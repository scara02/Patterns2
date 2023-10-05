package main

import "fmt"

func main() {
	gift := &DollGift{}

	giftWithWrap := &PinkWrap{gift: gift}

	giftWithWrapAndRibbon := &WhiteRibbon{gift: giftWithWrap}

	fmt.Printf("Price of doll gift with pink wrap and white ribbon is %d$ \n", giftWithWrapAndRibbon.getPrice())
}
