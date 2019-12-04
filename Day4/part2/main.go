package main

import (
	"fmt"
	"os"
	"strconv"
)

func checkNumber(num int, input1 int, input2 int) bool {
	if num > input2 || num < input1 {
		return false
	}
	digits := make([]int, 0)
	tmp := num
	for i := 1; i <= 6; i++ {
		digits = append(digits, tmp%10)
		tmp = tmp / 10
	}
	someFlag := false
	last := 10
	lastlast := 10
	for k, v := range digits {
		if v > last {
			return false
		}
		if v == last {
			if v != lastlast {
				if k < 5 {
					if v != digits[k+1] {
						someFlag = true
					}
				} else {
					someFlag = true
				}
			}

		}
		lastlast = last
		last = v
	}
	return someFlag
}

func main() {
	args := os.Args
	if len(args) <= 2 {
		fmt.Println("usage:\n", args[0], " min max")
		return
	}
	input1, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	input2, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	foundCount := 0
	i := input1
	for true {
		i++
		if i > input2 {
			break
		}
		if checkNumber(i, input1, input2) {
			foundCount++
		}
	}
	fmt.Println("found :", foundCount, " matching numbers")
}
