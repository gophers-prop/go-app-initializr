package api

import (
	"github.com/go-martini/martini"
    "{{ .AppName}}/pkg/jwt"
)

//registerRoute function to add your routes
func (ws *WebServer) registerRoute() {

	ws.server.Post("/welcome", welcome)

	//grouping routes based on their version
	ws.server.Group("/v1",func(r martini.Router){
		r.Get("/user/:userID", handleGetUser)
		r.Get("/user", handleGetAllUser)
		r.Post("/user", handleCreateUser)
		r.Put("/user/:userID", handleUpdateUser)
		r.Delete("/user/:userID", handleDeleteUser)
		r.Post("/login",handleLogin)
		r.Post("/signup",handleSignUp)
		r.Post("/refresh",jwt.MWHandle(handleRefreshToken))
		r.Put("/password",handleUpdatePassword)
	})
}
