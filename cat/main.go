package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	lenArgs := 0
	for range args {
		lenArgs++
	}
	if lenArgs == 0 {
		io.Copy(os.Stdout, os.Stdin)
	}
	for _, file_name := range args {
		file, err := os.Open(file_name)

		if err != nil {
			for _, s := range err.Error() {
				z01.PrintRune(s)
			}
			z01.PrintRune(10)
			return
		}

		b, err := ioutil.ReadAll(file)
		for _, s := range string(b) {
			z01.PrintRune(s)
		}
		if err != nil {
			for _, s := range err.Error() {
				z01.PrintRune(s)
			}
			z01.PrintRune('\n')
		}
	}
}
