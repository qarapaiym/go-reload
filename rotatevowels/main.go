package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var str string
	var tempRune rune
	argsLen := 0
	for i := range os.Args {
		argsLen = i
	}
	for i := range os.Args {
		if i > 0 {
			str += os.Args[i]
		}
		if i > 0 && i < argsLen {
			str += " "
		}
	}
	res := []rune(str)
	strLen := StrLen(str) - 1
	for i := 0; i < strLen; i++ {
		if IsVowel(res[i]) {
			for !IsVowel(res[strLen]) {
				strLen--
			}
			if i < strLen {
				tempRune = res[i]
				res[i] = res[strLen]
				res[strLen] = tempRune
				strLen--
			}
		}
	}
	for v := range res {
		z01.PrintRune(res[v])
	}
	z01.PrintRune('\n')
}

func StrLen(str string) int {
	runes := []rune(str)
	var count int = 0
	for i := range runes {
		count++
		i = i
	}
	return count
}

func IsVowel(r rune) bool {
	if r == 'a' || r == 'A' || r == 'e' || r == 'E' || r == 'i' || r == 'I' || r == 'o' || r == 'O' || r == 'u' || r == 'U' {
		return true
	}
	return false
}
