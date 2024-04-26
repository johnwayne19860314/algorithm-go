package leetcode

import "fmt"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	// loop untill num == 0
	for num != 0 {
		// increment i till the num > values[i]
		for values[i] > num {
			i++
		}
		// decrease the num by values[i] and increase the res by symbols[i]
		num -= values[i]
		res += symbols[i]
	}

	fmt.Println("the result is ", res)
	return res
}

func intToRomain0402(i int) string {
	// 1289
	// 	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	if i > 3999 || i < 1 {
		panic("the input is above the upper limit of 3999 or lower than the bottom limit of 1")
	}
	ctxMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	// form an array to get the string
	res := []string{}
	base := 10
	// record the times of dividing by the base. it works as a scale
	times := 1
	for i > 0 {
		// start from left side of the integer
		mod := i % base
		// make it map to the scales in the ctxMap
		mod *= times
		// if not in the ctxMap, then it is either 6,7 or 3,2
		phase1 := 5 * times
		phase2 := 1 * times
		if val, ok := ctxMap[mod]; ok {
			// if mapping, then just append
			res = append(res, val)

		} else {
			// if not, then either less or more than the phase1
			mod -= phase1
			abovePhase1 := false
			if mod > 0 {
				abovePhase1 = true
			} else {
				mod += phase1
			}
			// append the phase2 after making it less than phase1
			for mod > 0 {
				// make sure it matches with reminder
				res = append(res, ctxMap[phase2])
				mod -= phase2
			}
			// after phase2, then append phase1
			if abovePhase1 {
				res = append(res, ctxMap[phase1])
			}
		}
		// move to a larger scale
		i = i / base
		times *= base

	}
	// form a string from the array from right to left
	resStr := ""
	for i := len(res) - 1; i >= 0; i-- {
		resStr += res[i]
	}
	fmt.Println("the result is ", resStr)
	return resStr
}
