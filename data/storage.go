package storage

import (
	"github.com/Veoler/rest-api-coffee-server/models"
	"errors"
)

// Наше временное хранилище в памяти
var Drinks = []drink.Drink{
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

func GetAll() []drink.Drink {
	return Drinks
}

func GetByID(id int) (drink.Drink, error) {
	for _, r := range Drinks {
		if r.ID == id {
			return r, nil
		}
	}
	return drink.Drink{}, errors.New("напиток не найден")
}

func GetInStock() []drink.Drink {
	var inStock []drink.Drink
	for _, r := range Drinks {
		if r.InStock {
			inStock = append(inStock, r)
		}
	}
	return inStock
}

func Create(newDrink drink.Drink) drink.Drink {
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

func Update(id int, updatedDrink drink.Drink) (drink.Drink, error) {
	for i, r := range Drinks {
		if r.ID == id {
			Drinks[i] = updatedDrink
			Drinks[i].ID = id
			return Drinks[i], nil
		}
	}
	return drink.Drink{}, errors.New("не удалось обновить: напиток не найден")
}