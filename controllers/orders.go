package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"server/helpers"
	"server/models"
	"github.com/gin-gonic/gin"
)

func AddOrder(context *gin.Context) {
	var stock models.Stock
	var stocks []*models.Stock
	var stockIDs []string
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	jsonData, _ := io.ReadAll(context.Request.Body)
	var orderInput models.OrderInput
	if err := json.Unmarshal(jsonData, &orderInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
	}
	if len(orderInput.Stocks) >= 1 {
		for _, stockID := range orderInput.Stocks {
			stockIDs = append(stockIDs, stockID.Stock)
		}
		stocks = stock.FindIn(stockIDs)
		order := models.Order{
			Quantity: len(orderInput.Stocks),
			Stocks: stocks,
			UserID: user.ID,
		}
		savedOrder, err := order.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
			return
		}
		context.JSON(http.StatusCreated, gin.H{ "data":  savedOrder })
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{ "message": "No order made" })
}

func GetOrders(context *gin.Context) {
	var order models.Order
	result := order.FindAll()
	context.JSON(http.StatusOK, gin.H{ "data": result })
}

func GetOrder(context *gin.Context) {
	var order models.Order
	result, err := order.FindOrderByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "Record not found", })
		return
	}
	context.JSON(http.StatusOK, gin.H{ "data": result })
}
