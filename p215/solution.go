/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/7
*/
package p215

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	k = len(nums)  + 1 - k
	start, end :=  0, len(nums) - 1

	for start < end {
		left, right := start, end
		pivot := nums[right]

		for left < right {
			for nums[left] < pivot {
				left += 1
			}
			for left < right && nums[right] >= pivot {
				right -= 1
			}
			nums[left], nums[right] = nums[right], nums[left]
		}

		nums[left], nums[end] = nums[end], nums[left]

		if left == k - 1 {
			return nums[left]
		} else if left < k - 1 {
			start = left + 1
		} else {
			end = left - 1
		}
	}
	return nums[start]
}


func main() {
	nums := []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLargest(nums, 4))
}
