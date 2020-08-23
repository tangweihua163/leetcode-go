package leetcode_go

func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)

	m := make(map[int]bool, ls*lp)

	return match1(s, ls, p, lp, m)
}

func match1(s string, i int, p string, j int, m map[int]bool) bool {

	if _, ok := m[i*len(p)+j]; ok {
		return m[i*len(p)+j]
	}

	if i == 0 && j == 0 {
		return true
	}

	if i == 0 && j == 1 {
		return false
	}

	if i > 0 && j == 0 {
		return false
	}

	if i > 0 && j == 1 {
		if (p[0] == '*') || !(p[0] == '.' || p[0] == s[i-1]) {
			m[i*len(p)+j] = false
			return false
		}
		b := match1(s, i-1, p, 0, m)
		m[i*len(p)+j] = b
		return b
	}

	if j >= 2 {
		if p[j-1] == '.' {
			if i == 0 {
				m[i*len(p)+j] = false
				return false
			}
			b := match1(s, i-1, p, j-1, m)
			m[i*len(p)+j] = b
			return b
		}
		if p[j-1] != '*' {
			if i == 0 {
				m[i*len(p)+j] = false
				return false
			}
			if !(p[j-1] == s[i-1]) {
				m[i*len(p)+j] = false
				return false
			}
			b := match1(s, i-1, p, j-1, m)
			m[i*len(p)+j] = b
			return b
		}
		if p[j-2] == '.' {
			for k := i; k >= 0; k-- {
				b := match1(s, k, p, j-2, m)
				if b {
					m[i*len(p)+j] = true
					return true
				}
			}
			m[i*len(p)+j] = false
			return false
		} else {
			for k := i; k == i || k >= 0 && s[k] == p[j-2]; k-- {
				b := match1(s, k, p, j-2, m)
				if b {
					m[i*len(p)+j] = true
					return true
				}
			}
			m[i*len(p)+j] = false
			return false
		}
	}

	return false
}
