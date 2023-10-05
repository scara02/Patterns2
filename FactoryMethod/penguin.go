package main

type Penguin struct {
	PlushToy
}

func newPenguin() IPlushToy {
	return &Penguin{
		PlushToy: PlushToy{
			name:   "Penguin plush toy",
			colour: "Pink and White",
		},
	}
}
