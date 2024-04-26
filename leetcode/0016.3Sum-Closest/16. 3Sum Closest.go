package leetcode

import (
	"fmt"
	"math"
	"sort"
)

// 解法一 O(n^2)
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

// 解法二 暴力解法 O(n^3)
func threeSumClosest1(nums []int, target int) int {
	res, difference := 0, math.MaxInt16
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < difference {
					difference = abs(nums[i] + nums[j] + nums[k] - target)
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func threeSumClosest0409(nums []int, target int) int {
	
	sort.Ints(nums)
	n,diff, res := len(nums), math.MaxInt, 0
	a,b,c := 0,0,0
	if n < 2 {
		panic("array length must be at least three items")
	}
	for i := 0; i < n-2; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		j,k := i+1, n-1
		for  j < k {

			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < diff {
				a,b,c = nums[i],nums[j],nums[k]
				res, diff = sum, abs(sum-target)
			}
			if sum > target {
				k--
			}else if sum == target {
				fmt.Println(a,b,c)
				return res
			}else {
				j++
			}
		}
	}
	fmt.Println(a,b,c)
	return res
}

func absPlusVal(a,b int) int {
	res := a+b
	if res < 0{
		res *= (-1)
	}
	return res
}