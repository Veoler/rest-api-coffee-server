package data

import (
	"errors"
	"github.com/Veoler/rest-api-coffee-server/models"

)

var Drinks = []models.Drink{
	{
		ID:				1,
		Name:			"Эспрессо",
		Price:			150,
		InStock:		true,
		IsCaffeine:		true,
		Volume:			30,
		Description:	"Классический крепкий кофе",
	},
	{
		ID:				2,
		Name:			"Черный чай",
		Price:			100,
		InStock:		true,
		IsCaffeine:		false,
		Volume:			300,
		Description:	"Чай с пакетиком",
	},
	{
		ID:				3,
		Name:			"Nute Cream",
		Price:			370,
		InStock:		false,
		IsCaffeine:		true,
		Volume:			400,
		Description:	"Очень вкусный холодный кофе с Урбечом",
	},
}

func GetAll() []models.Drink {
	return Drinks
}

func GetByID(id int) (models.Drink, error) {
	for _, r := range Drinks {
		if r.ID == id {
			return r, nil
		}
	}
	return models.Drink{}, errors.New("напиток не найден")
}

func GetInStock() []models.Drink {
	var inStock []models.Drink
	for _, r := range Drinks {
		if r.InStock {
			inStock = append(inStock, r)
		}
	}
	return inStock
}

func Create(newDrink models.Drink) models.Drink {
	NextID := 4
	newDrink.ID = NextID
	NextID++
	Drinks = append(Drinks, newDrink)
	return newDrink
}

func Delete(id int) error {
	for i, r := range Drinks {
		if r.ID == id {
			Drinks = append(Drinks[:i], Drinks[i+1:]...)
			return nil
		}
	}
	return errors.New("не удалось удалить: напиток не найден")
}

func Update(id int, updatedDrink models.Drink) (models.Drink, error) {
	for i, r := range Drinks {
		if r.ID == id {
			Drinks[i] = updatedDrink
			Drinks[i].ID = id
			return Drinks[i], nil
		}
	}
	return models.Drink{}, errors.New("не удалось обновить: напиток не найден")
}