package main

import (
	"abs-factory/models"
	"fmt"
)

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s models.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShortDetails(s models.IShort) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
