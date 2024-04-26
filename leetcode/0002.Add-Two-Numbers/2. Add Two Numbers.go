package leetcode

import (
	"fmt"
	"sort"

	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

func multiply(num1 string, num2 string) string {
	// fmt.Println(int(num1[0])-48)
	// return ""
	conv := 48
	//fmt.Println(400000000000*800000000000)
	lenNum1, lenNum2 := len(num1), len(num2)
	// if lenNum1 < lenNum2 {
	//     num1, num2 = num2, num1
	// }
	res := &ListNode{Val: 0}
	//var base int = 10
	//intArr1, intArr2 := make([]int, lenNum1), make([]int, lenNum2)
	for i := 0; i < lenNum1; i++ {
		//intArr1[i] = (int(num1[i]) - conv)
		tmp1 := (int(num1[i]) - conv)

		// tmp1, _ := strconv.Atoi(string(num1[i]))
		// fmt.Println("tmp1 ", tmp1, string(num1[i]))

		// for j := i + 1; j < lenNum1; j++ {
		// 	tmp1 *= 10
		// }
		j := (lenNum1 - i - 1)
		for m := 0; m < lenNum2; m++ {
			tmp2 := (int(num2[m]) - conv)
			// tmp2, _ := strconv.Atoi(string(num2[m]))
			// fmt.Println("tmp2 ", tmp2, string(num2[m]))
			// for n := m + 1; n < lenNum2; n++ {
			// 	tmp2 *= 10
			// }
			n := (lenNum2 - m - 1)
			//fmt.Println("tmp1, tmp2", tmp1, tmp2, tmp1*tmp2)
			tmp := (tmp1 * tmp2)
			baseNode := &ListNode{Val: 0}
			tmpNode := baseNode
			for k := 0; k < (j + n); k++ {
				//fmt.Println("steps ", tmp, k)
				tmpNode.Next = &ListNode{Val: 0}
				tmpNode = tmpNode.Next
			}
			tmpNode.Next = &ListNode{Val: tmp % 10}
			tmpNode.Next = &ListNode{Val: tmp / 10}
			fmt.Println("the tmp is ", tmp)
			res = addTwoNumbers(res, baseNode)
			fmt.Println("result int is ", res)
		}
	}
	fmt.Printf("the number res is %v /n", res)
	for res.Next != nil {
		fmt.Print(res.Val)
		res = res.Next
	}
	// if res == 0 {
	// 	return "0"
	// }
	// resStr := []string{}
	// for res > 0 {
	// 	resStr = append(resStr, strconv.Itoa(int(res%10)))
	// 	res = res / 10
	// }
	// result := ""
	// for i := len(resStr) - 1; i >= 0; i-- {
	// 	result += resStr[i]
	// }

	return "result"
}

func multiply1(num1 string, num2 string) string {

	l1, l2 := len(num1), len(num2)
	res := [][]uint8{}
	var base uint8 = 10
	for i := l1 - 1; i >= 0; i-- {

		tmp1 := (num1[i] - byte('0'))
		fmt.Println(tmp1)
		resTmp := []uint8{}
		var carry uint8 = 0
		for i := l2 - 1; i >= 0; i-- {
			tmp2 := num2[i] - byte('0')
			fmt.Println(tmp2)
			tmp := (tmp1*tmp2 + carry)
			carry = tmp / base
			resTmp = append(resTmp, tmp%base)
		}
		if carry > 0 {
			resTmp = append(resTmp, carry)
		}
		res = append(res, resTmp)
	}
	fmt.Println("res is ", res)
	//result := []uint8{}
	//a := math.Multiply(12, 32)

	return "result"
}

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	res := make([]byte, l1+l2)
	//var base uint8 = 10
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			tmp := (num1[i] - '0') * (num2[j] - '0')
			res[i+j+1] += tmp
			fmt.Println("res is ", res)
			if res[i+j+1] >= 10 {
				res[i+j] += (res[i+j+1] / 10)
				res[i+j+1] %= 10
				fmt.Println("res is >>> ", res)
			}
		}
	}
	fmt.Println("res is ", res)
	if res[0] == 0 {
		res = res[1:]
	}
	for i := range res {
		res[i] += '0'
	}
	result := string(res)
	fmt.Println(result)
	return result
}

func multiply4(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	res := make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			val := (num1[i] - '0') * (num2[j] - '0')
			res[i+j+1] += val
			if res[i+j+1] >= 10 {
				res[i+j] += (res[i+j+1] / 10)
				res[i+j+1] %= 10
			}
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for i := range res {
		res[i] += '0'
	}
	result := string(res)
	fmt.Println(result)
	return result
}

