package web_transport

import (
	httpserver "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/transport/http"
)

type WebHTTPHandler struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
}

func NewWebHTTPHandler(webService WebService) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []httpserver.Route {
	return []httpserver.Route{
		{
			Path:    "/",
			Handler: h.GetMainPage,
		},
	}
}
