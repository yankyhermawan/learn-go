package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yankyhermawan/learn-go/models"

	"github.com/yankyhermawan/learn-go/services"

	"net/http"
)

func main() {
	r := gin.Default()
	db, err := models.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	r.GET("/findone", func(c *gin.Context) {
		products, err := services.FindByID(db, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"products": products,
		})
	})

	data := &models.Product{
		ProductName: "kecap",
		Description: "manis",
	}

	r.POST("/post", func(c *gin.Context) {
		services.CreateUser(db, data)
	})
	r.Run()
}
