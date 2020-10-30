package user

import (
	"errors"
	"fmt"
	"{{ .AppName }}/pkg/database"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"{{ .AppName }}/models"
	"net/http"
)

const (
	USER_TABLE_NAME = "user"
	USER_DB_EMAIL   = "email"
	USER_DB_NAME    = "name"
	USER_DB_AGE     = "age"
	USER_DB_ID      = "id"
)

func getAllUsers() ([]User, *models.Error) {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()
	sqlQuery := fmt.Sprintf("SELECT %s,%s,%s,%s FROM %s", USER_DB_NAME, USER_DB_EMAIL, USER_DB_AGE, USER_DB_ID, USER_TABLE_NAME)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer rows.Close()
	var response []User
	for rows.Next() {
		var user User
		err = rows.Scan(&(user.Name), &(user.Email), &(user.Age), &(user.ID))
		if err != nil {
			{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
			return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
		}
		response = append(response, user)
	}
	return response, nil
}

func update(user *User) *models.Error {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()
	sqlQuery := fmt.Sprintf("UPDATE %s SET %s=?,%s=?,%s=?,%s=? WHERE %s = ?", USER_TABLE_NAME, USER_DB_NAME, USER_DB_EMAIL, USER_DB_AGE, USER_DB_ID, USER_DB_ID)
	_, err = db.Exec(sqlQuery, user.Name, user.Email, user.Age, user.ID, user.ID)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	return nil
}

func create(u *User) *models.Error {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()

	sqlQuery := fmt.Sprintf("INSERT INTO %s (%s,%s,%s,%s) VALUES (?,?,?,?)", USER_TABLE_NAME, USER_DB_NAME, USER_DB_EMAIL, USER_DB_AGE, USER_DB_ID)

	insertReq, err := db.Prepare(sqlQuery)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer insertReq.Close()
	insertReq.Exec(u.Name, u.Email, u.Age, u.ID)
	return nil
}

func deleteByID(id string) *models.Error {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE ID = ?", USER_TABLE_NAME)
	_, err = db.Exec(sqlQuery, id)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	return nil
}

func getByID(id string) (*User, *models.Error) {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()
	sqlQuery := fmt.Sprintf("SELECT %s,%s,%s,%s FROM %s WHERE %s = ?", USER_DB_NAME, USER_DB_EMAIL, USER_DB_AGE, USER_DB_ID, USER_TABLE_NAME, USER_DB_ID)
	rows, err := db.Query(sqlQuery, id)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer rows.Close()
	if rows.Next() {
		var user User
		err = rows.Scan(&(user.Name), &(user.Email), &(user.Age), &(user.ID))
		if err != nil {
			{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
			return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
		}
		return &user, nil
	}
	return nil, &models.Error{Err:errors.New("Cannot find user"),HttpStatus:http.StatusNotFound}
}
