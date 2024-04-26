package leetcode

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whitespaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}
		whitespaceAllowed, signAllowed = false, false
		digits = append(digits, int(c-48))
	}
	var num, place int64
	place, num = 1, 0
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}
	var rtnMax int64
	if sign > 0 {
		rtnMax = maxInt - 1
	} else {
		rtnMax = maxInt
	}
	digitsCount := len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place
		place *= 10
		if digitsCount-i > 10 || num > rtnMax {
			return int(int64(sign) * rtnMax)
		}
	}
	num *= int64(sign)
	return int(num)
}

type linkedNode struct{
	val uint
	next *linkedNode
}

func myAtoi0318(s string) int {
	s = strings.Trim(s," ")
	head := rune(s[0])
	nodes := linkedNode{}
	positive := true
	var depth uint
	

	if ! unicode.IsDigit(head) {
		if head == '-' {
			positive = false
		}else {
			return 0
		}
	}
	if positive {
		nodes.val = uint(head-'0')	
	}
	i := 1
	depth ++
	tmp := &nodes
	for i< len(s){
		if unicode.IsDigit(rune(s[i])) {
			tmp.next = &linkedNode{
				val:uint(s[i]-'0'),
			} 
			tmp = tmp.next
			depth ++
		}else{
			break
		}
		i++
	}

	var res uint
	for  depth > 0 {
		depth --
		res += (uint(math.Pow10(int(depth)))*nodes.val)
		if nodes.next != nil {
			nodes = *nodes.next
		}
		
		
	}
	
	if !positive {
		if int(res) > int(math.Pow(2.0,231.0)) {
			return -1 * int(math.Pow(2.0,231.0))
		}
		return -1 * int(res)
	}else {
		if int(res) > int(math.Pow(2.0,230.0)) {
			return int(math.Pow(2.0,230.0))
		}
	}
	return int(res)
}