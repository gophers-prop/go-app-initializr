//add port in environment variable

package environment

import (
	"{{ .AppName }}/consts"
	"os"
)

var(
	defaultServerPort = ":8080"
	defaultDatabaseUserName = "root"
	defaultDatabasePassword = "my-secret-password"
	defaultDatabaseName   = "users"
)


//GetServerPort get server port from environment
func GetServerPort()string{
	if val := os.Getenv(consts.SERVER_PORT_ENVIRONMENT_VARIABLE); val != ""{
		return val
	}
	return defaultServerPort
}


//GetDatabaseDetails get database name , database username , database password
func GetDatabaseDetails() map[string]string{
	dbDetails := make(map[string]string,0)
	if val := os.Getenv(consts.DATABASE_USER_NAME); val != ""{
		dbDetails[consts.DATABASE_USER_NAME] = val
	}else{
		dbDetails[consts.DATABASE_USER_NAME] = defaultDatabaseUserName
	}
	if val := os.Getenv(consts.DATABSE_USER_PASSWORD); val != ""{
		dbDetails[consts.DATABSE_USER_PASSWORD] = val
	}else{
		dbDetails[consts.DATABSE_USER_PASSWORD] = defaultDatabasePassword
	}
	if val := os.Getenv(consts.DATABASE_NAME); val != ""{
		dbDetails[consts.DATABASE_NAME] = val
	}else{
		dbDetails[consts.DATABASE_NAME] = defaultDatabaseName
	}
	return dbDetails
}