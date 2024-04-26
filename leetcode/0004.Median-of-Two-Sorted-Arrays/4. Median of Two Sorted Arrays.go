package leetcode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMeidanUtil(nums []int) int {
	l := len(nums)
	if l%2 == 0 {
		return (nums[l/2] + nums[l/2-1]) / 2
	} else {
		return nums[l/2]
	}
}

func moveCompare(a, b, q int, num1, num2, num []int) (int, int, int) {
	num[q] = num1[a]
	fmt.Println(num[q])
	q++
	if a < (len(num1) - 1) {
		a++
	} else {
		// if a move to the end of num1, then copy all values from num2 to num
		for b < len(num2) {
			num[q] = num2[b]
			fmt.Println(num[q])
			b++
			q++
		}
	}
	return a, b, q
}

// merge two sorted arrays and then get the median
func getMedian(num1, num2 []int) int {
	m, n := len(num1), len(num2)
	if m == 0 && n == 0 {
		return 0
	}
	if n == 0 {
		return getMeidanUtil(num1)
	}
	if m == 0 {
		return getMeidanUtil(num2)
	}
	num := make([]int, m+n)
	i, j, q := 0, 0, 0
	for q < (m + n) {
		// get the lesser num to fill in
		if num1[i] <= num2[j] {
			i, j, q = moveCompare(i, j, q, num1, num2, num)
		} else {
			j, i, q = moveCompare(j, i, q, num2, num1, num)
		}
	}
	return getMeidanUtil(num)
}

func mergeTwoSortedArrsAndGetMedian(a, b []int)  float64 {
	i, j, k, la, lb := 0, 0, 0, len(a), len(b)
	merged := make([]int, la+lb)
	if la == 0 && lb == 0 {
		return 0.0
	}else if la == 0 {
		merged = b
		k = lb
	}else if lb == 0 {
		merged = a
		k = la
	}else {
		for  k < len(merged) {
			if i < la && j < lb && a[i] <= b[j] {
				merged[k] = a[i]
				i++
			} else {
				merged[k] = b[j]
				j++
			}
			k++
		}
	}

	k -=1
	var res float64
	if k%2 == 0 {
		res =float64(merged[k/2]) 
	} else {
		res = (float64(merged[k/2]) + float64(merged[k/2+1])) / 2
	}
	fmt.Println("the median is ", res)
	return res
}


func biSearchGetMedian(a,b []int) float64{
	if len(a) == 0 && len(b) == 0 {
		return 0.0
	}

	if len(a) == 0 {
		return getMedianFromSortedArr(b)
	}else if len(b) == 0 {
		return getMedianFromSortedArr(a)
	}

	if len(a) > len(b)  {
		biSearchGetMedian(b,a)
	}
	
	// search and find in a which meets a[A-1] < b[B] && a[A] > b[B-1]
	medianALow, la, lb, medianA, medianB :=  0, len(a), len(b) ,0,0
	medianPosition := (la+lb)/2
	medianAHigh := la
	
	flag := true
	for flag {
		medianA = (medianALow + medianAHigh)/2
		medianB = (medianPosition - medianA)
		if medianA > 0 && a[medianA-1] > b[medianB] {
			medianAHigh = medianA
		} else if medianB > 0 && a[medianA] > b[medianB-1] {
			medianALow = medianA
		} else {
			flag = false
		}
	}

	// if odd ; max(a[medianA-1], a[medianB-1])
	// if even ; (max(a[medianA-1], a[medianB-1]) + min(a[medianA] , b[medianB]))/2
	if (la+lb)%2 == 1 {
		return float64(max(a[medianA-1], a[medianB-1]))
	} else {
		return (float64(max(a[medianA-1], a[medianB-1])) + float64(min(a[medianA] , b[medianB])))/2
	}

}

func getMedianFromSortedArr(a []int) float64 {
	k := len(a)/2
	if len(a) %2 == 1 {
		return float64(a[k])
	}else {
		return (float64(a[k-1]) + float64(a[k]))/2
	}
}