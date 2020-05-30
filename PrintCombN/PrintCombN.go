package student

import (
	"fmt"

	"github.com/01-edu/z01"
)

func printi(b int, n int, a int, i int, l string) {
	a++
	for j := b + 1; j < 10-n+a; j++ {
		if a < n {
			printi(j, n, a, i, l+ConvertInt64ToString(j))
		} else {
			runeL := []rune(l)
			lenL := 0
			for range runeL {
				lenL++
			}
			for q := 0; q < lenL; q++ {
				z01.PrintRune(runeL[q])
			}
			runeJ := []rune(string(j))
			lenJ := 0
			for range runeJ {
				lenJ++
			}
			for w := 0; w < lenJ; w++ {
				z01.PrintRune(runeJ[w] + 48)
			}
			if i < 10-n {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}

func PrintCombN(n int) {
	for i := 0; i < 10; i++ {
		if n > 1 {
			printi(i, n, 1, i, ConvertInt64ToString(i))
		} else {
			fmt.Print(i)
			if i < 10-n {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
}

func ConvertInt64ToString(n int) string {
	str := ""
	str1 := ""
	if n == 0 {
		str1 = "0"
	} else if n < 0 {
		n = -n
		for n >= 1 {
			str = str + string(n%10+48)
			n = n / 10
		}
		str = str + "-"
	} else {
		for n >= 1 {
			str = str + string(n%10+48)
			n = n / 10
		}
	}
	l := 0
	for range str {
		l++
	}
	for i := l - 1; i >= 0; i-- {
		str1 = str1 + string(str[i])
	}
	if str1 != "-" {
		return str1
	}
	return "0"
}
