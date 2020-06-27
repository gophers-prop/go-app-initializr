package api

import (
	"{{ .AppName}}/pkg/jwt"
	"goji.io/pat"
)

//registerRoute function to add your routes
func (ws *WebServer) registerRoute() {

	ws.server.HandleFunc(pat.Post("/welcome"), welcome)
	ws.server.HandleFunc(pat.Get("/v1/user/:userID"), handleGetUser)
	ws.server.HandleFunc(pat.Get("/v1/user"), handleGetAllUser)
	ws.server.HandleFunc(pat.Post("/v1/user"), handleCreateUser)
	ws.server.HandleFunc(pat.Put("/v1/user/:userID"), handleUpdateUser)
	ws.server.HandleFunc(pat.Delete("/v1/user/:userID"), handleDeleteUser)
	ws.server.HandleFunc(pat.Post("/v1/login"),handleLogin)
	ws.server.HandleFunc(pat.Post("/v1/signup"),handleSignUp)
	ws.server.HandleFunc(pat.Post("/v1/refresh"),jwt.MWHandle(handleRefreshToken))
	ws.server.HandleFunc(pat.Put("/v1/password"),handleUpdatePassword)
	
}
