package leetcode

import "fmt"

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		
		width := end - start
		high := 0
		//fmt.Printf("====== the start is %v ; the end is %v ; the width is %v \n", start,end, width)
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		//fmt.Printf("the start is %v ; the end is %v ; the width is %v \n", start,end, width)

		temp := width * high
		if temp > max {
			max = temp
		}
		fmt.Printf("the start is %v ; the end is %v ; the width is %v; the high is %v; the maxVal is %v \n", start,end, width, high, max)
	}
	return max
}

func maxArea0402(input []int) int {
	maxVal := 0

	i, l := 0, len(input)
	// two loops make the algorithm complex as O(n*n)
	for i < l {
		tmpIndex := i + 1
		for tmpIndex < l {
			tmpVal := (tmpIndex - i) * max(input[i], input[tmpIndex])
			if tmpVal > maxVal {
				maxVal = tmpVal
			}
			tmpIndex++
		}
		i++

	}
	return maxVal
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
