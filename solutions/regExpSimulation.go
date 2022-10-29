package solutions

func IsMatch(s string, p string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != p[i] {
			if p[i] == '.' {
				continue
			}
			if p[i] == '*' {
				if s[i-1] != s[i] && s[i-1] != '.' {
					return false
				}
			}
		}
	}
	return true
}
