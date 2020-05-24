package api

import (
	"github.com/go-martini/martini"
)

const(
	SWAGGER_PORT = ":9090"
)

//SwaggerServer hosts a swagger docs server
func SwaggerServer() {

	server := martini.Classic()
	server.Use(martini.Recovery())
    server.Use(martini.Static("./assets/swagger",martini.StaticOptions{Prefix:"/swagger"}))
	//server.("/swagger", http.Dir("./assets/swagger"))
	server.Use(martini.Static("./assets/swagger/favicon.ico",martini.StaticOptions{Prefix:"/favicon.ico"}))
	//server.StaticFile("/favicon.ico", "../assets/swagger/favicon.ico")

	server.RunOnAddr(SWAGGER_PORT)

}
