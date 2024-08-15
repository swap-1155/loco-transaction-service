package api

import (
	"loco/domain"
	"loco/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetTransactionById(c *gin.Context) {

	transaction_id := c.Param("transaction_id")
	if transaction_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "transaction_id not found",
			"status":  http.StatusBadRequest,
		})
		return
	}

	request := &domain.Transaction{}
	request.TransactionId = cast.ToInt(transaction_id)

	response, err := service.GetTransactionById(request)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func PostTransactionById(c *gin.Context) {

	transaction_id := c.Param("transaction_id")
	if transaction_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "transaction_id not found",
			"status":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	request := &domain.Transaction{}
	request.TransactionId = cast.ToInt(transaction_id)

	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	response, err := service.PostTransactionById(request)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetTransactionsByType(c *gin.Context) {

	transaction_type := c.Param("type")
	if transaction_type == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "transaction_type not found",
			"status":  http.StatusBadRequest,
		})
		return
	}

	request := &domain.Transaction{}
	request.TransactionType = cast.ToString(transaction_type)

	response, err := service.GetTransactionsByType(request)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func GetTotalTransactionAmount(c *gin.Context) {

	transaction_id := c.Param("transaction_id")
	if transaction_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "transaction_id not found",
			"status":  http.StatusBadRequest,
		})
		return
	}

	request := &domain.Transaction{}
	request.TransactionId = cast.ToInt(transaction_id)

	response, err := service.GetTotalTransactionAmount(request)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}
