package _4_stack_queue

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	l := len(s)
	//长度如果不是偶数，那肯定配对不了
	if l%2 != 0 {
		return false
	}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}

	for i := 0; i < l; i++ {
		pairA, ok := pairs[s[i]]
		if ok {
			stackLen := len(stack)
			if stackLen == 0 || stack[stackLen-1] != pairA {
				return false
			} else {
				stack = stack[:stackLen-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
