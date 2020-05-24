package api

import (
	"{{ .AppName }}/consts"
	"github.com/gin-gonic/gin"
	"fmt"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"{{ .AppName }}/pkg/user"
	"net/http"
	"{{ .AppName }}/pkg/auth"
	"{{ .AppName }}/pkg/jwt"
)

func welcome(ctx *gin.Context){
	fmt.Println("welcome")
}

func handleCreateUser(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.CREATE_USER}} {{end}}
	var newUser user.User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	_,newUserError := newUser.Create()
	if newUserError != nil{
		ctx.JSON(newUserError.HttpStatus,gin.H{"error":newUserError.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,"User added successfully")
}

func handleGetUser(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.GET_USER}} {{end}}
	newUser := user.User{
		ID : ctx.Param("userID"),
	}
	response,getUserError := newUser.Get()
	if getUserError != nil{
		ctx.JSON(getUserError.HttpStatus,gin.H{"error":getUserError.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,response)
	return	
}

func handleUpdateUser(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.UPDATE_USER}} {{end}}
	var newUser user.User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	updateErr := newUser.Update()
	if updateErr != nil{
		//TODO: log error
		ctx.JSON(updateErr.HttpStatus,gin.H{"error":updateErr.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,"User updated successfully")
	return
}

func handleDeleteUser(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.DELETE_USER}} {{end}}
	newUser := user.User{
		ID : ctx.Param("userID"),
	}
	deleteUserErr := newUser.Delete()
	if deleteUserErr != nil{
		//TODO: log error
		ctx.JSON(deleteUserErr.HttpStatus,gin.H{"error":deleteUserErr.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,"User updated successfully")
}

func handleGetAllUser(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.GET_ALL_USER}} {{end}}
	response,getUserErr := user.GetAll()
	if getUserErr != nil{
		//TODO: log error
		ctx.JSON(getUserErr.HttpStatus,gin.H{"error":getUserErr.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,response)
}

func handleLogin(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.LOGIN}} {{end}}
	var auth auth.Auth
	if err := ctx.ShouldBind(&auth);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	resp , loginErr := auth.LogIn()
	if loginErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.LOGIN_ERROR}} {{end}}
		ctx.JSON(loginErr.HttpStatus,gin.H{"error":loginErr.Err.Error()})
	   return
	}
	ctx.JSON(http.StatusOK,resp)
}

func handleSignUp(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.SIGNUP}} {{end}}	
	var auth auth.Auth
	if err := ctx.ShouldBind(&auth);err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	resp , signupErr := auth.SignUp()
	 if signupErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.SIGNUP_ERROR}} {{end}}	
	    ctx.JSON(signupErr.HttpStatus,gin.H{"error":signupErr.Err.Error()})
		return
	 }
	 ctx.JSON(http.StatusOK,resp)
}

func handleRefreshToken(ctx *gin.Context){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.REFRESH_TOKEN}} {{end}}	
	token,err := jwt.GetTokenFromReq(ctx)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.EMPTY_TOKEN_ERROR}} {{end}}	
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	email,err := jwt.GetValueFromAKeyInToken(token,consts.JWT_EMAIL_KEY)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.ERROR_KEY_NOT_FOUND_JWT}} {{end}}
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	var auth auth.Auth
	auth.Email = email
	newTokenResp,refreshTokenErr := auth.RefreshToken()
	if refreshTokenErr != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.ERROR_REFRESH_TOKEN}} {{end}}
		ctx.JSON(refreshTokenErr.HttpStatus,gin.H{"error":refreshTokenErr.Err.Error()})
	}
	ctx.JSON(http.StatusOK,newTokenResp)

}
func handleUpdatePassword(ctx *gin.Context){
}



