package user

import(
	"errors"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"{{ .AppName }}/models"
	"net/http"
)

const(
	currentContext = "user"
)

// User : Struct for getting user data
type User struct {
	Name   string        `form:"name" json:"name" xml:"name"  binding:"required"`
	Email string 	`form:"email" json:"email" xml:"email"  binding:"required"`
	ID   string        `form:"id" json:"id" xml:"id"  binding:"required"`
	Age string       `form:"age" json:"age" xml:"age"  binding:"required"`
	
}

//GetAll fetch all users
func GetAll()([]User,*models.Error){

	//this line is db dependent
	resp , err := getAllUsers()
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.User.GET_ALL}} {{end}}
		return nil,err
	}
	return resp,nil
}



//Create logic for creating new user
func (user *User)Create()(*User,*models.Error){
	//this line db dependent
	err := create(user)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.User.CREATE}} {{end}}
		return nil,err
	}
	return user,nil
}


//Get get user details
func (user *User)Get()(*User,*models.Error){
	id := user.ID
	if id == ""{
		return nil,&models.Error{Err:errors.New("Id cannot be empty"),HttpStatus:http.StatusBadRequest}
	}

	//TODO: this is db dependent
	response,err := getByID(id)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.User.GET}} {{end}}
		return nil,err
	}
	return response,nil
}

//Update updating user details
func (user *User)Update()(*models.Error){
	//TODO: this is db dependent
	err := update(user)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.User.UPDATE}} {{end}}
		return err
	}
	return nil
}

//Delete deleting user 
func (user *User)Delete()(*models.Error){	
	//TODO: this is db dependent
	err := deleteByID(user.ID)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.User.DELETE}} {{end}}
		return err
	}
	return nil
}