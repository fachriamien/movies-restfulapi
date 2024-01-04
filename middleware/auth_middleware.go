package middleware

import (
	"fchramn/movies-restfulapi/helper"
	"fchramn/movies-restfulapi/model/web"
	"net/http"
)

type AuthMiddleWare struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleWare {
	return &AuthMiddleWare{Handler: handler}
}

func (middleware *AuthMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "SECRETNUMBERID" == request.Header.Get("API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
