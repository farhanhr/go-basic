package middleware

import (
	"farhanhr/go-restful-api/helper"
	"farhanhr/go-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middlware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "AUTHKEY" == request.Header.Get("X-Api-Key") {
		middlware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}