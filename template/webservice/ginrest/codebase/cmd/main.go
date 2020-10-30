package main

import (
	"os"
	"os/signal"
	"syscall"
	"{{ .AppName }}/api"
	"fmt"
	"{{ .AppName }}/environment"
	
)

//gin-gonic 
func main(){

	cleanUp := make(chan int , 1 )

	server := api.Create(cleanUp)
	//starting server
	go server.Start(environment.GetServerPort())

	//starting swagger server
	go api.SwaggerServer()
	
	//stopChan is use in case of server is closed intentionally or due to any server failure
	stopChan := make(chan os.Signal ,2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	{	
		fmt.Println("exiting due to signal" , <-stopChan)
	}
}