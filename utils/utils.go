package utils

func RuneAt(s string, idx int) rune {
	rs := []rune(s)
	if idx >= len(rs) {
		return 0
	}
	return rs[idx]
}
