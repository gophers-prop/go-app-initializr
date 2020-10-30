package auth

import (

	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"{{ .AppName }}/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"{{ .AppName }}/models"
	"net/http"
)

type Auth struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpLoginResp struct{
	Email string `json:"email"`
	Token string `json:"token"`
}
const(
	currentContext = "auth"
)
//SignUp registers an account in application
func (a *Auth)SignUp()(*SignUpLoginResp,*models.Error){
	var err error
	a.Password ,err = hash(a.Password)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Auth.ERROR_ENCRYPTING_PASSWORD}} {{end}}
		return nil,&models.Error{Err:err,HttpStatus:http.StatusBadRequest}
	}
	var signUpErr *models.Error

	{{ if .Database.ImportPath }}{{ .Database.Messages.General.SIGNUP}} {{end}}
	
	if signUpErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Auth.ERROR_SIGNING_UP}} {{end}}
		return nil,signUpErr
	}

	token,tokenErr := jwt.GenerateAccountLoginToken(a.Email)
	if tokenErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.ERROR_GENERATING_NEW_TOKEN}} {{end}}
		return nil,&models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	return &SignUpLoginResp{
		Email: a.Email,
		Token: token,
	},nil    
}

//LogIn logging logic 
func (a *Auth)LogIn()(*SignUpLoginResp,*models.Error){
	
	var auth *Auth
	var err *models.Error
	
	{{ if .Database.ImportPath }}{{ .Database.Messages.General.LOGIN}} {{end}}

	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Auth.ERROR_LOGGING_IN}} {{end}}
		return nil,err
	}
	//comparing password from db 
	if !checkHash(a.Password,auth.Password){
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Auth.ERROR_INCORRECT_PASSWORD}} {{end}}
		return nil,&models.Error{Err:errors.New("Invalid password"),HttpStatus:http.StatusForbidden}
	}
	token,tokenErr := jwt.GenerateAccountLoginToken(a.Email)
	if tokenErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.ERROR_GENERATING_NEW_TOKEN}} {{end}}
		return nil,&models.Error{Err:tokenErr,HttpStatus:http.StatusInternalServerError}
	}
	return &SignUpLoginResp{
		Email: a.Email,
		Token: token,
	},nil
}

//RefreshToken refreshed the token and return it as map
func (a *Auth)RefreshToken()(map[string]string,*models.Error){
	token,tokenErr := jwt.GenerateAccountLoginToken(a.Email)
	if tokenErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.ERROR_GENERATING_NEW_TOKEN}} {{end}}
		return nil,&models.Error{Err:tokenErr,HttpStatus:http.StatusInternalServerError}
	}
	return map[string]string{"token":token},nil
}

func hash(password string)(string,error){
	 pass , err := bcrypt.GenerateFromPassword([]byte(password),12)
	 return string(pass),err
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
