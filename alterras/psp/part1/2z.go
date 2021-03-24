package main

import "fmt"


var COINS []int = []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 1000}

func main() {
	fmt.Println(getMinimumCoins[543])
	fmt.Println(getMinimumCoins[7752])
}

func getMinimumCoins(totalMoney int) []int {
	var usedCoins []int
	checkTotalMoney := 0
	for totalMoney > 0 {
		for i := len(COINS) - 1; i >= 0; i-- {
			checkTotalMoney = totalMoney - COINS[i]
			fmt.Println("check", checkTotalMoney)
			if checkTotalMoney >= 0 {
				usedCoins = append(usedCoins, COINS[i])
				totalMoney = totalMoney - COINS[i]
				fmt.Println(totalMoney, usedCoins)
				break
			} else {
				continue
			}
		}
	}
	return usedCoins
}