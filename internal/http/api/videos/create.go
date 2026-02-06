package videosapi

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	apierrors "github.com/sillkiw/gotube/internal/http/api/apierrors"
	"github.com/sillkiw/gotube/internal/http/api/videos/dto"
	"github.com/sillkiw/gotube/internal/http/httpjson"
	"github.com/sillkiw/gotube/internal/videos"
)

func (vh *VideosHandler) create(w http.ResponseWriter, r *http.Request) {
	const op = "http.api.videos.create"
	l := vh.logger.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req dto.CreateRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		l.Info("failed to decode request body",
			slog.Any("err", err),
		)
		// TODO: more detailed json error response
		body := apierrors.New("failed_decode_json", "failed to decode request body")
		httpjson.WriteJSON(w, r, http.StatusBadRequest, body)
		return
	}
	l.Debug("request was decoded", slog.Any("req", req))

	if verrs := vh.validator.CreateRequest(req); !verrs.Empty() {
		l.Info("failed to validate request",
			slog.Any("err", verrs),
		)
		status, body := apierrors.Map(err)
		httpjson.WriteJSON(w, r, status, body)
		return
	}

	video := videos.Video{
		Title:  req.Title,
		Size:   req.Size,
		Status: videos.StatusCreated,
	}

	id, err := vh.svc.Create(video)
	if err != nil {
		l.Error("failed to create video record",
			slog.Any("err", err),
		)
		status, body := apierrors.Map(err)
		httpjson.WriteJSON(w, r, status, body)
		return
	}

	l.Info("video record created",
		slog.String("id", id),
	)
	resp := dto.CreateResponse{
		ID:     id,
		Status: "created",
		Upload: dto.UploadDetails{
			Method: http.MethodPut,
			URL:    "/api/videos/" + id + "/upload",
			Headers: map[string]string{
				"Content-Type": req.ContentType,
			},
			MaxBytes: vh.validator.Cfg.UplLimit.MaxSize,
		},
		Links: dto.Links{
			Self: "/api/videos/" + id,
		},
	}
	httpjson.WriteJSON(w, r, http.StatusCreated, resp)
}
