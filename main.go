package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tbd20/check-me-out/checkout"
)

func main() {
	fmt.Println("Hello and welcome to the checkout service!")

	filePath := os.Args[1]
	fmt.Printf("Looking for file in path %s \n", filePath)

	store, err := checkout.NewJsonStore(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	basket := checkout.NewCheckout(store)

	scannedItems := os.Args[2]
	for _, character := range scannedItems {
		err := basket.Scan(string(character))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	yourTotal, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
	}

	myTotal, err := basket.GetTotalPrice()
	if err != nil {
		fmt.Println(err)
	}

	if myTotal != yourTotal {
		fmt.Printf("Oops! Looks like our totals don't match! \nMy total was %d, and yours was %d \n \n", myTotal, yourTotal)
		return
	}

	fmt.Printf("Yay! Our totals match.\nWe both got the total as: %d \n \n ", myTotal)
}
