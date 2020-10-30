[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gophers-prop/golang-appgen/blob/master/LICENSE)

# Golang Appsgen

A application that generates Golang application code based on user inputs. It generates different types of application like CLI, Web Serverices, Web client etc using librarey of your choice. While generating sample application, it follows proper directory structure and coding best practices. Start using the hosted version of this application running at [golangapps.com](http://golangapps.com).


### Contributing

Golang Appsgen is an open source project run by volunteers, and contributions are welcome! Check out the [Issues](https://github.com/gophers-prop/golang-aappgen/issues) page to see if your idea for a contribution has already been mentioned, and feel free to raise an issue or submit a pull request.

### Credits
The Go gopher on the home page was designed by Renée French.

Credits of all packages used are given in respective application template README.md file. Also these files are part of downloaded output folder.

### License

The code is licensed under the permissive Apache v2.0 licence. This means you can do what you like with the software, as long as you include the required notices. [Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for a summary.


###Setup

1) Install go > 1.12
2) git clone https://github.com/gophers-prop/golang-appgen.git --do not clone inside GOPATH
3) cd golang-appgen
4) mkdir output
5) mkdir outputzip
4) go mod download
6) go run cmd/main.go