package main

import (
	"fmt"
	"math"
)

func main() {
	scaffolding := []float64{2, 9, 4, 5, 1, 6, 10}

	var minCost float64 = solve(scaffolding)
	fmt.Printf("重みの総和は%gです。", minCost)
}

func solve(argScaffolding []float64) float64 {
	jumpOne := []float64{}
	for i := 0; i < len(argScaffolding); i++ {
		if i < 1 {
			jumpOne = append(jumpOne, 0)
		} else {
			var tmp float64 = argScaffolding[i] - argScaffolding[i-1]
			jumpOne = append(jumpOne, math.Abs(tmp))
		}
	}

	jumpTwo := []float64{}
	for i := 0; i < len(argScaffolding); i++ {
		if i < 2 {
			jumpTwo = append(jumpTwo, 0)
		} else {
			var tmp float64 = argScaffolding[i] - argScaffolding[i-2]
			jumpTwo = append(jumpTwo, math.Abs(tmp))
		}
	}

	dp := []float64{}
	for i := 0; i < len(argScaffolding); i++ {
		if i < 1 {
			dp = append(dp, 0)
		} else if i < 2 {
			dp = append(dp, jumpOne[1])
		} else {
			var edgeOne float64 = dp[i-1] + jumpOne[i]
			var edgeTwo float64 = dp[i-2] + jumpTwo[i]
			var minCost float64 = math.Min(edgeOne, edgeTwo)
			dp = append(dp, minCost)
		}
	}
	maxIndexNum := len(argScaffolding) - 1
	return dp[maxIndexNum]
}
