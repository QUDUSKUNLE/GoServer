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
	var stocks []models.Stock
	var products []models.Product
	var productInput models.ProductInput
	var stockIDs []string
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	jsonData, _ := io.ReadAll(context.Request.Body)
	var orderInput models.OrderInputs
	if err := json.Unmarshal(jsonData, &orderInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
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
