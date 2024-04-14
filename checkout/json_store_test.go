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

	t.Run("Test new Json store - With no special offers", func(t *testing.T) {
		filePath := "../testFiles/noSpecialOffers.json"
		store, err := NewJsonStore(filePath)
		if err != nil {
			t.Errorf("Error - reading file should not fail")
		}

		got := store.items
		want := map[string]StoreItem{
			"C": {sku: "C", value: 20},
			"D": {sku: "D", value: 15},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})
	t.Run("Test new Json store - Error in reading file", func(t *testing.T) {
		filePath := "../testFiles/notValidJson.json"
		_, err := NewJsonStore(filePath)
		if err == nil {
			t.Errorf("Error - reading file should not fail")
		}
	})

	t.Run("Test newJsonStore - can read special offers", func(t *testing.T) {
		filePath := "../testFiles/basic.json"
		store, err := NewJsonStore(filePath)
		if err != nil {
			t.Errorf("Error - reading file should not fail")
		}

		got := store.items
		want := map[string]StoreItem{
			"A": {sku: "A", value: 50, specialOffer: &SpecialOffer{threshold: 3, thresholdAmountValue: 130}},
			"B": {sku: "B", value: 30, specialOffer: &SpecialOffer{threshold: 2, thresholdAmountValue: 45}},
			"C": {sku: "C", value: 20},
			"D": {sku: "D", value: 15},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in testing: got %v, want %v", got, want)
		}
	})
}
