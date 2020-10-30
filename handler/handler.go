package handler

import (
	"go-initializer/pkg/engine"

	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Test : test function ...Must be removed
func Test(ctx *gin.Context) {
	var request engine.GenerateTemplateRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(request)
}

//Liveness sds
func Liveness(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "liveness", "active": "true"})
	ctx.Abort()
	return
}

//GenerateTemplate Create a zip file of a template code
func GenerateTemplate(ctx *gin.Context) {
	engine.GenerateTemplate(ctx)
}
