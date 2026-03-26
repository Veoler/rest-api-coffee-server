package service

import (
	"github.com/Veoler/rest-api-coffee-server/data"
	"github.com/Veoler/rest-api-coffee-server/models"
	"errors"
)

// GetDrinksShort конвертирует список полных структур в сокращенные
func GetDrinksShort() ([]models.DrinkShort, error) {
    data := data.GetAll()

    var result []models.DrinkShort
    for _, r := range data {
        result = append(result, models.DrinkShort{
            ID:    r.ID,
            Name:  r.Name,
            Price: r.Price,
        })
    }
    return result, nil
}

func GetInStockShort() ([]models.DrinkShort, error) {
	data := data.GetInStock() // берем только те, что InStock: true
	
	var result []models.DrinkShort
	for _, d := range data {
		result = append(result, models.DrinkShort{
			ID:    d.ID,
			Name:  d.Name,
			Price: d.Price,
		})
	}
	return result, nil
}

// GetDrinkFull возвращает полную информацию о напитке по ID
func GetDrinkFull(id int) (models.Drink, error) {
	// Просто пробрасываем запрос в репозиторий
	return data.GetByID(id)
}

func CreateDrink(input models.Drink) (models.Drink, error) {
	if input.Name == "" {
		return models.Drink{}, errors.New("название напитка обязательно")
	}
    
    if input.Price <= 0 {
        return models.Drink{}, errors.New("цена должна быть больше нуля")
    }

	return data.Create(input), nil
}