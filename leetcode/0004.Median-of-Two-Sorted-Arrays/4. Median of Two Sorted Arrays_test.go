package leetcode

import (
	"fmt"
	"testing"
)

type question4 struct {
	para4
	ans4
}

// para 是参数
// one 代表第一个参数
type para4 struct {
	nums1 []int
	nums2 []int
}

// ans 是答案
// one 代表第一个答案
type ans4 struct {
	one float64
}

func Test_Problem4(t *testing.T) {

	qs := []question4{

		// {
		// 	para4{[]int{1, 3}, []int{2}},
		// 	ans4{2.0},
		// },

		{
			para4{[]int{1, 3}, []int{}},
			ans4{2.0},
		},

		// {
		// 	para4{[]int{}, []int{}},
		// 	ans4{0.0},
		// },


		// {
		// 	para4{[]int{}, []int{2,3,4,5}},
		// 	ans4{3.5},
		// },

		// {
		// 	para4{[]int{1, 2}, []int{3, 4}},
		// 	ans4{2.5},
		// },
		// {
		// 	para4{[]int{1,3,5},[]int{2,3,4}},
		// 	ans4{3.0},
		// },

		// {
		// 	para4{[]int{1,3,5,7,8,8,10,15},[]int{2,3,4,8,9,11,13,14,19}},
		// 	ans4{8.0},
		// },
	}

	fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

	for _, q := range qs {
		q, p := q.ans4, q.para4
		//fmt.Printf("【input】:%v       【output】:%v\n", p, findMedianSortedArrays(p.nums1, p.nums2) == q.one)
		//fmt.Printf(" input : %v		output : %v \n", p, mergeTwoSortedArrsAndGetMedian(p.nums1,p.nums2)== q.one)

		fmt.Printf(" input : %v		output : %v \n", p, biSearchGetMedian(p.nums1,p.nums2)== q.one)
	}
	fmt.Printf("\n\n\n")
}

func TestGetMedian(t *testing.T) {
	qs := []question4{
		{
			para4{[]int{1,3,5,7,8,8,10,15},[]int{2,3,4,8,9,11,13,14,19}},
			ans4{2.0},
		},
	}
	for _, q := range qs {
		q,p := q.ans4,q.para4
		fmt.Printf(" input : %v		output : %v \n", p, getMedian(p.nums1,p.nums2))
		fmt.Printf(" input : %v		output : %v \n", p, mergeTwoSortedArrsAndGetMedian(p.nums1,p.nums2)== q.one)
	}
}