package leetcode

// 解法一 位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++

		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getLoggestSubStr(input string) string{

	if input == "" {
		return ""
	}
	longest := ""
	strMap := map[byte]uint8{}
	i,j:=0,0
	for j < len(input){
		// add the tmp into strMap one by one
		tmp := input[j]
		strMap[tmp] ++
		// handle the case that the tmp already in strMap; j stop 
		if strMap[tmp] >1 {
			// if j stop; then compute and compare the substring
			tmpLoggest := input[i:j]
			if len(tmpLoggest) > len(longest) {
				longest = tmpLoggest
			}
			// move i one step forward first and subtract it from the map
			strMap[input[i]]--
			i++
			// in case of "pwwkew". please note there is no way that i will be over j
			for strMap[input[i]]>1{
				strMap[input[i]]--
				i++
			}
		}
		j++
	}
	return longest
}

func getLonggestSubStr0313(input string) int {
	if input == "" {
		return 0
	}
	res := 1
	i,j := 0,1
	cache := map[byte]int{}
	// init the cache when i is at the position of 0
	// no need for val1 after this
	val1 := input[i]
	cache[val1] = i
	for i < len(input) && j < len(input) {
		val2 :=  input[j]	
		if _, ok := cache[val2]; !ok {
			// cache val2 index
			cache[val2] = j
			j++
		}else{
			// compare to get the longgest
			res = max(res, (j-i))
			// move index i one step forward from the old index
			i = (cache[val2]+1)
			// update the index 
			cache[val2] = j
			j++
		}
		
	}
	return res
}
