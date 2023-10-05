package main

type WhiteRibbon struct {
	gift IGift
}

func (c *WhiteRibbon) getPrice() int {
	price := c.gift.getPrice()
	return price + 5
}
