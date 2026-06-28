package response

import (
	"encoding/json"
	"errors"
	"net/http"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"go.uber.org/zap"
)

type ResponseHandler struct {
	http.ResponseWriter
	log *logger.Logger
}

func NewResponseHandler(rw http.ResponseWriter, log *logger.Logger) *ResponseHandler {
	return &ResponseHandler{
		ResponseWriter: rw,
		log:            log,
	}
}

func (h *ResponseHandler) JSONResponse(statusCode int, responseBody any) {
	h.Header().Set("Content-Type", "application/json")

	h.WriteHeader(statusCode)

	if err := json.NewEncoder(h.ResponseWriter).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response:", zap.Error(err))
	}
}

func (h *ResponseHandler) ErrorResponse(err error, message string) {
	var (
		statusCode int
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
	case errors.Is(err, core_errors.ErrUnauthorized):
		statusCode = http.StatusUnauthorized
	case errors.Is(err, core_errors.ErrForbidden):
		statusCode = http.StatusForbidden
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict

	default:
		statusCode = http.StatusInternalServerError
	}

	if statusCode >= 500 {
		h.log.Error(message, zap.Error(err))
	} else {
		h.log.Warn(message, zap.Error(err))
	}

	h.JSONResponse(
		statusCode,
		map[string]string{
			"error": message,
		},
	)
}

func (h *ResponseHandler) PanicResponse(p any, message string) {
	h.log.Error(
		message,
		zap.Any("panic", p),
		zap.Stack("stacktrace"),
	)

	h.JSONResponse(
		http.StatusInternalServerError,
		map[string]string{"error": message},
	)
}
