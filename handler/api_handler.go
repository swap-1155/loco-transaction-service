package handler

import (
	"fmt"
	"loco/handler/api"
	"loco/util/middleware"

	"github.com/gin-gonic/gin"
)

// SetupApiRoutes ... setup all the API routes
func SetupApiRoutes(engine *gin.Engine) {
	fmt.Println("SetupRoutes")

	engine.Use(middleware.DBConnectionMiddleware)

	TransactionRouter(engine)
	TestRouter(engine)
}

func TransactionRouter(engine *gin.Engine) {
	incomingRoutes := engine.Group("/transactionservice")

	incomingRoutes.POST("/transaction/:transaction_id", api.PostTransactionById)
	incomingRoutes.GET("/transaction/:transaction_id", api.GetTransactionById)
	incomingRoutes.GET("/types/:type", api.GetTransactionsByType)
	incomingRoutes.GET("/sum/:transaction_id", api.GetTotalTransactionAmount)
}

func TestRouter(engine *gin.Engine) {
	engine.GET("/", api.Testing)
}
