package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	length := len(os.Args)
	if length == 1 {
		fmt.Print("File name missing")
	} else if length < 2 {
		fmt.Print("Too many arguments")
	} else if arguments[0] == "quest8.txt" {

		content, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		fmt.Print(string(content))
	}
}
