package videosapi

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	apierrors "github.com/sillkiw/gotube/internal/http/api/errors"
	"github.com/sillkiw/gotube/internal/http/api/videos/dto"
	"github.com/sillkiw/gotube/internal/http/httpjson"
)

const (
	decodeJsonFailed = "failed_decode_json"
)

func (vh *VideosHandler) create(w http.ResponseWriter, r *http.Request) {
	const op = "http.api.videos.create"
	l := vh.logger.With(
		slog.String("op", op),
		slog.String("requets_id", middleware.GetReqID(r.Context())),
	)

	var req dto.CreateRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		l.Error("failed to decode request body",
			slog.Any("err", err),
		)
		httpjson.WriteError(w, r, http.StatusBadRequest, decodeJsonFailed, "failed to decode request body")
		return
	}
	l.Debug("request was decodec", slog.Any("req", req))

	if err := vh.validator.CreateRequest(req); err != nil {
		l.Error("failed to validate request",
			slog.Any("err", err),
		)
		code, errCode, msg := apierrors.CreateRequestValidation(err, vh.validator)
		httpjson.WriteError(w, r, code, errCode, msg)
		return
	}
}
