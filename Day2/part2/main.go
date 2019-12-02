package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func compute(in []int) int {
	input := make([]int, len(in))
	copy(input, in)
	i := 0
	for i < len(input) {
		if input[i] == 1 {
			arg1 := input[input[i+1]]
			arg2 := input[input[i+2]]
			arg3 := &input[input[i+3]]
			*arg3 = arg1 + arg2
			i += 4
		} else if input[i] == 2 {
			arg1 := input[input[i+1]]
			arg2 := input[input[i+2]]
			arg3 := &input[input[i+3]]
			*arg3 = arg1 * arg2
			i += 4
		} else if input[i] == 99 {

			/*for _, v := range input {
				print(v, ", ")
			}
			print("\n")*/

			return input[0]
		} else {
			for _, v := range input {
				print(v, ", ")
			}
			panic("invalid command")
		}
	}
	return 0
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	pattern := regexp.MustCompile("[0-9]+")
	str, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	strMatches := pattern.FindAllString(string(str), -1)
	matches := make([]int, len(strMatches))
	for k, v := range strMatches {
		matches[k], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	for noun := 0; noun < len(matches); noun++ {
		for verb := 0; verb < len(matches); verb++ {
			matches[1] = noun
			matches[2] = verb
			out := compute(matches)
			if out == 19690720 {
				fmt.Print("noun:", noun, " verb:", verb)
				break
			}
		}
	}

}
