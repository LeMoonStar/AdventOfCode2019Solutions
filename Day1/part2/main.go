package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcFuelUsage(fuel int) int {
	val := fuel/3 - 2
	if val > 0 {
		val += calcFuelUsage(val)
	}

	if val < 0 {
		val = 0
	}
	fmt.Println(val)
	return val
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	f, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	token := 0

	for scanner.Scan() {
		currLine, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		token += calcFuelUsage(currLine)
	}
	fmt.Println("The token is: " + strconv.Itoa(token))
}
