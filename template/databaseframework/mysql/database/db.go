package database

import (
	"{{ .AppName }}/consts"
	"database/sql"
 _ "github.com/go-sql-driver/mysql"
 "{{ .AppName }}/environment"
 {{ if .Logging.ImportPath }}
 "{{ .Logging.ImportPath }}"
 {{end}}
)
const(
	currentContext = "database"
)

//GetConn connecting to db 
//TODO:  this func is highly dependent on mysql , need to make it more modular
func GetConn()(*sql.DB,error){
	//get this from environment.go
	dbDetails := environment.GetDatabaseDetails()
	dbDriver := "mysql"
    dbUser := dbDetails[consts.DATABASE_USER_NAME]
    dbPass := dbDetails[consts.DATABSE_USER_PASSWORD]
	dbName := dbDetails[consts.DATABASE_NAME]
	
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return nil,err
	}
	err = db.Ping()
	if err != nil{
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Database.CONNECTION_ERROR}} {{end}}
		return nil,err
	}
	return db ,nil
}

