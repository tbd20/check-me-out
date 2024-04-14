package checkout

import (
	"reflect"
	"testing"
)

func TestJsonStore(t *testing.T) {
	t.Run("Test Set Item", func(t *testing.T) {
		store := JsonStore{
			items: make(map[string]StoreItem),
		}
		store.Set("A", StoreItem{
			sku:   "A",
			value: 1,
		})

		got := store.items
		want := map[string]StoreItem{
			"A": {sku: "A", value: 1},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})
	t.Run("Test Get Item", func(t *testing.T) {
		store := JsonStore{
			items: make(map[string]StoreItem),
		}
		store.Set("A", StoreItem{
			sku:   "A",
			value: 1,
		})

		got, err := store.Get("A")
		if err != nil {
			t.Errorf("Error get should not fail")
		}
		want := StoreItem{sku: "A", value: 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})
}
