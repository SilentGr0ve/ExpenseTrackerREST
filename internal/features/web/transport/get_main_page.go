package web_transport

import (
	"net/http"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/logger"
	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http/response"
)

func (h *WebHTTPHandler) GetMainPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	rh := response.NewResponseHandler(rw, log)

	html, err := h.webService.GetMainPage()
	if err != nil {
		rh.ErrorResponse(
			err,
			"failed to get index.html for main page",
		)
	}

	rh.HTMLResponse(html)

}
