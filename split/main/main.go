package main

import (
	"fmt"

	student ".."
)

func main() {
	str := "HAHAHelloHAhowHAareHAyou?HAHA"
	fmt.Println(student.Split(str, "HA"))
}
