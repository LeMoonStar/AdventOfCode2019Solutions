package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	f, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	token := 0

	for scanner.Scan() {
		currLine, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		val := currLine/3 - 2
		token += val
	}
	fmt.Println("The token is: " + strconv.Itoa(token))
}
