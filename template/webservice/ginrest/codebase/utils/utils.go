package utils

import (
	"io/ioutil"
	"os"
	"reflect"
)

//HasElem checks an array if element exists or not
func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

//ListDir will return list of directory inside path depth is 1
func ListDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var dirName []string
	for _, f := range files {
		if f.IsDir() {
			dirName = append(dirName, f.Name())
		}
	}
	return dirName, nil
}

//GetWorkingDir get current working directory
func GetWorkingDir() (string, error) {

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}

//GetWorkingDirNoError get current working directory
func GetWorkingDirNoError() string {

	path, _ := os.Getwd()

	return path
}
