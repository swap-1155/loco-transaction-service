package middleware

import (
	"fmt"
	"loco/util/connectToDB"

	"github.com/gin-gonic/gin"
)

func DBConnectionMiddleware(c *gin.Context) {
	const functionName = "middleware.DBConnectionMiddleware"
	fmt.Println(functionName, "Connecting to database...")
	connectToDB.ConnectToDatabase()
	c.Next()
}
