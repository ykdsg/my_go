package _5_hash

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c-'a']++
		}
		m[key] = append(m[key], str)
	}
	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
