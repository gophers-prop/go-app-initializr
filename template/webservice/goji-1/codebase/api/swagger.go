package api

import (
	
	"net/http"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)

const(
	SWAGGER_PORT = ":9090"
)

//SwaggerServer hosts a swagger docs server
func SwaggerServer() {

	http.Handle("/", http.StripPrefix("/swagger", http.FileServer(http.Dir("./assets/swagger"))))
	swaggerErr := http.ListenAndServe(SWAGGER_PORT, nil)
  if swaggerErr != nil {
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Swagger.SERVER_STARTED_ERROR}} {{end}}
	return
	}
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Swagger.SERVER_STARTED_SUCCESS}} {{end}}
}
