package controllers

import (
	"net/http"
	"server/helpers"
	"server/models"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func AddOrder(context *gin.Context) {
	var stocks []models.Stock
	var products []models.Product
	var productInput models.ProductInput
	var stockIDs []string
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	orderInput := models.OrderInputs{}
	if err := context.ShouldBindJSON(&orderInput); err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errors": middlewares.CompileErrors(err),
				"message": "Required fields are essential",
			},
		)
		return
	}

	var totalQuantity int
	var stock models.Stock
	orderedProducts := make(map[string]models.ProductInput)
	if len(orderInput.Products) >= 1 {
		for _, product := range orderInput.Products {
			stockIDs = append(stockIDs, (product.StockID).String())
			var productInput = models.ProductInput{
				Quantity: product.Quantity,
				StockID: product.StockID,
			}
			orderedProducts[(product.StockID).String()] = productInput
		}
		stocks = stock.FindIn(stockIDs)

		if len(stocks) >= 1  {
			for _, stock := range stocks {
				productInput = orderedProducts[(stock.ID).String()]
				product := models.Product{
					Quantity: productInput.Quantity,
					StockID: stock.ID,
				}
				products = append(products, product)
				totalQuantity += productInput.Quantity
			}
		}
		order := models.Order{
			TotalQuantity: totalQuantity,
			Products: products,
			ShippingAddress: orderInput.ShippingAddress,
			UserID: user.ID,
		}
		_, err := order.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
			return
		}
		context.JSON(http.StatusCreated, gin.H{ "data":  "Order created successfully." })
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
