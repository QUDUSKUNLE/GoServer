package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddStock(context *gin.Context) {
	stockInputModel := models.CreateStockInput{}
	if err := context.ShouldBindJSON(&stockInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if stockInputModel.Price == 0 && stockInputModel.Cost > 0 && stockInputModel.Unit > 0 {
		stockInputModel.Price = (stockInputModel.Cost / float32(stockInputModel.Unit))
	}

	stockModel := models.Stock{
		Type:        stockInputModel.Type,
		Description: stockInputModel.Description,
		Cost:        stockInputModel.Cost,
		Province:    stockInputModel.Province,
		Price:       stockInputModel.Price,
		Unit:        stockInputModel.Unit,
		Slot:        stockInputModel.Slot,
	}

	if err := stockModel.Save(); err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": stockModel})
}

func GetStocks(context *gin.Context) {
	stockModel := models.Stock{}
	stocks := stockModel.FindAll()
	context.JSON(http.StatusOK, gin.H{"data": stocks})
}

func GetStock(context *gin.Context) {
	stockModel := models.Stock{}
	result, err := stockModel.FindStockByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateStock(context *gin.Context) {
	updateStockInputModel := models.UpdateStockInput{}
	if err := context.ShouldBindJSON(&updateStockInputModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateStockInputModel.Price == 0 && updateStockInputModel.Cost > 0 && updateStockInputModel.Unit > 0 {
		updateStockInputModel.Price = (updateStockInputModel.Cost / float32(updateStockInputModel.Unit))
	}

	stockModel := models.Stock{}
	updatedStock, err := stockModel.Update(updateStockInputModel, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedStock})
}

func DeleteStock(context *gin.Context) {
	stockModel := models.Stock{}
	if _, err := stockModel.Delete(context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{"data": "Stock deleted successfully"})
}
