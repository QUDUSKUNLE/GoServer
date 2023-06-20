package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/models"
)

func AddStock(context *gin.Context) {
	var stockInput models.CreateStockInput
	if err := context.ShouldBindJSON(&stockInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	stock := models.Stock{
		Type:        stockInput.Type,
		Description: stockInput.Description,
		Cost:        stockInput.Cost,
		Province:    stockInput.Province,
		Price:       (stockInput.Cost / float32(stockInput.Unit)),
		Unit:        stockInput.Unit,
		Slot:        stockInput.Slot,
	}

	savedStock, err := stock.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedStock})
}

func GetStocks(context *gin.Context) {
	var stock models.Stock
	result := stock.FindAll()
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func GetStock(context *gin.Context) {
	var stock models.Stock
	result, err := stock.FindStockByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": result})
}
