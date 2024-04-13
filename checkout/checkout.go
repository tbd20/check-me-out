package checkout

type Checkout struct {
	basket map[string]int
}

func NewCheckout() *Checkout {
	emptyBasket := make(map[string]int)
	return &Checkout{basket: emptyBasket}
}

func (c Checkout) Scan(s string) error {
	count, ok := c.basket[s]
	if !ok {
		c.basket[s] = 1
		return nil
	}

	c.basket[s] = count + 1
	return nil
}
