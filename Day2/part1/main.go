package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

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

	matches[1] = 12
	matches[2] = 2

	i := 0
	for i < len(matches) {
		if matches[i] == 1 {
			arg1 := matches[matches[i+1]]
			arg2 := matches[matches[i+2]]
			arg3 := &matches[matches[i+3]]
			*arg3 = arg1 + arg2
			i += 4
		} else if matches[i] == 2 {
			arg1 := matches[matches[i+1]]
			arg2 := matches[matches[i+2]]
			arg3 := &matches[matches[i+3]]
			*arg3 = arg1 * arg2
			i += 4
		} else if matches[i] == 99 {

			/*for _, v := range matches {
				print(v, ", ")
			}
			print("\n")*/

			fmt.Println("The flag is: ", matches[0])
			break
		} else {
			for _, v := range matches {
				print(v, ", ")
			}
			panic("invalid command")
		}
	}
}
