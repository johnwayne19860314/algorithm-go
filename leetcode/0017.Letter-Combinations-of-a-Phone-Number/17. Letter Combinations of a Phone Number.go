package leetcode

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

// 解法一 DFS
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

// 解法二 非递归
func letterCombinations_(digits string) []string {
	if digits == "" {
		return []string{}
	}
	index := digits[0] - '0'
	letter := letterMap[index]
	tmp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res, "")
		}
		for j := 0; j < len(res); j++ {
			tmp = append(tmp, res[j]+string(letter[i]))
		}
	}
	res = tmp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		tmp = res
		res = []string{}
	}
	return tmp
}

// 解法三 回溯（参考回溯模板，类似DFS）
var result []string
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]
	}
}
func DfsCombines0412(digits string) []string {
	if digits == "" {
		return res
	}
	dfsCombines0412(&digits,0,"")
	return res


}

func dfsCombines0412(digits *string, idx int, s string) {
	if idx == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[idx] 
	for _,item := range letterMap[num-'0'] {
		dfsCombines0412(digits,idx+1, s+string(item))
	}

}

// func allLetterCombines(s string) []string{

// 	strMap := make(map[string]bool)
// 	strArr := make([]string,0)
// 	for i := 0 ; i < len(s); i++ {
// 		//strArr = append(strArr, s[i:i+1])
// 		tmp := s[i:i+1]
// 		if _, ok := strMap[tmp]; !ok {
// 			strArr = append(strArr, tmp)
// 		}
// 	}
// 	mapArr := make([][]string,0)
// 	for i:= 0 ; i < len(strArr) ;i++ {
// 		tmp := strArr[i]
// 		mapArr = append(mapArr, dict[tmp])
// 		// tmpStr := ""
// 		// for _, item := range dict[tmp] {
// 		// 	tmpStr += item
// 		// 	for _, item := range dict[]
// 		// }
// 	}
	
// 	res := []string{}
// 	len, lastIdx := len(mapArr),len(mapArr)-1
// 	cursor := lastIdx
	
// 	for i := 0; i < 3;i++ {
// 		tmpStr := ""
		
// 		for cursor > 0 {
// 			l := 0
// 			for  l < 3 {
			
// 				for j := 0; j < lastIdx ; j++ {
					
// 					if j == (cursor) {
// 						tmpStr += mapArr[l][j]
// 					}else {
// 						tmpStr += mapArr[i][j]
// 					}
					
// 				}
// 				for k := 0; k < 3; k++ {
// 					tmp := tmpStr+mapArr[i][lastIdx]
// 					res = append(res, tmp)
// 				}
				
// 				l++
	
// 			}
// 		}
		
// 		cursor--
		


// 	}

// }

// func matricComb(input [][]string, res *[]string) {
// 	i,j := len(input), len(input[0])

// }