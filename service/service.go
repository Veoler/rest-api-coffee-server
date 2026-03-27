package service

import (
	"github.com/Veoler/rest-api-coffee-server/models"
	"github.com/Veoler/rest-api-coffee-server/data"
)

func GetList(onlyInStock bool) []models.DrinkShort {
	data := data.GetAll()
	var result []models.DrinkShort
	for _, d := range data {
		if onlyInStock && (d.InStock == nil || !*d.InStock) { continue }
		
		// Разыменовываем указатели для безопасного вывода
		name, price := "", 0
		if d.Name != nil { name = *d.Name }
		if d.Price != nil { price = *d.Price }

		result = append(result, models.DrinkShort{ID: d.ID, Name: name, Price: price})
	}
	return result
}