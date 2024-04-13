package checkout

import "strings"

type Checkout struct {
	basket map[string]int
}

func NewCheckout() *Checkout {
	emptyBasket := make(map[string]int)
	return &Checkout{basket: emptyBasket}
}

func (c Checkout) Scan(s string) error {
	upperCaseInput := strings.ToUpper(s)
	count, ok := c.basket[upperCaseInput]
	if !ok {
		c.basket[upperCaseInput] = 1
		return nil
	}

	c.basket[upperCaseInput] = count + 1
	return nil
}
