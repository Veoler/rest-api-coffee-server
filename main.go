package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/Veoler/rest-api-coffee-server/models"
	"github.com/Veoler/rest-api-coffee-server/data"
	"github.com/Veoler/rest-api-coffee-server/service"
)

func main() {
	r := gin.Default()

	r.GET("/drinks", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.GetList(false))
	})

	r.GET("/drinks/in-stock", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.GetList(true))
	})

	r.GET("/drinks/:id", func(c *gin.Context) {
		id, errr := strconv.Atoi(c.Param("id"))
		if errr != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
    		return
		}

		drink, err := data.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, drink)
	})

	r.POST("/drinks", func(c *gin.Context) {
		var input models.Drink
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при добавлении: некорректные данные"})
			return
		}
		c.JSON(http.StatusCreated, data.Create(input))
	})

	r.DELETE("/drinks/:id", func(c *gin.Context) {
		id, errr := strconv.Atoi(c.Param("id"))
		if errr != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
    		return
		}

		if err := data.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})

	r.PATCH("/drinks/:id", func(c *gin.Context) {
		id, errr := strconv.Atoi(c.Param("id"))
		if errr != nil {
    		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
    		return
		}

		var input models.Drink
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при изменении"})
			return
		}
		
		updated, err := data.Update(id, input)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updated)
	})

	r.Run(":8080")
}