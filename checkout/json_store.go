package checkout

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type JsonStore struct {
	items map[string]StoreItem
}

type JsonStoreInputItem struct {
	Sku          string `json:"sku"`
	Value        int    `json:"value"`
	SpecialOffer struct {
		Threshold            int `json:"threshold"`
		ThresholdAmountValue int `json:"thresholdAmountValue"`
	} `json:"specialOffer,omitempty"`
}

type JsonStoreInput []JsonStoreInputItem

func NewJsonStore(filePath string) (*JsonStore, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var storeInput JsonStoreInput

	err = json.Unmarshal(byteValue, &storeInput)
	if err != nil {
		return nil, err
	}

	jsonStore := JsonStore{items: make(map[string]StoreItem)}

	for _, item := range storeInput {
		storeItem := StoreItem{
			sku:   item.Sku,
			value: item.Value,
		}

		if item.SpecialOffer.Threshold != 0 {
			specialOffer := SpecialOffer{
				threshold:            item.SpecialOffer.Threshold,
				thresholdAmountValue: item.SpecialOffer.ThresholdAmountValue,
			}
			storeItem.specialOffer = &specialOffer
		}

		jsonStore.Set(item.Sku, storeItem)

	}

	return &jsonStore, nil

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
