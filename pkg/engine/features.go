package engine

import(
	"go-initializer/utils"
	"path/filepath"
	"io/ioutil"
	"os"
	"fmt"
	"go-initializer/logger"
)
var featureMap map[string]featureFunc 

type Functionality struct{
	Name string   	`json:"name,omitempty"`
	Library string	`json:"library,omitempty"`
	OutputFolder string  
}

const(
	LOGGING = "logging"
	SWAGGER = "swagger"
	DATABASE = "database"
	D_CTX = "d-ctx"
)

type featureFunc func(Functionality)

func init(){
	featureMap = map[string]featureFunc{
		LOGGING : injectLogging,
		SWAGGER : injectSwagger,
		DATABASE: injectDatabase,

	}
}

func injectLogging(f Functionality){
	homePath := utils.GetWorkingDirNoError()

	goFilePath := filepath.Join(homePath, "template", "logframework", f.Library, "logger.go")

	goFileContent, err := ioutil.ReadFile(goFilePath)
	if err != nil {
		fmt.Printf("Error getting log file content for framework %s with error %s", f.Library, err.Error())
	}
	if _, err := os.Stat(filepath.Join(f.OutputFolder, "logger")); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(f.OutputFolder, "logger"), os.ModePerm)
		if err != nil {
			fmt.Printf("Cannot create logging directory for request %+v with error %s", f, err.Error())
		}
	}
	file, err := os.Create(filepath.Join(f.OutputFolder, "logger/logger.go"))
	if err != nil {
		 fmt.Printf("Cannot create logging file inside output folder for request %+v with error %s", f, err.Error())
	}
	defer file.Close()
	_, err = file.Write(goFileContent)
	if err != nil {
		 fmt.Printf("Error writing contents to output logging file for request %+v with error %s", f, err.Error())
	}
}

func injectDatabase(f Functionality){
	logger.Log("DUMMY").Debug("Inject database files %+v",f)
	//check if database folder exist in the output folder
	
	if _,err := os.Stat(filepath.Join(f.OutputFolder,"pkg","database"));os.IsNotExist(err){
		err = os.MkdirAll(filepath.Join(f.OutputFolder,"pkg","database"),os.ModePerm)
		if err != nil{
			logger.Log(D_CTX).Errorf("Error while creating pkg/database directory %+v",err)
			return 
		}
	}
	if _,err := os.Stat(filepath.Join(f.OutputFolder,"pkg","user"));os.IsNotExist(err){
		err = os.MkdirAll(filepath.Join(f.OutputFolder,"pkg","user"),os.ModePerm)
		if err != nil{
			logger.Log(D_CTX).Errorf("Error while creating pkg/user directory %+v",err)
			return 
		}
	}
	if _,err := os.Stat(filepath.Join(f.OutputFolder,"pkg","auth"));os.IsNotExist(err){
		err = os.MkdirAll(filepath.Join(f.OutputFolder,"pkg","auth"),os.ModePerm)
		if err != nil{
			logger.Log(D_CTX).Errorf("Error while creating pkg/auth directory %+v",err)
			return 
		}
	}
	homePath := utils.GetWorkingDirNoError()

	databaseGoFilePath := filepath.Join(homePath, "template", "databaseframework", f.Library, "database","db.go")
	databseGoFileContent, err := ioutil.ReadFile(databaseGoFilePath)
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while reading db.go file %+v",err)
		return
	}
	file, err := os.Create(filepath.Join(f.OutputFolder, "pkg","database","db.go"))
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while creating file db.go inside output folder %+v",err)
		return
	}
	defer file.Close()
	_, err = file.Write(databseGoFileContent)
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while writing file content for db.go inside output folder %+v",err)
		return
	}
	userQueryGoFilePath := filepath.Join(homePath, "template", "databaseframework", f.Library, "user","query.go")
	userQueryGoFileContent, err := ioutil.ReadFile(userQueryGoFilePath)
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while reading user/query.go file %+v",err)
		return
	}
	userQueryfile, err := os.Create(filepath.Join(f.OutputFolder, "pkg","user","query.go"))
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while creating file query.go inside output folder %+v",err)
		return
	}
	defer userQueryfile.Close()
	_, err = userQueryfile.Write(userQueryGoFileContent)
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while writing file content for query.go inside output folder %+v",err)
		return
	}

	authQueryGoFilePath := filepath.Join(homePath, "template", "databaseframework", f.Library, "auth","query.go")
	authQueryGoFileContent, err := ioutil.ReadFile(authQueryGoFilePath)
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while reading auth/query.go file %+v",err)
		return
	}
	authQueryfile, err := os.Create(filepath.Join(f.OutputFolder, "pkg","auth","query.go"))
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while creating file query.go inside output folder %+v",err)
		return
	}
	defer authQueryfile.Close()
	_, err = authQueryfile.Write(authQueryGoFileContent)
	if err != nil {
		logger.Log(D_CTX).Errorf("Error while writing file content for query.go inside output folder %+v",err)
		return
	}
}

func injectSwagger(f Functionality){
    
}