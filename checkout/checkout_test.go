package checkout

import (
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
}
