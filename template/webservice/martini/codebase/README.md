#GINREST -GOLANG Restful Starter Kit

[![Build Status](https://travis-ci.org/ribice/gorsk.svg?branch=master)](https://travis-ci.org/ribice/gorsk)
[![codecov](https://codecov.io/gh/ribice/gorsk/branch/master/graph/badge.svg)](https://codecov.io/gh/ribice/gorsk)
[![Go Report Card](https://goreportcard.com/badge/github.com/ribice/gorsk)](https://goreportcard.com/report/github.com/ribice/gorsk)
[![Maintainability](https://api.codeclimate.com/v1/badges/c3cb09dbc0bc43186464/maintainability)](https://codeclimate.com/github/ribice/gorsk/maintainability)

##About

GinRest is a starter kit for developing RESTful services in GOLANG . It provides boilerplate code with some examples . It is designed for developer's to save some time and directly implement business logic and dont need to start from scratch .    

GinRest follows golang principles and standard insipired by several package designs like Ben Johnson's [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1).The idea here is to provide a raw skeleton for an application.

This boilerplate code includes following functionalities:
* RESTful endpoint for login , password change and CRUD operations on the user.
* Application configuration from a config file (yaml and json )
* Logging from sirupsen/logrus +(some details on this loggin package)
* API docs using SwaggerUI
* With Complete testing coverage like unit test.
* Docker support to provide docker image.
* Docker-swarm / kubernetes support
* Make file.
* Package management using go mod (link of go mod).
* Multiple server's in a single microservice.
* Highly compatible with microservice based architecture.


##Getting Started

Setting up project
* Install any [Go lang](https://golang.org/doc/install) version and create a GOPATH.
* Create a src folder inside GOPATH and copy the extracted zip/tar file inside src. The path for you application should be like $GOPATH/src/ginRest 
* Run app using
```bash 
go run cmd/main.go
```

Setting up database

* Install mysql or have a dburl , dbname , username , password and tables ready .
* Update database details in database/db.go.

The application runs as an HTTP server . Check localhost:9090/swagger to check swagger docs


##Package Definations

*  "api"
  * This package has server configurations , routes , handler functions and swagger server .
*  "pkg/database"
  * This package has database configurations.
*  "pkg/user"
  * All the user ralated task like Create , Update , Read and Delete logic.
*  "config"
  * Reads config data from JSON/YAML file and provide functions to access contents of a file.
*  "consts"
  * Global constants that are used accross application.
*  "environment"
  * Reads environment variable from machine
*  "utils"
  * Common funcs that can be used accross application.
*  "models"
  * Collections of struct which are common accross multiple packages
*  "logger"
  * logging configurations like setting logging syntax , logging level. 
*  "pkg"
  * This package is highly important as this is the core of application



##LICENSE

ginRest is licensed under the MIT license. Check the [LICENSE](LICENSE) file for details.

##Author

[golangapps](http://golangapps.com)

##Credits

The following dependencies are used in this project .

```bash
|-------------------------------------|--------------------------------------------|--------------|
|             DEPENDENCY              |                  REPOURL                   |   LICENSE    |
|-------------------------------------|--------------------------------------------|--------------|
| github.com/gin-contrib/static       | github.com/gin-contrib/static              | MIT          |
| github.com/gin-gonic/gin            | github.com/gin-gonic/gin                   | bsd-2-clause |
| github.com/go-sql-driver/mysql      | github.com/go-sql-driver/mysql             | MIT          |
| github.com/sirupsen/logrus          | https://github.com/sirupsen/logrus         | MIT          |
| gopkg.in/yaml.v2                    | https://github.com/go-yaml/yaml            |              |
| github.com/lib/pq                   | https://github.com/lib/pq                  | Other        |
|-------------------------------------|--------------------------------------------|--------------|
```

##Directory Structure

GINREST
├── api
│   ├── handlers.go
│   ├── routes.go
│   ├── server.go
│   └── swagger.go
├── assets
│   └── swagger
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   ├── config.json
│   └── config.yaml
├── consts
│   └── consts.go
├── deployment
│   └── deployment.yaml
├── dockerfile
├── environment
│   └── environment.go
├── go.mod
├── go.sum
├── LICENCE.md
├── logger
│   └── logger.go
├── models
│   └── user.go
├── pkg
│   ├── database
│   │   └── db.go
│   └── user
│       ├── query.go
│       └── user.go
├── README.md
├── scripts
└── utils


** things to improve
//postgres db connection gets created for every db call this can be optimized to make it by creating single db connection for every rest request