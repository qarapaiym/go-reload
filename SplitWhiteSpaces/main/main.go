package main

import (
	"fmt"

	student ".."
)

func main() {
	str := "Hello how are you?"
	fmt.Println(str)
	fmt.Println(student.SplitWhiteSpaces(str))
}
