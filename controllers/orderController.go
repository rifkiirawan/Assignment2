package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     uint
}

type Order struct {
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
}

func ShowOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "Data Fetched!",
		"Order":  []Order{},
	})
}

func NewOrder(ctx *gin.Context) {
	db := database.GetDB()
	var newOrder models.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder.OrderedAt = time.Now()

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Creating Order Data: %v", err),
		})
		return
	}

	fmt.Println("New product data:", newOrder)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "Data successfully added!",
		"Data":   newOrder,
	})
}

func DeleteOrder(ctx *gin.Context) {
	var (
		orderId string
	)
	db := database.GetDB()
	orderId = ctx.Param("orderId")

	if err := db.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.First(&models.Item{}, "order_id = ?", orderId).Error; err == nil {
		if err := db.Where("order_id = ?", orderId).Delete(&models.Item{}).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"result": fmt.Sprintf("Error Deleting Item: %v", err.Error()),
			})
			return
		}
	}

	if err := db.Where("order_id = ?", orderId).Delete(&models.Order{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Order: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %v Has Been Successfully Deleted", orderId),
		"success": true,
	})
}

func UpdateOrder(ctx *gin.Context) {
	var (
		orderId      string
		updatedOrder models.Order
		countItem    int
	)
	db := database.GetDB()
	orderId = ctx.Param("orderId")

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	if err := db.Model(&models.Order{}).Where("order_id = ?", orderId).Updates(updatedOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Updating Order Data: %v", err.Error()),
		})
		return
	}

	countItem = 0
	for _, item := range updatedOrder.Items {
		if err := db.Model(&models.Item{}).Where("item_id = ?", item.ItemId).Updates(item).Error; err == nil {
			countItem++
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Order":   updatedOrder,
		"message": fmt.Sprintf("Order with id %v Has Been Successfully Updated (%v Items Updated)", orderId, countItem),
		"success": true,
	})
}

func GetOrder(ctx *gin.Context) {
	var (
		orders []models.Order
	)

	db := database.GetDB()

	if err := db.Model(&models.Order{}).Preload("Items").Order("order_id asc").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Getting Order Data: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": orders,
	})
}
