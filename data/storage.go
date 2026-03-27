package data

import (
	"errors"
	"github.com/Veoler/rest-api-coffee-server/models"
)

// Вспомогательные функции добавления значений в структуры с указателями
func strPtr(s string) *string { 
	return &s 
}
func intPtr(i int) *int { 
	return &i 
}
func boolPtr(b bool) *bool { 
	return &b 
}

var drinks = []models.Drink{
	{
		ID:				1, 
		Name:			strPtr("Эспрессо"), 
		Price:			intPtr(150), 
		InStock:		boolPtr(true), 
		IsCaffeine:		boolPtr(true), 
		Volume:			intPtr(30), 
		Description:	strPtr("Крепкий кофе"),
	},
	{
		ID:				2,
		Name:			strPtr("Черный чай"),
		Price:			intPtr(100),
		InStock:		boolPtr(true),
		IsCaffeine:		boolPtr(false),
		Volume:			intPtr(300),
		Description:	strPtr("Чай с пакетиком"),
	},
	{
		ID:				3,
		Name:			strPtr("Nute Cream"),
		Price:			intPtr(370),
		InStock:		boolPtr(false),
		IsCaffeine:		boolPtr(true),
		Volume:			intPtr(400),
		Description:	strPtr("Очень вкусный холодный кофе с Урбечом"),
	},
}

var nextID = 4

func GetAll() []models.Drink { 
	return drinks 
}

func GetByID(id int) (models.Drink, error) {
	for _, d := range drinks {
		if d.ID == id { 
			return d, nil 
		}
	}
	return models.Drink{}, errors.New("напиток не найден")
}

func Create(d models.Drink) models.Drink {
	d.ID = nextID
	nextID++
	drinks = append(drinks, d)
	return d
}

func Delete(id int) error {
	for i, d := range drinks {
		if d.ID == id {
			drinks = append(drinks[:i], drinks[i+1:]...)
			return nil
		}
	}
	return errors.New("id не найден")
}

func Update(id int, input models.Drink) (models.Drink, error) {
	for i := range drinks {
		if drinks[i].ID == id {
			// Обновляем только те поля, которые пришли в JSON (не nil)
			if input.Name != nil { drinks[i].Name = input.Name }
			if input.Price != nil { drinks[i].Price = input.Price }
			if input.InStock != nil { drinks[i].InStock = input.InStock }
			if input.IsCaffeine != nil { drinks[i].IsCaffeine = input.IsCaffeine }
			if input.Volume != nil { drinks[i].Volume = input.Volume }
			if input.Description != nil { drinks[i].Description = input.Description }
			return drinks[i], nil
		}
	}
	return models.Drink{}, errors.New("напиток не найден")
}