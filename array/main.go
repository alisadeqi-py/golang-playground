package main 

import "fmt"

func main() {
	var productNames [4]string = [4]string{
		"book",
	}

	prices := [4]float64{44.3, 55.1, 2.4, 99.7}
	fmt.Println(prices[2])

	productNames[2] = "bike"

	fmt.Println(productNames)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.8
	highlighttedPrices := featuredPrices[1:]
	fmt.Println(highlighttedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlighttedPrices), cap(highlighttedPrices))

	highlighttedPrices = highlighttedPrices[:2]
	fmt.Println(len(highlighttedPrices), cap(highlighttedPrices))
	fmt.Println(highlighttedPrices)
}
