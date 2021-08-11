package main

import (
	"io/ioutil"
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
			userinsert, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				printStr("\n")
				os.Exit(1)
				return
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					printStr("ERROR: ")
					printStr(err.Error())
					printStr("\n")
					os.Exit(1)
					break
				} else {
					printStr(string(userinsert))
					printStr(string(data))
				}
			}
		}
	}
}
