package practice1

import "camp/common"

func IsValid1(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := common.NewStack[byte]()

	for i := 0; i < l; i++ {
		cur := s[i]
		v, ok := pairs[cur]
		if ok {
			t := stack.Pop()
			if t != v {
				return false
			} else {
				continue
			}
		} else {
			stack.Push(cur)
		}
	}
	return stack.IsEmpty()
}
