package leetcode

import (
	"sort"
)

// 解法一 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

// 解法二
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

func threeSum0406(input []int) [][]int {
	sort.Ints(input)
	//record the times of the dup nums
	inputMap := make(map[int]int)
	// separate the input into pos and neg
	posSli, negSli, res := []int{}, []int{}, [][]int{}
	for _, i := range input {
		inputMap[i]++
		if i >= 0 {
			posSli = append(posSli, i)
		} else {
			negSli = append(negSli, i)
		}
	}
	//fmt.Println(jumpMap)
	// deal with unique case
	if len(negSli) == 0 {
		if inputMap[0] >= 3 {
			return [][]int{{0, 0, 0}}
		} else {
			panic("will fail")
		}
	}
	i := 0
	for i < len(negSli) {
		val := negSli[i]
		// two neg and one pos
		j := i + 1
		for j < (len(negSli)) {
			valj, jumpVal := negSli[j], 1
			//fmt.Println(val, valj)
			addSlice(val, valj, inputMap, &res)
			// algo on how to jump
			// if dup more than two
			if inputMap[valj] > 2 {
				if valj == val {
					// case : val and valj are the dup
					jumpVal = (inputMap[valj] - 1)
				} else {
					// case : val and valj are not in the same dup
					jumpVal = (inputMap[valj])
				}
			}
			j += jumpVal
		}
		// one neg and two pos
		l := 0
		for l < len(posSli) {
			valL := posSli[l]
			addSlice(val, valL, inputMap, &res)
			l += (inputMap[valL])

		}
		// jump i -- the first ele
		i += (inputMap[val])
	}
	return res
}

func addSlice(i, j int, inputMap map[int]int, res *[][]int) {
	tmp := 0 - (i + j)
	// if condition meet
	if _, ok := inputMap[tmp]; ok {
		// subtract the the used ele in inputMap.
		inputMap[i]--
		inputMap[j]--
		// if tmp exit in the input and remove the dup by asserting the increasing order
		// i <=j <= tmp
		if inputMap[tmp] > 0 && j <= tmp {
			// if i == -10 && j == 8 && tmp == 2 {
			// 	fmt.Println("denig")
			// }
			tmpSlice := []int{i, j, tmp}
			*res = append(*res, tmpSlice)
		}
		// restore the counter inputMap
		inputMap[i]++
		inputMap[j]++
	}
}
