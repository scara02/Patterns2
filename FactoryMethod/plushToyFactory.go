package main

import "fmt"

func getPlushToy(toyType string) (IPlushToy, error) {
	if toyType == "cat" {
		return newCat(), nil
	}
	if toyType == "penguin" {
		return newPenguin(), nil
	}
	return nil, fmt.Errorf("wrong toy type")
}
