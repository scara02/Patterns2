package main

type PlushToy struct {
	name   string
	colour string
}

func (p *PlushToy) setName(name string) {
	p.name = name
}

func (p *PlushToy) setColour(colour string) {
	p.colour = colour
}

func (p *PlushToy) getName() string {
	return p.name
}

func (p *PlushToy) getColour() string {
	return p.colour
}
