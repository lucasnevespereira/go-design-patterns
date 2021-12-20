package builders

import (
	"abs-factory/models"
)

type Nike struct{}

func (n *Nike) MakeShoe() models.IShoe {
	return &models.NikeShoe{
		models.Shoe{
			Logo: "nike",
			Size: 39,
		},
	}
}

func (n *Nike) MakeShort() models.IShort {
	return &models.NikeShort{
		models.Short{
			Logo: "nike",
			Size: 16,
		},
	}
}
