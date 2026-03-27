package models

type Drink struct {
	ID			int		`json:"id"`
	Name		*string `json:"name"`
	Price		*int	`json:"price"`
	InStock		*bool	`json:"inStock"`
	IsCaffeine	*bool	`json:"isCaffeine"`
	Volume		*int	`json:"volume"`
	Description	*string	`json:"description"`
}

type DrinkShort struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Price	int		`json:"price"`
}
