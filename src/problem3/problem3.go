package problem3

// CountPrimes has a time complexity of O(n * log(log(n))) as it uses Sieve of Eratosthenes
func CountPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	// Create and initialize array of size n
	isPrimeNumber := make([]bool, n)
	for i := range isPrimeNumber {
		isPrimeNumber[i] = true
	}

	isPrimeNumber[0] = false
	isPrimeNumber[1] = false

	for i := 2; i*i < n; i++ {
		if isPrimeNumber[i] {
			for j := 2; i*j < n; j++ {
				isPrimeNumber[i*j] = false
			}
		}
	}

	// Count the prime numbers
	count := 0
	for i := range isPrimeNumber {
		if isPrimeNumber[i] {
			count++
		}
	}
	return count
}
