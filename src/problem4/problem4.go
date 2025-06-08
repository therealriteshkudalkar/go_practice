package problem4

import "fmt"

func MyPowRec(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var answer float64
	if n%2 == 0 {
		answer = MyPowRec(x, n/2) * MyPowRec(x, n/2)
	} else {
		if n < 0 {
			answer = 1 / x * MyPowRec(x, n/2) * MyPowRec(x, n/2)
		} else {
			answer = x * MyPowRec(x, n/2) * MyPowRec(x, n/2)
		}
	}
	return answer
}

func MyPowRecMemoized(memo map[string]float64, x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	key := fmt.Sprintf("%v,%v", x, n)
	result, ok := memo[key]
	if ok {
		return result
	}

	if n%2 == 0 {
		result = MyPowRecMemoized(memo, x, n/2) * MyPowRecMemoized(memo, x, n/2)
	} else {
		if n < 0 {
			result = 1 / x * MyPowRecMemoized(memo, x, n/2) * MyPowRecMemoized(memo, x, n/2)
		} else {
			result = x * MyPowRecMemoized(memo, x, n/2) * MyPowRecMemoized(memo, x, n/2)
		}
	}
	memo[key] = result
	return result
}

func MyPow(x float64, n int) float64 {
	result := 1.0
	temp := 0
	for temp != n {

	}
	return result
}
