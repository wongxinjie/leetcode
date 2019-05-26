/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/5/9
*/
package main

func countPrimes(n int) int {
	isPrimeArray := make([]bool, 0, n)
	for i := 0; i < n; i++ {
		isPrimeArray = append(isPrimeArray, true)
	}

	for i := 2; i * i < n; i++ {
		if !isPrimeArray[i] {
			continue
		}

		for j := i * i; j < n; j += i {
			isPrimeArray[j] = false
		}
	}

	count := 0
	for  i := 2; i < n; i++ {
		if isPrimeArray[i] {
			count += 1
		}
	}
	return count
}
