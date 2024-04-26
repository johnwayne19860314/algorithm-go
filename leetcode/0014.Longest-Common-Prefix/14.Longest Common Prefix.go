package leetcode

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	// get the common prefix two by two from left to right
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			// if this condition happens then move on to the next element
			// and cut off the prefix with index 0 and j
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}

func getLonggestCommonPrefix0405(input []string) string {

	firstEle := input[0]
	i:=0
	moveOnFlag := true
	for i< len(firstEle) && moveOnFlag {
		for _ , ele := range input[1:] {
			// if i reached to the end of any element or notEqual happens
			// then break the loop 
			// set the flag as false to break the outerloop and break statement to
			// break the inner loop
			if  i == len(ele) ||  firstEle[i] != ele[i] {
				moveOnFlag = false
				break
			}
		}
		if moveOnFlag{
			i++
		}
	}

	res := firstEle[0:i]
	return res
}
