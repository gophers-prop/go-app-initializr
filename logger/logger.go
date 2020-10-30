package logger

import (
	"os"
  "github.com/sirupsen/logrus"
)
func init(){

  logrus.SetFormatter(&logrus.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  logrus.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  logrus.SetLevel(logrus.DebugLevel)
}

//Log returns a logging entry that will have a label appName
func Log(appName string) (*logrus.Entry){
	return logrus.WithField("go-initializer-app-name",appName)
}