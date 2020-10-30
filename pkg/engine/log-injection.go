package engine

import (
	"encoding/json"
	"fmt"
	"go-initializer/types"
	"go-initializer/utils"
	"io/ioutil"
	"path/filepath"
)

func readLogJson(logFramework string) (*types.LoggingFramework, error) {
	homePath := utils.GetWorkingDirNoError()
	jsonFilePath := filepath.Join(homePath, "template", "logframework", logFramework, "logger.json")
	types.Mutex.Lock()
	
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	
	types.Mutex.Unlock()
	
	if err != nil {
		return nil, fmt.Errorf("Error while reading json log file for logging framework %s with error %s", logFramework, err.Error())
	}

	var frameworkData types.LoggingFramework

	err = json.Unmarshal([]byte(jsonData), &frameworkData)

	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling  json log file for logging framework %s with error %s", logFramework, err.Error())
	}
    
	return &frameworkData, nil

}

func readDatabaseJson(databaseFramework string)(*types.DatabaseFramework,error){
	homePath := utils.GetWorkingDirNoError()
	jsonFilePath := filepath.Join(homePath, "template", "databaseframework", databaseFramework, "metadata.json")
	types.Mutex.Lock()
	
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	
	types.Mutex.Unlock()
	
	if err != nil {
		return nil, fmt.Errorf("Error while reading json log file for database framework %s with error %s", databaseFramework, err.Error())
	}

	var frameworkData types.DatabaseFramework

	err = json.Unmarshal([]byte(jsonData), &frameworkData)

	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling  json log file for database framework %s with error %s", databaseFramework, err.Error())
	}
    
	return &frameworkData, nil
}
