package renderer

func trim(s string, width int) string {
	if width < 0 {
		return ""
	}
	minLen := width
	if minLen > len(s) {
		minLen = len(s)
	}
	return s[:minLen]
}
