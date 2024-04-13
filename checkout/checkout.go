package checkout

import (
	"errors"
	"strings"
)

type StoreChecker interface {
	Get(string) (StoreItem, error)
}

type StoreItem struct {
	sku          string
	value        int
	specialOffer *SpecialOffer
}

type SpecialOffer struct {
	threshold            int
	thresholdAmountValue int
}

type Checkout struct {
	basket map[string]int
	store  StoreChecker
}

func NewCheckout(store StoreChecker) *Checkout {
	emptyBasket := make(map[string]int)
	return &Checkout{basket: emptyBasket, store: store}
}

func (c Checkout) Scan(s string) error {
	if len(s) != 1 {
		return errors.New("input is too long")
	}
	upperCaseInput := strings.ToUpper(s)

	_, err := c.store.Get(upperCaseInput)
	if err != nil {
		return err
	}

	count, ok := c.basket[upperCaseInput]
	if !ok {
		c.basket[upperCaseInput] = 1
		return nil
	}

	c.basket[upperCaseInput] = count + 1
	return nil
}