func Multiply5(s1, s2 string) string {
	if s1 == "0" || s2 == "0" {
		return "0"
	}
	//base :=
	l1, l2 := len(s1), len(s2)
	res := make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			tmp := (s1[i] - '0') * (s2[j] - '0')
			res[i+j+1] += tmp
			if res[i+j+1] >= 10 {
				res[i+j] += (res[i+j+1] / 10)
				res[i+j+1] %= 10

			}
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for i, item := range res {
		res[i] = item + '0'
	}
	fmt.Println(string(res))
	return string(res)
}

type linkedNode struct {
	val  byte
	next *linkedNode
}

func convStrtoLinkedNode(str string) *linkedNode {
	head := &linkedNode{}
	tmp := head
	for i := 0; i < len(str); i++ {
		// by - '0] ; make the val as the same apparance as str[i] in ascill format
		tmp.val = (str[i] - '0')
		//fmt.Println(tmp.val)
		// put this check here in order to avoid extra nil node
		if i == (len(str) - 1) {
			break
		}
		tmp.next = &linkedNode{}
		tmp = tmp.next
	}
	return head
}

func convertLinkedNodeToStr(src *linkedNode) string {
	resBytes := []byte{}
	for src != nil {
		resBytes = append(resBytes, src.val+'0')
		src = src.next

	}
	return string(resBytes)

}

func equalLinkedNodes(target []*linkedNode) bool {
	if len(target) == 0 {
		return true
	}
	standard := convertLinkedNodeToStr(target[0])
	for i := 1; i < len(target); i++ {
		if standard != convertLinkedNodeToStr(target[i]) {
			return false
		}
	}
	return true
}

// 1, convert all items to linkedNode; 2, category the items by the linkedNode head; the depth is 1
// 3, keep grouping the subItems by the depth and adding the depth by 1
func maxStr(src []string) string {
	// no need to sort the src
	// sort.Strings(src)
	// fmt.Println(src)
	result := []byte{}
	//res := []*linkedNode{}
	// convert the item to linkedNode and category them by the head
	srcMap := make(map[byte][]*linkedNode)
	for _, str := range src {
		//fmt.Println("=========", str)
		tmp := convStrtoLinkedNode(str)
		//res = append(res, tmp)
		key := tmp.val
		//keys = append(keys, key)
		//srcMap[key] = []*linkedNode{}
		srcMap[key] = append(srcMap[key], tmp)
	}
	// pass the pointer of the result slice and and the depth now is 1
	appendRecursively(srcMap, &result, 1)
	for k, v := range result {
		result[k] = (v + '0')
	}
	fmt.Println("the result is ", string(result))
	return string(result)
}

func appendRecursively(src map[byte][]*linkedNode, result *[]byte, n uint) {
	// sort the keys of the map
	keys := []byte{}
	for k, _ := range src {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	// iterate the categories in the order of keys
	for _, k := range keys {
		subItems := src[k]
		if len(subItems) == 0 {
			continue
		}
		subHead := subItems[0]
		if len(subItems) > 1 {

			// recusive to next level
			if equalLinkedNodes(subItems) {
				// if the items are the same in the category, then append them
				for _, subHead := range subItems {
					// append the item of linkedNode
					for subHead != nil {
						*result = append(*result, subHead.val)
						subHead = subHead.next
					}
				}
			} else {
				// if the items are not the same then recursive to next level
				subItemMap := make(map[byte][]*linkedNode)
				for i := 0; i < len(subItems); i++ {
					// get the item
					head := subItems[i]
					// get the k according to the level
					headTmp := head
					k := headTmp.val
					level := n
					for level > 0 {
						// heatTmp have no case to be nil. it would not appear in this category
						headTmp = headTmp.next
						if headTmp != nil {
							k = headTmp.val
						}
						level--
					}
					subItemMap[k] = append(subItemMap[k], head)
				}
				// recursive and add depth by 1
				appendRecursively(subItemMap, result, n+1)
			}
		} else {
			// there is only one item in this category.
			for subHead != nil {
				*result = append(*result, subHead.val)
				subHead = subHead.next
			}
		}
	}
}

func isZiShu(n int) bool {
	if n == 1 {
		return false
	}
	res := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			res = false
		}
	}
	return res
}
// refactor the number to a slice of zhishu
func getZhiShu(n int) []int {
	resArr := []int{}
	for i := 2; i <= n; i++ {
		//fmt.Println(i,n)
		if n%i == 0 {
			n /= i
			// divide the i part first
			if isZiShu(i) {
				resArr = append(resArr, i)
			} else {
				resArr = append(resArr, getZhiShu(i)...)
			}
			// then divide the n part
			// if n can not divide
			if isZiShu(n) {
				resArr = append(resArr, n)
				// return to avoid zhishu being added twice
				// return resArr
				break
			} else {
				// avoid not dividing n in case of (i+1) > n
				if i >= n {
					fmt.Println(i,n)
					resArr = append(resArr, getZhiShu(n)...)
				}
			}
		}
	}
	//fmt.Println(resArr)
	return resArr
}
