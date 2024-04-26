package leetcode

import "strconv"

// 解法一
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 解法二 数字转字符串
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome0322(input int) bool {
	// 10987
	if input < 0 {
		return false
	}
	base := 10
	tmp := input
	intArr := make([]int,0)
	for tmp > 0 {
		intArr = append(intArr, tmp%base)
		tmp /= base
	}
	// now it shall be [7,8,9,0,1]
	len := len(intArr)
	for i:=0;i<len/2;i++{
		// how abut len/2+1 in the case of odd
		if intArr[i] != intArr[len-i-1]{
			return false
		}
	}
	return true
}