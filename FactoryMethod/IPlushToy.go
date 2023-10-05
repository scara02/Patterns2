package main

type IPlushToy interface {
	setName(name string)
	setColour(colour string)
	getName() string
	getColour() string
}
