package checkout

type Checkout struct {
	basket map[string]int
}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c Checkout) Scan(s string) error {
	return nil
}
