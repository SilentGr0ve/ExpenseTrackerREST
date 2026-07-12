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

	h.errorResponse(
		statusCode,
		message,
	)
}

func (h *ResponseHandler) errorResponse(statusCode int, msg string) {
	response := ErrorResponse{
		Error: msg,
	}

	h.JSONResponse(
		statusCode,
		response,
	)
}

func (h *ResponseHandler) PanicResponse(p any, msg string) {
	h.log.Error(
		msg,
		zap.Any("panic", p),
		zap.Stack("stacktrace"),
	)

	h.errorResponse(
		http.StatusInternalServerError,
		msg,
	)
}

func (h *ResponseHandler) NoContent() {
	h.WriteHeader(http.StatusNoContent)
}

func (h *ResponseHandler) HTMLResponse(html []byte) {
	h.ResponseWriter.WriteHeader(http.StatusOK)

	h.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := h.ResponseWriter.Write(html); err != nil {
		h.log.Error("write HTML response", zap.Error(err))
	}
}
