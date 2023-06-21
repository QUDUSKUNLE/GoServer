package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddStock(context *gin.Context) {
	var stockInput models.CreateStockInput
	if err := context.ShouldBindJSON(&stockInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if stockInput.Price == 0 && stockInput.Cost > 0 && stockInput.Unit > 0 {
		stockInput.Price = (stockInput.Cost / float32(stockInput.Unit))
	}

	stock := models.Stock{
		Type:        stockInput.Type,
		Description: stockInput.Description,
		Cost:        stockInput.Cost,
		Province:    stockInput.Province,
		Price:       stockInput.Price,
		Unit:        stockInput.Unit,
		Slot:        stockInput.Slot,
	}

	savedStock, err := stock.Save()
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
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

func UpdateStock(context *gin.Context) {
	var updateStockInput models.UpdateStockInput
	if err := context.ShouldBindJSON(&updateStockInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateStockInput.Price == 0 && updateStockInput.Cost > 0 && updateStockInput.Unit > 0 {
		updateStockInput.Price = (updateStockInput.Cost / float32(updateStockInput.Unit))
	}

	var stock models.Stock
	updatedStock, err := stock.Update(updateStockInput, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedStock})
}

func DeleteStock(context *gin.Context) {
	var stock models.Stock
	_, err := stock.Delete(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{"data": "Stock deleted successfully"})
}
