package student

func SplitWhiteSpaces(str string) []string {
	stringa := ""
	len := 0
	prev := false
	for _, s := range str {
		if (s == ' ' || s == '\n' || s == '\t') && !prev {
			prev = true
			len++
		} else {
			prev = false
		}
	}
	len++

	arr := make([]string, len)

	count := 0
	for _, s := range str {

		if s == ' ' || s == '\n' || s == '\t' {
			length := 0
			for i := range stringa {
				length = i + 1
			}
			if length == 0 {
				continue
			}
			arr[count] = stringa
			count++
			stringa = ""
			continue
		}
		stringa += string(s)
	}

	length := 0
	for i := range stringa {
		length = i + 1
	}
	if length != 0 {
		arr[count] = stringa
	}
	newSize := 0
	for i := range arr {
		if arr[i] != "" {
			newSize++
		}
	}
	cnt := 0
	newArr := make([]string, newSize)
	for i := range arr {
		if arr[i] != "" {
			cnt++
			newArr[cnt-1] = arr[i]
		}
	}
	return newArr
}
