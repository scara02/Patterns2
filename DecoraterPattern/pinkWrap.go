package main

type PinkWrap struct {
	gift IGift
}

func (c *PinkWrap) getPrice() int {
	price := c.gift.getPrice()
	return price + 8
}
