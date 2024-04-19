package routes

import (
	"github.com/alefwhite/api-users-go/docs/custom"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	docsURL = "http://localhost:8080/docs/doc.json"
)

// @title		API users
// @version	1.0
// @in			header
// @name		Authorization
func InitDocsRoutes(r chi.Router) {
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL),
		httpSwagger.AfterScript(custom.JS),
		httpSwagger.DocExpansion("none"),
		httpSwagger.UIConfig(map[string]string{
			"defaultModelsExpandDepth": `"-1"`,
		}),
	))
}
