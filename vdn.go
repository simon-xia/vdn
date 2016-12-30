package vdn

// rfc1035 2.3.4. Size limits
const (
	maxLableSize  = 63
	maxDomainSize = 254
	minPartCnt    = 2
)

// IsValidDomain check whether a domain is valid
// https://tools.ietf.org/html/rfc1035
// https://tools.ietf.org/html/rfc3696
func IsValidDomain(s string) bool {
	l := len(s)
	if l == 0 || l > maxDomainSize || l == maxDomainSize && s[l-1] != '.' {
		return false
	}

	last := byte('.')
	ok := false // Ok once we've seen a letter.
	partlen := 0
	partCnt := 1
	for i := l - 1; i >= 0; i-- {
		c := s[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			ok = true
			partlen++
		case '0' <= c && c <= '9':
			partlen++
		case c == '-':
			// Byte after dash cannot be dot.
			if last == '.' {
				return false
			}
			partlen++
		case c == '.':
			// Byte after dot cannot be dot, dash.
			if last == '-' || last == '.' {
				return false
			}
			if partlen > maxLableSize || partlen == 0 {
				return false
			}
			if partCnt == 1 {
				if _, ok := TLDs[s[i+1:]]; !ok {
					return false
				}
			}
			partlen = 0
			partCnt++
		}
		last = c
	}

	if partCnt < minPartCnt || last == '.' || last == '-' || partlen > maxLableSize {
		return false
	}

	return ok
}
