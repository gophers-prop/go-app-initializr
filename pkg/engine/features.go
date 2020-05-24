package engine

import(
	"go-initializer/utils"
	"path/filepath"
	"io/ioutil"
	"os"
	"fmt"
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
)

type featureFunc func(Functionality)

func init(){
	featureMap = map[string]featureFunc{
		LOGGING : injectLogging,
		SWAGGER : injectSwagger,
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

func injectSwagger(f Functionality){

}