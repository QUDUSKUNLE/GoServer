package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/helpers"
	"server/models"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	// "github.com/satori/go.uuid"
)

func AddOrder(context *gin.Context) {
	var stock models.Stock
	var orderMadeInput models.OrderMadeInput
	var stocks []models.Stock
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	orderMadeInput.UserID = user.ID
	jsonData, _ := io.ReadAll(context.Request.Body)
	var orderInput models.OrderInput
	if err := json.Unmarshal(jsonData, &orderInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
	}

	for _, stoc := range orderInput.Stocks {
		orderStock, err := stock.FindStockBy(stoc.Stock)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(orderStock)
		stocks = append(stocks, orderStock)
	}
	orderMadeInput.Quantity = len(orderInput.Stocks)
	orderMadeInput.Stocks = stocks

	order := models.Order{
		Quantity: orderMadeInput.Quantity,
		Stocks: orderMadeInput.Stocks,
		UserID: orderMadeInput.UserID,
	}

	savedOrder, err := order.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	context.JSON(http.StatusCreated, gin.H{ "data":  savedOrder })
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
