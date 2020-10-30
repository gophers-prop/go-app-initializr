package api

import (
	"{{ .AppName }}/consts"
	"encoding/json"
	"fmt"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"{{ .AppName }}/pkg/user"
	"{{ .AppName }}/utils"
	"net/http"
	"{{ .AppName }}/pkg/auth"
	"{{ .AppName }}/pkg/jwt"
	"goji.io/pat"
)

func welcome(w http.ResponseWriter, r *http.Request){
	fmt.Println("welcome")
	fmt.Fprintf(w,"welcome")
}


func handleCreateUser(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.CREATE_USER}} {{end}}
	var newUser user.User
	unMarshalErr := json.NewDecoder(r.Body).Decode(&newUser)
	if unMarshalErr != nil{
		utils.GenerateAndSendErrorResponse(http.StatusInternalServerError,"Error while reading request data",unMarshalErr,w)
	return
	}
	_,newUserError := newUser.Create()
	fmt.Println(newUserError)
	if newUserError != nil{
		utils.GenerateAndSendErrorResponse(newUserError.HttpStatus,"Error while creating new user",newUserError.Err,w)
		return
	}
	
   		 w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,`{"msg":"User created successfull"`)
}


func handleGetUser(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.GET_USER}} {{end}}
	newUser := user.User{
		ID : pat.Param(r,"userID"),
	}
	response,getUserError := newUser.Get()
	if getUserError != nil{
		//TODO: log error
		utils.GenerateAndSendErrorResponse(getUserError.HttpStatus,"Error fetching details of the user",getUserError.Err,w)
		return
	}
	utils.GenerateAndSendSuccessResponse(http.StatusOK,"User details",response,w)
	return	
}
func handleUpdateUser(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.UPDATE_USER}} {{end}}
	var newUser user.User
	if decodeErr := json.NewDecoder(r.Body).Decode(&newUser);decodeErr != nil{
		utils.GenerateAndSendErrorResponse(http.StatusInternalServerError,"Error while reading request data",decodeErr,w)
		return
	}
	updateErr := newUser.Update()
	if updateErr != nil{
		utils.GenerateAndSendErrorResponse(updateErr.HttpStatus,"Error while updating user data",updateErr.Err,w)
		return
	}
	utils.GenerateAndSendSuccessResponse(http.StatusOK,"User updated successfully",nil,w)
	return
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.DELETE_USER}} {{end}}
	newUser := user.User{
		ID : pat.Param(r,"userID"),
	}
	deleteUserErr := newUser.Delete()
	if deleteUserErr != nil{
		utils.GenerateAndSendErrorResponse(deleteUserErr.HttpStatus,"Error while deleting user data",deleteUserErr.Err,w)
		return
	}
	utils.GenerateAndSendSuccessResponse(http.StatusOK,"User deleted successfully",nil,w)
}

func handleGetAllUser(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.GET_ALL_USER}} {{end}}
	response,getUserErr := user.GetAll()
	if getUserErr != nil{
		utils.GenerateAndSendErrorResponse(getUserErr.HttpStatus,"Error while fetching users ",getUserErr.Err,w)
		return
	}
	utils.GenerateAndSendSuccessResponse(http.StatusOK,"Users list ",response,w)
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.LOGIN}} {{end}}
	var auth auth.Auth
	if decodeErr := json.NewDecoder(r.Body).Decode(&auth);decodeErr != nil{
		utils.GenerateAndSendErrorResponse(http.StatusInternalServerError,"Error while reading request data",decodeErr,w)
		return
	}
	resp , loginErr := auth.LogIn()
	if loginErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.LOGIN_ERROR}} {{end}}
	  utils.GenerateAndSendErrorResponse(loginErr.HttpStatus,"Error while logging ",loginErr.Err,w)
	   return
	}
	utils.GenerateAndSendSuccessResponse(http.StatusOK,"Login Successfull  ",resp,w)
}


func handleSignUp(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.SIGNUP}} {{end}}	
	var auth auth.Auth
	if decodeErr := json.NewDecoder(r.Body).Decode(&auth);decodeErr != nil{	
		utils.GenerateAndSendErrorResponse(http.StatusInternalServerError,"Error while reading request data",decodeErr,w)
		return
	}
	resp , signupErr := auth.SignUp()
	 if signupErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.SIGNUP_ERROR}} {{end}}	
		utils.GenerateAndSendErrorResponse(signupErr.HttpStatus,"Error while sign up ",signupErr.Err,w)
		return
	 }
	 utils.GenerateAndSendSuccessResponse(http.StatusOK,"Signup Successfull  ",resp,w)
}

func handleRefreshToken(w http.ResponseWriter, r *http.Request){
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.REFRESH_TOKEN}} {{end}}	
	token,err := jwt.GetTokenFromReq(r)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.EMPTY_TOKEN_ERROR}} {{end}}	
		utils.GenerateAndSendErrorResponse(http.StatusBadRequest,"Error while reading jwt token ",err,w)
	}
	email,err := jwt.GetValueFromAKeyInToken(token,consts.JWT_EMAIL_KEY)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.ERROR_KEY_NOT_FOUND_JWT}} {{end}}
		utils.GenerateAndSendErrorResponse(http.StatusBadRequest,"Error while reading jwt token ",err,w)
	}
	var auth auth.Auth
	auth.Email = email
	newTokenResp,refreshTokenErr := auth.RefreshToken()
	if refreshTokenErr != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Handler.ERROR_REFRESH_TOKEN}} {{end}}
		utils.GenerateAndSendErrorResponse(http.StatusBadRequest,"Error while generating jwt token",err,w)
	}
	utils.GenerateAndSendSuccessResponse(http.StatusOK,"Token generated successfully",newTokenResp,w)

}
func handleUpdatePassword(w http.ResponseWriter, r *http.Request){
}



