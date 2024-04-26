package leetcode

func isValid(s string) bool {
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

type Stack struct {
	val  string
	next *Stack
}

func stringToStack(s string) (uint,*Stack) {
	head := Stack{}
	tmp := &head
	var n uint
	for _, r := range s {
		//fmt.Println()
		tmp.val = string(r)
		tmp.next = &Stack{}
		tmp = tmp.next
		n++
	}
	return n, &head
}

func findMatchWithHead(s *Stack) (bool,*Stack) {
	tmp := s
	var preNode *Stack
	if isMatch(s.val, s.next.val) {
		s = s.next.next
		return true, s
	} else {
		tmp = s.next
		for tmp.next.next != nil {
			preNode = tmp.next
			tmp = tmp.next.next
			if isMatch(s.val, tmp.val) {
				s = s.next
				preNode.next = tmp.next
				return true,s
			}
		}
	}
	return false,s
}


func isParenthesesValid(s string) bool {
	level, stack := stringToStack(s)
	if level%2 != 0 {
		return false
	}
	head := stack
	res := false
	for head.next != nil {
		res,head = findMatchWithHead(head)
		if !res {
			return false
		}
	}
	return true

}

func isMatch(a, b string) bool {
	if a == "{" && b == "}" {
		return true
	} else if a == "[" && b == "]" {
		return true
	} else if a == "(" && b == ")" {
		return true
	}
	return false
}
