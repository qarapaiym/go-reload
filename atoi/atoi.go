package student

func overflow(s string, maxint string) bool {
	if len(s) > len(maxint) {
		return true
	} else if len(s) == len(maxint) {
		for i := range s {
			if s[i] > maxint[i] {
				return true
			}
		}
	}
	return false
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	str := []rune(s)
	var sign int = 1

	switch {
	case str[0] == '+':
		sign = 1
		str = str[1:]
		if overflow(string(str), "9223372036854775807") {
			return 0
		}
	case str[0] == '-':
		sign = -1
		str = str[1:]
		if overflow(string(str), "9223372036854775808") {
			return 0
		}
	default:
		if overflow(string(str), "9223372036854775807") {
			return 0
		}
	}

	var sum int
	for _, n := range str {
		if n < 48 || n > 57 {
			return 0
		}
		sum = sum*10 + (int(n) - 48)
	}

	return sum * sign
}
