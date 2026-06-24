package auth_service

type AuthHTTPHandler struct {
	authService AuthServiceInterface
}

type AuthServiceInterface interface{}
