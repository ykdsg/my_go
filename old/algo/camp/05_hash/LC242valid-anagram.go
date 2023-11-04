package _5_hash

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
