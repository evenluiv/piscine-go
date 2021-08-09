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
		fmt.Println("File name missing")
	} else if length < 2 {
		fmt.Println("Too many arguments")
	} else if arguments[0] == "quest8.txt" {

		content, err := ioutil.ReadFile(arguments[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(content))
	}
}
