package types

import (
	"sync"
)
type Messages struct {
	General map[string]string `json:"GENERAL"`
	User    map[string]string `json:"USER"` 
	Database map[string]string `json:"DATABASE"`
	Jwt      map[string]string `json:"JWT"`
	Auth map[string]string `json:"AUTH"`
	Handler map[string]string `json:"HANDLER"`
	Config   map[string]string `json:"CONFIG"`
	Swagger  map[string]string  `json:"SWAGGER"`
	Server   map[string]string   `json:"SERVER"`
  }

type LoggingFramework struct {
	LibraryName string   `json:"LIBRARY_NAME"`
	ImportPath  string   `json:"IMPORT_PATH"`
	Version     string   `json:"VERSION"`
	Messages    Messages `json:"MESSAGES"`
}

type Configuration struct {
	AppName string
	Logging LoggingFramework
}

var Mutex = &sync.Mutex{}
