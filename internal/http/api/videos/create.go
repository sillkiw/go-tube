package videosapi

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	apierrors "github.com/sillkiw/gotube/internal/http/api/apierrors"
	"github.com/sillkiw/gotube/internal/http/api/videos/dto"
	"github.com/sillkiw/gotube/internal/http/httpjson"
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
		l.Error("failed to decode request body",
			slog.Any("err", err),
		)
		// TODO: more detailed json error response
		errResp := apierrors.New("failed_decode_json", "failed to decode request body")
		httpjson.WriteJSON(w, r, http.StatusBadRequest, errResp)
		return
	}
	l.Debug("request was decoded", slog.Any("req", req))

	if verrs := vh.validator.CreateRequest(req); !verrs.Empty() {
		l.Error("failed to validate request",
			slog.Any("err", verrs),
		)
		code, errResp := apierrors.Map(err)
		httpjson.WriteJSON(w, r, code, errResp)
		return
	}

	id, err


}
