package app

import (
	"fmt"
	"gotube/internal/cookie"
	"gotube/internal/templates"
	"gotube/internal/user"
	"gotube/internal/utils"

	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Login handles basic HTTP authentication: it retrieves the username and password,
// checks them against users.yaml, and sets a signed cookie on success. Login handles requests to the /auth endpoint.
func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	ip := r.RemoteAddr
	if !ok {
		app.logger.Info("AUTH FAIL: no credentials", slog.String("ip", ip))
		w.Header().Set("WWW-Authenticate", `Basic realm="RetryLogin"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	app.logger.Info("AUTH ATTEMPT",
		slog.String("ip", ip),
	)

	// Find the user with the given username
	var currentUser *user.User
	for i := range app.users {
		if app.users[i].Username == username {
			currentUser = &app.users[i]
			break
		}
	}
	if currentUser == nil {
		app.logger.Info("AUTH FAIL: user not found",
			slog.String("user", username),
			slog.String("ip", ip),
		)
		app.renderError(w, "Invalid username or password")
		return
	}

	// Verify the password
	err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(password))
	if err != nil {
		app.logger.Info("AUTH FAIL: invalid password",
			slog.String("user", username),
			slog.String("ip", ip),
		)
		// Passwords don't match, show an error message
		app.renderError(w, "Invalid username or password")
		return
	}

	// Authentication successful
	app.logger.Info("AUTH SUCCESS",
		slog.String("user", username),
		slog.String("ip", ip),
		slog.String("role", currentUser.Role),
	)

	// Create cookie
	expiration := time.Now().Add(24 * time.Hour)
	value := currentUser.Username + "|" + currentUser.Role
	cookie := cookie.CreateSignedCookie("auth", value, expiration)
	http.SetCookie(w, cookie)

	// Redirect the user to the home page
	http.Redirect(w, r, "/", http.StatusFound)

}

func (app *Application) upload(w http.ResponseWriter, r *http.Request) {
	// Upload a file
	file, header, err := r.FormFile("video")
	if err != nil {
		app.renderError(w, "Failed to upload the file. Please try again.")
		return
	}
	defer file.Close()
	app.logger.Info("File upload started", slog.String("filename", header.Filename))

	// Check if the file size is valid
	if header.Size > app.config.Upload.MaxSize {
		maxSizeMB := app.config.Upload.MaxSize / (1024 * 1024)
		app.logger.Info("File size is too big",
			slog.String("file", header.Filename),
			slog.Int64("limit_bytes", app.config.Upload.MaxSize),
		)

		msg := fmt.Sprintf(`File "%s" exceeds the max allowed size of %dMB.`, header.Filename, maxSizeMB)
		app.renderError(w, msg)
		return
	}

	// Check if the file name is valid
	filename := header.Filename
	if len(filename) > app.config.Upload.MaxNameLength || !utils.IsSafeFileName(filename) {
		app.logger.Info("Invalid file name", slog.String("file", header.Filename))
		msg := fmt.Sprintf("Invalid file name: either it contains invalid characters or it's longer than %d characters", app.config.Upload.MaxNameLength)
		app.renderError(w, msg)
		return
	}

	// Check if the maxium video per h is reached
	if app.videoSrv.GetUploadedCount() >= app.config.Upload.MaxPerHour {
		app.logger.Info("Upload limit reached",
			slog.Int64("upload_count", app.videoSrv.GetUploadedCount()),
			slog.Int64("limit", app.config.Upload.MaxPerHour),
		)
		msg := fmt.Sprintf("Can't upload more than %d videos per hour", app.config.Upload.MaxPerHour)
		app.renderError(w, msg)
		return
	}

	// Get the file extension and remove it from the filename
	extension := path.Ext(filename)
	filenamenoext := strings.TrimSuffix(filename, extension)

	// Clean the filename and combine with upload path
	filePath := filepath.Join(app.config.Upload.Path, filepath.Clean(filename))

	// Get the absolute path to the upload directory
	absUploadPath, err1 := filepath.Abs(app.config.Upload.Path)

	// Get the absolute path to the requested file
	absFilePath, err2 := filepath.Abs(filePath)

	// If any error occurs while getting absolute paths, return internal error
	if err1 != nil || err2 != nil {
		app.renderError(w, "Internal path error")
		return
	}

	// Check if the file is truly inside the upload directory
	if !strings.HasPrefix(absFilePath, absUploadPath) {
		app.renderError(w, "Invalid file name")
		return
	}

	// Check if the file already exists
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		msg := "File already exists: " + filename
		app.renderError(w, msg)
		return
	}

	// Create the destination file where the uploaded content will be saved
	out, err := os.Create(filePath)
	if err != nil {
		app.logger.Error("Failed to save file",
			slog.String("path", filePath),
			slog.String("error", err.Error()),
		)
		app.renderError(w, "Failed to save file")
		return
	}
	defer out.Close()

	// Copy the uploaded file's content into the newly created file
	_, err = io.Copy(out, file)
	if err != nil {
		app.logger.Error("Failed to upload file",
			slog.String("path", filePath),
			slog.String("error", err.Error()),
		)
		app.renderError(w, "Failed to upload file")
		return
	}
	app.videoSrv.IncrementUploaded()

	// Start the video conversion process in a separate goroutine
	go app.videoSrv.StartConvertVideo(filePath, filenamenoext)

	// Prepare data for the template that confirms upload to the user
	p := &templates.PageUploaded{
		FileName:      filename,
		FileNameNoExt: filenamenoext,
		QuequeSize:    0,
	}

	// Render the "uploaded" template with the provided data
	// This gives the user feedback that the upload succeeded and conversion has started
	app.render(w, "uploaded", p)
}

func (app *Application) showUploadForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, "sendfile", nil)
}

// listFolderHandler handles GET /lst (and /) requests, paginates video folders,
// and renders the filelist template.
func (app *Application) listFolderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0 {
			page = n
		}
	}
	dir := app.config.Video.ConvertPath
	perPage := app.config.Video.PerPage
	folders, total, err := listFolders(dir, page, perPage)
	if err != nil {
		app.logger.Error("Failed to read video directory",
			slog.String("error", err.Error()),
		)
		app.renderError(w, "No video available")
		return
	}
	totalPages := (total + perPage - 1) / perPage

	data := &templates.PageList{
		Files:     folders,
		Page:      page,
		TotalPage: totalPages,
		PrevPage:  max(1, page-1),
		NextPage:  min(totalPages, page+1),
		CanDelete: app.canDeleteQ(r, app.config.Auth.DeleteAllowedRoles),
	}
	app.render(w, "filelist", data)
}

func listFolders(dirPath string, page, perPage int) ([]utils.FolderInfo, int, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot read video directory: %w", err)
	}

	// Filter and collect folderInfo
	infos := make([]utils.FolderInfo, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			fi, err := e.Info()
			if err != nil {
				continue // skip unreadable
			}
			infos = append(infos, utils.FolderInfo{
				Name:    e.Name(),
				ModTime: fi.ModTime(),
			})
		}
	}

	total := len(infos)
	if total == 0 {
		return nil, 0, fmt.Errorf("no videos available")
	}

	// Sort by modification time (newest first)
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].ModTime.After(infos[j].ModTime)
	})

	// Calculate slice indices
	start := (page - 1) * perPage
	if start >= total {
		return nil, total, fmt.Errorf("invalid page number: %d", page)
	}
	end := start + perPage
	if end > total {
		end = total
	}

	return infos[start:end], total, nil
}

func (app *Application) faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, app.config.UI.Favicon)
}

func (app *Application) quequeSize(w http.ResponseWriter, r *http.Request) {
	p := &templates.PageQueque{
		QuequeSize: int(app.videoSrv.GetQueueLen()),
	}
	app.render(w, "queque", p)
}

// deleteVideo handles DELETE requests (or GET with ?videoname=…) to remove
// a video’s converted directory. It validates the name, attempts deletion,
// logs any error, and then redirects back to the video list.
func (app *Application) deleteVideo(w http.ResponseWriter, r *http.Request) {
	// // Only allow POST/DELETE methods
	// if r.Method != http.MethodPost && r.Method != http.MethodDelete {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// Parse and validate the videoname parameter
	videoname := r.URL.Query().Get("videoname")
	if videoname == "" || !utils.IsSafeFileName(videoname) {
		app.logger.Info("deleteVideo: invalid videoname", slog.String("videoname", videoname))
		http.Error(w, "Invalid video name", http.StatusBadRequest)
		return
	}

	// Build the path and attempt to delete
	dir := filepath.Join(app.config.Video.ConvertPath, videoname)
	if err := os.RemoveAll(dir); err != nil {
		app.logger.Error("deleteVideo: failed to remove directory",
			slog.String("path", dir),
			slog.String("error", err.Error()),
		)
		app.renderError(w, "Failed to delete video")
		return
	}

	app.logger.Info("deleteVideo: successfully deleted",
		slog.String("videoname", videoname),
		slog.String("path", dir),
	)

	// Redirect back to the video list (or send JSON success for an API)
	http.Redirect(w, r, "/lst", http.StatusSeeOther)
}

// videoPlayerHandler renders the video playback page.
// It supports three modes: embedded iframe, no-JS fallback, and full JS player.
// Filename is validated and sanitized before rendering.
func (app *Application) videoPlayerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters once
	q := r.URL.Query()
	videoname := q.Get("videoname")

	// Validate filename
	if !utils.IsSafeFileName(videoname) || len(videoname) > app.config.Upload.MaxNameLength {
		app.logger.Info("invalid video name",
			slog.String("videoname", videoname),
		)
		app.renderError(w, fmt.Sprintf(
			"Invalid file name: only A-Z, a-z, 0-9, -, _ allowed and max %d characters",
			app.config.Upload.MaxNameLength,
		))
		return
	}

	// Determine display mode
	embedded := q.Get("embedded") == "1" && app.config.Video.AllowEmbed
	noJS := q.Get("nojs") == "1"

	// Select template and data
	var (
		tmplName string
		data     interface{}
	)

	switch {
	case embedded:
		tmplName = "embedded"
		data = &templates.PageVPEMB{VidNm: videoname}

	case noJS:
		tmplName = "vpnojs"
		data = &templates.PageVPNoJS{VidNm: videoname}

	default:
		tmplName = "vp"
		data = &templates.PageVP{
			VidNm: videoname,
			Embed: app.config.Video.AllowEmbed,
		}
	}

	// Render the chosen template
	app.render(w, tmplName, data)
}
