package models

import (
	"github.com/Veoler/rest-api-coffee-server/data"
)

// GetDrinksShort конвертирует список полных структур в сокращенные
func GetDrinksShort() ([]DrinkShort, error) {
    data := data.GetAll()

    var result []DrinkShort
    for _, r := range data {
        result = append(result, DrinkShort{
            ID:    r.ID,
            Name:  r.Name,
            Price: r.Price,
        })
    }
    return result, nil
}