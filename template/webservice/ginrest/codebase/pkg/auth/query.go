package auth

import (
	"errors"
	"fmt"
	"{{ .AppName }}/pkg/database"
	"net/http"
	"{{ .AppName }}/models"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)

const (
	AUTH_TABLE_NAME  = "auth"
	AUTH_DB_EMAIL    = "email"
	AUTH_DB_PASSWORD = "password"
)

func signup(a *Auth) *models.Error {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()

	sqlQuery := fmt.Sprintf("INSERT INTO %s (%s,%s) VALUES (?,?)", AUTH_TABLE_NAME, AUTH_DB_EMAIL, AUTH_DB_PASSWORD)

	insertReq, err := db.Prepare(sqlQuery)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer insertReq.Close()
	insertReq.Exec(a.Email, a.Password)
	return nil
}

func getByEmail(email string) (*Auth, *models.Error) {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer db.Close()

	sqlQuery := fmt.Sprintf("SELECT %s,%s FROM %s WHERE %s = ?", AUTH_DB_EMAIL, AUTH_DB_PASSWORD, AUTH_TABLE_NAME, AUTH_DB_EMAIL)

	rows, err := db.Query(sqlQuery, email)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	defer rows.Close()

	if rows.Next() {
		var email, password string
		err = rows.Scan(&email, &password)
		if err != nil {
			{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
			return nil, &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
		}

		return &Auth{
			Email:    email,
			Password: password,
		}, nil
	}
	return nil,  &models.Error{Err:errors.New("No User found"),HttpStatus:http.StatusNotFound}
}

func update(email, password string) *models.Error {
	db, err := database.GetConn()
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	sqlQuery := fmt.Sprintf("UPDATE %s SET %s=?,%s=?  WHERE %s = ?", AUTH_TABLE_NAME, AUTH_DB_EMAIL, AUTH_DB_PASSWORD, AUTH_DB_EMAIL)

	_, err = db.Exec(sqlQuery, email, password, email)

	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.QUERY_ERROR}} {{end}}
		return &models.Error{Err:err,HttpStatus:http.StatusInternalServerError}
	}
	return nil
}

