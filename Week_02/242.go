package main

// 字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	origin := make(map[byte]int)
	new := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		addItem(origin, s[i])
		addItem(new, t[i])
	}

	for k, v := range origin {
		cnt, ok := new[k]

		if !ok || cnt != v {
			return false
		}
	}

	return true
}

func addItem(m map[byte]int, b byte) {
	cnt, ok := m[b]
	if ok {
		m[b] = cnt + 1
	} else {
		m[b] = 1
	}
	return
}
