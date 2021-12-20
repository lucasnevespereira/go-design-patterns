package builders

import "abs-factory/models"

type Adidas struct{}

func (a *Adidas) MakeShoe() models.IShoe {
	return &models.AdidasShoe{
		models.Shoe{
			"adidas",
			42,
		},
	}
}

func (a *Adidas) MakeShort() models.IShort {
	return &models.AdidasShort{
		models.Short{
			Logo: "adidas",
			Size: 14,
		},
	}
}
