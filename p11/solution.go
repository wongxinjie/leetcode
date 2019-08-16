/**
* @File: solution.go
* @Author: wongxinjie
* @Date: 2019/8/16
*/
package p11

func maxArea(height []int) int {
	//area, currentArea := 0, 0
	//size := len(height)
	//
	//for i := 0; i < size - 1; i ++ {
	//	yAxis := height[i]
	//	for j := i + 1; j < size; j++ {
	//		if height[j] < yAxis {
	//			currentArea = height[j] * (j - i)
	//		} else {
	//			currentArea = yAxis * (j - i)
	//		}
	//		if currentArea > area {
	//			area = currentArea
	//		}
	//	}
	//}
	//return area

	maxArea := 0
	l, r := 0, len(height) - 1
	for l < r {
		y := 0
		if height[l] < height[r] {
			y = height[l]
			l += 1
		} else {
			y = height[r]
			r -= 1
		}
		currentArea := y * (r - l + 1)
		if currentArea > maxArea {
			maxArea = currentArea
		}
	}
	return maxArea
}
