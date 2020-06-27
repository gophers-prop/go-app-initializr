package models

//Error error model for all request
type Error struct{
	Err error 
	HttpStatus int
}
type ErrorResponse struct{
	Message string   `json:"message"`
	Error   interface{} `json:"error"`
}