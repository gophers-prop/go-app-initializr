package models

//Error error model for all request
type Error struct{
	Err error 
	HttpStatus int
}