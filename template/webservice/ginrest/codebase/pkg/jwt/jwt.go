package jwt

import (
	"errors"

	"{{ .AppName }}/consts"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	jwtAuthAlgo = "HS256"
	//set it to hours
	jwtTokenDuration = 7
	jwtKey           = "my_secret_key_which_is_not_s"
)

//GenerateToken generate token based on payload provided in arguement
func GenerateToken(payload jwt.MapClaims) (string, error) {
	authAlgo := jwt.GetSigningMethod(jwtAuthAlgo)
	expire := time.Now().Add(time.Hour * time.Duration(jwtTokenDuration))
	payload["exp"] = expire.Unix()
	token := jwt.NewWithClaims(authAlgo, payload)

	tokenString, tokenErr := token.SignedString([]byte(jwtKey))
	if tokenErr != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.ERROR_GENERATING_NEW_TOKEN}} {{end}}
		return "", tokenErr
	}
	return tokenString, nil
}

//GenerateAccountLoginToken it generates a token after user is logged / signed up
func GenerateAccountLoginToken(email string) (string, error) {
	claims := jwt.MapClaims{
		consts.JWT_EMAIL_KEY: email,
	}
	return GenerateToken(claims)
}

//ParseToken parses the token and returns a jwt.Token reference
func ParseToken(tok string) (*jwt.Token, error) {
	tkn, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(jwtAuthAlgo) != token.Method {
			return "", errors.New("Authorization algorithm is incorrect")
		}
		return []byte(jwtKey), nil
	})
	return tkn, err
}

//MWHandle reading authorization token and validating token .This func works as a middleware for validating jwt token  as a handler func .
func MWHandle(handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.EMPTY_TOKEN_ERROR}} {{end}}
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Empty token or token not present in header")})
			return
		}
		newToken, err := ValidateToken(token)
		if newToken == "" || err != nil {
			{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.INVALID_TOKEN}} {{end}}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		handle(c)
	}
}

//GetTokenFromReq returns token from header
func GetTokenFromReq(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return "", errors.New("Empty token value")
	}
	return token, nil
}

//ValidateToken check if the token is valid or not
func ValidateToken(token string) (string, error) {

	tkn, err := ParseToken(token)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.ERROR_PARSING_TOKEN}} {{end}}
		return "", err
	}
	if !tkn.Valid {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.INVALID_TOKEN}} {{end}}
		return "", errors.New("Invalid token")
	}
	return tkn.SignedString([]byte(jwtKey))
}

//GetValueFromAKeyInToken parses the token and return value of the key
func GetValueFromAKeyInToken(token, key string) (string, error) {
	tkn, err := ParseToken(token)
	if err != nil {
		{{ if .Logging.ImportPath }} {{ .Logging.Messages.Jwt.ERROR_PARSING_TOKEN}} {{end}}
		return "", err
	}
	claims := tkn.Claims.(jwt.MapClaims)

	if value, ok := claims[key].(string); ok {
		return value, nil
	}
	return "", errors.New("Key does not exist in token")
}
