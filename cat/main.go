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
	userinput, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	} else {
		printStr(string(userinput))
	}

	if len(os.Args) == 1 {
		return
	} else {
		os.Args = os.Args[1:]

		for _, v := range os.Args {
			file, err := os.Open(v)
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				os.Exit(1)
				return
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					printStr("ERROR: ")
					printStr(err.Error())
					os.Exit(1)
					break
				} else {
					printStr(string(data))
				}
			}
		}
	}
}
