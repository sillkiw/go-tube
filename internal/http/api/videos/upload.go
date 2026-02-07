package videosapi

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sillkiw/gotube/internal/http/api/apierrors"
	"github.com/sillkiw/gotube/internal/http/httpjson"
)

func (vh *VideosHandler) upload(w http.ResponseWriter, r *http.Request) {
	const op = "http.api.videos.upload"
	l := vh.logger.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	videoID := chi.URLParam(r, "video_id")

	videoRec, err := vh.svc.Get(videoID)
	if err != nil {
		l.Info("failed to get video record", slog.Any("err", err))
		code, body := apierrors.Map(err)
		httpjson.WriteJSON(w, r, code, body)
		return
	}

	if videoRec.Status != "create" {
		l.Info("video record has another status",
			slog.String("id", videoID),
			slog.String("status", string(videoRec.Status)),
		)
		body := apierrors.New("not_create_status", "video status isn't create")
		httpjson.WriteJSON(w, r, http.StatusBadRequest, body)
		return
	}
	

}
