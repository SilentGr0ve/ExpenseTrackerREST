package auth_transport

type AuthHTTPHandler struct {
	authService AuthServiceInterface
}

type AuthServiceInterface interface{}
