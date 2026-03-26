package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Veoler/rest-api-coffee-server/service"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Сообщение": "Отклик есть",
		})
	})

	r.GET("/drinks", func(c *gin.Context) {
		drinks, err := service.GetDrinksShort()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, drinks)
	});

	r.GET("/drinks/in-stock", func(c *gin.Context) {
		drinks, err := service.GetInStockShort()
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, drinks)
	})

	r.GET("/drinks/:id", func(c *gin.Context) {
		id := c.Param("id")
		idGet, err := strconv.Atoi(id)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат ID"})
			return
		}

		// Вызываем сервис
		drink, err := service.GetDrinkFull(idGet)
		if err != nil {
			// Если напиток не найден (репозиторий вернул ошибку)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, drink)
	})

	r.Run()
}