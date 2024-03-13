package main

import "fmt"

func pivotInteger(n int) int {
	sumLeft := 1
	sumRight := n
	leftValue := 1
	rightValue := n

	if n == 1 {
		return 1
	}

	for i := 1; i <= n; i++ {
		if sumLeft < sumRight {
			sumLeft += leftValue + 1
			leftValue++
		} else {
			sumRight += rightValue - 1
			rightValue--
		}
		if sumLeft == sumRight && leftValue+1 == rightValue-1 {
			return leftValue + 1
		}
	}
	return -1
}

func betterPivot(n int) int {
	if n == 1 {
		return 1
	}
	left := 1
	right := n
	mid := 0
	totalSum := n * (n + 1) / 2

	for i := 0; i < n; i++ {
		mid = (left + right) / 2
		if mid*mid < totalSum {
			left = mid + 1
		} else if mid*mid >= totalSum {
			right = mid
		}
		if left*left == totalSum {
			return left
		}
	}

	return -1
}

func main() {
	fmt.Println(betterPivot(4))
}
