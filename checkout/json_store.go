package checkout

import "fmt"

type JsonStore struct {
	items map[string]StoreItem
}

func (store JsonStore) Get(s string) (StoreItem, error) {
	storeItem, ok := store.items[s]
	if !ok {
		return StoreItem{}, fmt.Errorf("store item %v not found", s)
	}

	return storeItem, nil
}

func (store JsonStore) Set(s string, storeItem StoreItem) {
	store.items[s] = storeItem

}
