package api

import (
    "{{ .AppName }}/pkg/jwt"
)

//registerRoute function to add your routes
func (ws *WebServer) registerRoute() {

	ws.server.POST("/welcome", welcome)

	//grouping routes based on their version
	v1 := 	ws.server.Group("/v1")
	{
		v1.GET("/user/:userID", handleGetUser)
		v1.GET("/user", handleGetAllUser)
		v1.POST("/user", handleCreateUser)
		v1.PUT("/user/:userID", handleUpdateUser)
		v1.DELETE("/user/:userID", handleDeleteUser)
		v1.POST("/login",handleLogin)
		v1.POST("/signup",handleSignUp)
		v1.POST("/refresh",jwt.MWHandle(handleRefreshToken))
		v1.PUT("/password",handleUpdatePassword)
	}

	
	

}
