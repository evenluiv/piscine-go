package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) == 1 {
		return
	} else {
		os.Args = os.Args[1:]

		for _, v := range os.Args {
			file, err := os.Open(v)
			if err != nil {
				printStr(err.Error())
				return
			} else {
				data := make([]byte, 443)
				file.Read(data)
				if len(os.Args) == 1 {
					printStr(string(data))
				} else {
					printStr(string(data))
				}
			}
		}
	}
}
