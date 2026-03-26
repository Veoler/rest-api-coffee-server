package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Veoler/rest-api-coffee-server/models"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Сообщение": "Отклик есть",
		})
	})

	r.GET("/drinks", func(c *gin.Context) {
		drinks, err := models.GetDrinksShort()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, drinks)
	});

	r.Run()
}