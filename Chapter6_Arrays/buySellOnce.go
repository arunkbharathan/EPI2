package main

func maxProfit(arr [10]int) int {
	minsofar := int(^uint(0) >> 1)
	profit := 0
	for _, price := range arr {
		observedProfit := price - minsofar
		profit = max(observedProfit, profit)
		minsofar = min(minsofar, price)
	}
	return profit
}
