module {{ .AppName }}

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	{{ .Logging.ImportPath }} {{ .Logging.Version }}
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79
	gopkg.in/yaml.v2 v2.2.2
)
