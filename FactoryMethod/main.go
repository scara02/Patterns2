package main

import "fmt"

func main() {
	cat, _ := getPlushToy("cat")
	penguin, _ := getPlushToy("penguin")

	printDetails(cat)
	printDetails(penguin)
}

func printDetails(p IPlushToy) {
	fmt.Printf("Toy: %s\n", p.getName())
	fmt.Printf("Colour: %s\n", p.getColour())
}
