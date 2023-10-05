package main

type Cat struct {
	PlushToy
}

func newCat() IPlushToy {
	return &Cat{
		PlushToy: PlushToy{
			name:   "Cat plush toy",
			colour: "Black",
		},
	}
}
