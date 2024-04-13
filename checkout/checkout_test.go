package checkout

import (
	"errors"
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	t.Run("Add one item successfully", func(t *testing.T) {
		checkout := NewCheckout()

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
		checkout := NewCheckout()

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
		checkout := NewCheckout()

		err := checkout.Scan("Scanning error")
		if err == nil {
			t.Errorf("Should error on account of being too long")
		}

		got := err.Error()
		want := errors.New("input is too long").Error()

		if got != want {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})

}
