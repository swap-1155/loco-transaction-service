package main

import (
	"fmt"

	"loco/handler"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambda
var router *gin.Engine

func main() {
	godotenv.Load()
	const functionName = "main.main"
	fmt.Println(functionName)
	router = gin.Default()
	handler.SetupApiRoutes(router)

	router.Run("localhost:8080")
}
