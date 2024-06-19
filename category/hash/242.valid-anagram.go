package hash

func IsAnagram(s string, t string) bool {
	m := make(map[string]int)
	for _, c := range s {
		m[string(c)] += 1
	}
	for _, c := range t {
		m[string(c)] -= 1
	}
	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}

func IsAnagramByArray(s string, t string) bool {
	m := [26]int{}
	for _, c := range s {
		m[c-'a'] += 1
	}
	for _, c := range t {
		m[c-'a'] -= 1
	}
	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
