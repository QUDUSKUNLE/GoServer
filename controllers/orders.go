package controllers

import (
	"net/http"
	"server/helpers"
	"server/middlewares"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddOrder(context *gin.Context) {
	productsModel := []models.Product{}

	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	orderInputModel := models.OrderInputs{}
	if err := context.ShouldBindJSON(&orderInputModel); err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errors":  middlewares.CompileErrors(err),
				"message": "Required fields are essential",
			},
		)
		return
	}

	if (orderInputModel.AddressID).String() == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "order address id is required"})
		return
	}

	userProfile, err := helpers.UserProfile(user.ID.String())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": middlewares.CompileErrors(err)})
		return
	}

	totalQuantity := 0
	stockModel := models.Stock{}
	if len(orderInputModel.Products) >= 1 {
		for _, product := range orderInputModel.Products {
			stockID := product.StockID
			productModel := models.Product{
				Quantity: product.Quantity,
				StockID:  product.StockID,
			}
			productsModel = append(productsModel, productModel)
			totalQuantity += product.Quantity
			stockModel.UpdateSlot(models.UpdateSlotInput{StockID: stockID, Slot: product.Quantity})
		}
		orderModel := models.Order{
			TotalQuantity: totalQuantity,
			Products:      productsModel,
			AddressID:     orderInputModel.AddressID,
			ProfileID:     userProfile.ID,
		}

		if err := orderModel.Save(); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Error making order"})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"data": "Order created successfully."})
		return
	}

	context.JSON(http.StatusBadRequest, gin.H{"message": "No order made"})
}

func GetOrders(context *gin.Context) {
	orderModel := models.Order{}
	orders := orderModel.FindAll()
	context.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetOrder(context *gin.Context) {
	orderModel := models.Order{}

	order, err := orderModel.FindOrderByID(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": order})
}

func PatchOrder(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Patch an Order"})
}

func DeleteOrder(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Delete a Order"})
}
