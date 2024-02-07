package main

import "fmt"

func main() {

	priceArray := []int{7, 1, 5, 5, 4, 8, 2}

	minPurchasePrice := priceArray[0]
	maxProfit := 0

	for i := 1; i < len(priceArray); i++ {

		// Check If the next element have lower purchase price
		// Replace the minimum purchase amout value

		if priceArray[i] < minPurchasePrice {
			minPurchasePrice = priceArray[i]
		} else if (priceArray[i] - minPurchasePrice) > maxProfit {
			//check if the difference is greater than max profit
			maxProfit = priceArray[i] - minPurchasePrice
		}

	}

	fmt.Println("The maximum Profit is", maxProfit)

}
