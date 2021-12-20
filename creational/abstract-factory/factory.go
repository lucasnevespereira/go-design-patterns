package main

import (
	"abs-factory/builders"
	"abs-factory/models"
	"fmt"
)

type ISportsFactory interface {
	MakeShoe() models.IShoe
	MakeShort() models.IShort
}

func getSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &builders.Adidas{}, nil
	}

	if brand == "nike" {
		return &builders.Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
