package checkout

import (
	"errors"
	"reflect"
	"testing"
)

type testScanStoreChecker struct{}

func (c testScanStoreChecker) Get(s string) (StoreItem, error) {
	if s == "X" {
		return StoreItem{}, errors.New("scanned item is not in the store")
	}

	return StoreItem{}, nil
}

type basicTestGetForTotalPrice struct{}

func (b basicTestGetForTotalPrice) Get(s string) (StoreItem, error) {
	outputStoreItem := StoreItem{
		sku:          s,
		value:        int(s[0]-'A') + 1,
		specialOffer: nil,
	}
	return outputStoreItem, nil
}

type errorTestGetTotalForPrice struct{}

func (b errorTestGetTotalForPrice) Get(s string) (StoreItem, error) {
	return StoreItem{}, errors.New("error in get")
}

type specialOfferTestGetTotalForPrice struct{}

func (store specialOfferTestGetTotalForPrice) Get(s string) (StoreItem, error) {
	charValue := int(s[0]-'A') + 1
	outputStoreItem := StoreItem{
		sku:   s,
		value: charValue,
		specialOffer: &SpecialOffer{
			threshold:            charValue * 100,
			thresholdAmountValue: charValue * 10,
		},
	}
	return outputStoreItem, nil
}

func TestScan(t *testing.T) {
	t.Run("Add one item successfully", func(t *testing.T) {
		store := testScanStoreChecker{}
		checkout := NewCheckout(store)

		err := checkout.Scan("A")
		if err != nil {
			t.Errorf("Should not error")
		}

		got := checkout.basket
		want := map[string]int{
			"A": 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Add one item successfully - ignoring case of input", func(t *testing.T) {
		store := testScanStoreChecker{}
		checkout := NewCheckout(store)

		err := checkout.Scan("a")
		if err != nil {
			t.Errorf("Should not error")
		}

		got := checkout.basket
		want := map[string]int{
			"A": 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Fails as input is too long", func(t *testing.T) {
		store := testScanStoreChecker{}
		checkout := NewCheckout(store)

		err := checkout.Scan("Scanning error")
		if err == nil {
			t.Errorf("Should error on account of being too long")
			return
		}

		got := err.Error()
		want := errors.New("input is too long").Error()

		if got != want {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Adding many of one item", func(t *testing.T) {
		store := testScanStoreChecker{}
		checkout := NewCheckout(store)

		items := []string{
			"A",
			"A",
			"A",
		}

		for i := range items {
			err := checkout.Scan(items[i])
			if err != nil {
				t.Errorf("Should not error")
			}
		}

		got := checkout.basket
		want := map[string]int{
			"A": len(items),
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Adding many different items out of order", func(t *testing.T) {
		store := testScanStoreChecker{}
		checkout := NewCheckout(store)

		items := []string{
			"A",
			"B",
			"A",
			"C",
			"B",
		}

		for i := range items {
			err := checkout.Scan(items[i])
			if err != nil {
				t.Errorf("Should not error")
			}
		}

		got := checkout.basket
		want := map[string]int{
			"A": 2,
			"B": 2,
			"C": 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Scanning items that are not in the store will return error", func(t *testing.T) {
		store := testScanStoreChecker{}
		checkout := NewCheckout(store)

		//item X will not be one of the store's available items. Thus you cannot scan it.
		err := checkout.Scan("X")
		if err == nil {
			t.Errorf("Should error on account of item not being in stock")
			return
		}

		got := err.Error()
		want := errors.New("scanned item is not in the store").Error()

		if got != want {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}

	})
}

func TestGetTotalPrice(t *testing.T) {
	t.Run("Add one item and succesfully and get total price", func(t *testing.T) {
		store := basicTestGetForTotalPrice{}
		checkout := NewCheckout(store)

		err := checkout.Scan("A")
		if err != nil {
			t.Errorf("Should not error")
			return
		}

		got, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Should not error")
			return
		}
		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Add different items and get correct price", func(t *testing.T) {
		store := basicTestGetForTotalPrice{}
		checkout := NewCheckout(store)

		items := []string{
			"A",
			"B",
			"C",
		}

		for i := range items {
			err := checkout.Scan(items[i])
			if err != nil {
				t.Errorf("Should not error")
			}
		}

		got, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Should not error")
			return
		}
		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Add same item multiple times with no offer gets the correct price", func(t *testing.T) {
		store := basicTestGetForTotalPrice{}
		checkout := NewCheckout(store)

		items := []string{
			"A",
			"A",
			"A",
		}

		for i := range items {
			err := checkout.Scan(items[i])
			if err != nil {
				t.Errorf("Should not error")
			}
		}

		got, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Should not error")
			return
		}
		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Adding multiple items out of order gets the correct price", func(t *testing.T) {
		store := basicTestGetForTotalPrice{}
		checkout := NewCheckout(store)

		items := []string{
			"A",
			"B",
			"C",
			"A",
			"A",
		}

		for i := range items {
			err := checkout.Scan(items[i])
			if err != nil {
				t.Errorf("Should not error")
			}
		}

		got, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Should not error")
			return
		}
		want := 8

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Error in get results in returned error", func(t *testing.T) {
		store := basicTestGetForTotalPrice{}
		checkout := NewCheckout(store)

		err := checkout.Scan("A")
		if err != nil {
			t.Errorf("Should not error")
			return
		}

		errorStore := errorTestGetTotalForPrice{}
		checkout.store = errorStore

		_, err = checkout.GetTotalPrice()
		if err == nil {
			t.Errorf("Should error")
			return
		}
		got := err.Error()
		want := errors.New("error in get").Error()

		if got != want {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Adding enough of one item to get special offer results in correct price", func(t *testing.T) {
		store := specialOfferTestGetTotalForPrice{}
		checkout := NewCheckout(store)

		for i := range 100 {
			err := checkout.Scan("A")
			if err != nil {
				t.Errorf("Should not error: errored on iteration %v", i)
			}
		}

		got, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Should not error")
			return
		}
		want := 10

		if got != want {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

	t.Run("Adding one more than threshold results in special price + 1 unit price", func(t *testing.T) {
		store := specialOfferTestGetTotalForPrice{}
		checkout := NewCheckout(store)

		for i := range 101 {
			err := checkout.Scan("A")
			if err != nil {
				t.Errorf("Should not error: errored on iteration %v", i)
			}
		}

		got, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Should not error")
			return
		}
		want := 11

		if got != want {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})
}
