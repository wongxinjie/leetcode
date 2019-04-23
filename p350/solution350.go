package main

<<<<<<< HEAD
import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	counter1 := make(map[int]int)
	counter2 := make(map[int]int)

	for _, n := range nums1 {
		if c, ok := counter1[n]; ok {
			counter1[n] = c + 1
		} else {
			counter1[n] = 1
		}
	}

	for _, n := range nums2 {
		if c, ok := counter2[n]; ok {
			counter2[n] = c + 1
		} else {
			counter2[n] = 1
		}
	}

	result := make([]int, 0)
	for k, v := range counter1 {
		if c, ok := counter2[k]; ok {
			if c > v {
				c = v
			}
			for i := 0; i < c; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}

func main() {
	num1 := []int{4, 9, 5}
	num2 := []int{9, 4, 9, 8, 4}
	fmt.Printf("%v\n", intersect(num1, num2))
}
