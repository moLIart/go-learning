package handlers

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	// This is for Swagger documentation
	_ "github.com/moLIart/go-course/internal/docs"
)

func SwaggerUIHandler() http.Handler {
	return httpSwagger.WrapHandler
}
