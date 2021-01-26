package main

import (
    "fmt"
    "math"
    )

func main() {
	var price int = 100
	fmt.Println("Price is ", price, " dollars")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is ", tax, " dollars")

	var total float64 = float64(price) + tax
	fmt.Println("total cost is ", total, " dollars")

	var availableFunds = 120
	fmt.Println(availableFunds, " dollars available")
	fmt.Println("Within the budget? ", total <= float64(availableFunds))

	lower := math.Floor(4.566)
	fmt.Println(lower)
}
