package controllers

import (
	"net/http"
	"server/helpers"
	"server/middlewares"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddOrder(context *gin.Context) {
	// var stocks []models.Stock
	products := []models.Product{}
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	orderInput := models.OrderInputs{}
	if err := context.ShouldBindJSON(&orderInput); err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errors":  middlewares.CompileErrors(err),
				"message": "Required fields are essential",
			},
		)
		return
	}

	if (orderInput.AddressID).String() == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "order address id is required"})
		return
	}
	userProfile, err := helpers.UserProfile(user.ID.String())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	totalQuantity := 0
	stock := models.Stock{}
	if len(orderInput.Products) >= 1 {
		for _, product := range orderInput.Products {
			stockID := product.StockID
			product := models.Product{
				Quantity: product.Quantity,
				StockID:  product.StockID,
			}
			products = append(products, product)
			totalQuantity += product.Quantity
			stock.UpdateSlot(models.UpdateSlotInput{StockID: stockID, Slot: product.Quantity})
		}
		order := models.Order{
			TotalQuantity: totalQuantity,
			Products:      products,
			AddressID:     orderInput.AddressID,
			ProfileID:     userProfile.ID,
		}
		_, err := order.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Error making order"})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"data": "Order created successfully."})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"message": "No order made"})
}

func GetOrders(context *gin.Context) {
	order := models.Order{}
	result := order.FindAll()
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func GetOrder(context *gin.Context) {
	order := models.Order{}
	result, err := order.FindOrderByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": result})
}

func PatchOrder(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Patch an Order"})
}

func DeleteOrder(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Delete a Order"})
}
