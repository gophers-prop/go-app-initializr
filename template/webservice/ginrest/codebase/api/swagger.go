package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const(
	SWAGGER_PORT = ":9090"
)

//SwaggerServer hosts a swagger docs server
func SwaggerServer() {

	server := gin.Default()

	server.Use(gin.Recovery())

	server.StaticFS("/swagger", http.Dir("./assets/swagger"))

	server.StaticFile("/favicon.ico", "../assets/swagger/favicon.ico")

	server.Run(SWAGGER_PORT)

}
