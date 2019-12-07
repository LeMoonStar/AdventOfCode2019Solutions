package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
)

func anyequals(args ...int) bool {
	for k, v := range args {
		for k2, v2 := range args {
			if v == v2 && k != k2 {
				return true
			}
		}
	}
	return false
}

func getDigitOfPos(in int, pos int, digitCout int) int {
	if pos == 0 {
		return in % int(math.Pow(10, float64(digitCout)))
	}
	return in / int(math.Pow(10, float64(pos))) % int(math.Pow(10, float64(digitCout)))
}

func compute(in []int, simInput []int) []int {
	inputcount := 0
	output := make([]int, 0)
	input := make([]int, len(in))
	copy(input, in)
	i := 0
	for i < len(input) {
		if getDigitOfPos(input[i], 0, 2) == 1 {
			var arg1, arg2, arg3 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else {
				arg2 = &input[i+2]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else {
				arg3 = &input[i+3]
			}
			*arg3 = *arg1 + *arg2
			i += 4
		} else if getDigitOfPos(input[i], 0, 2) == 2 {
			var arg1, arg2, arg3 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else {
				arg2 = &input[i+2]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else {
				arg3 = &input[i+3]
			}

			*arg3 = *arg1 * *arg2
			i += 4
		} else if getDigitOfPos(input[i], 0, 2) == 3 {
			var arg1 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if inputcount < len(simInput) {
				*arg1 = simInput[inputcount]
				inputcount++
			} else {
				*arg1 = 0
			}
			/*	reader := bufio.NewReader(os.Stdin)
				in, _ := reader.ReadString('\n')
				inInt, err := strconv.Atoi(in[:len(in)-1])
				if err != nil {
					panic(err)
				}
				*arg1 = inInt*/

			i += 2
		} else if getDigitOfPos(input[i], 0, 2) == 4 {
			var arg1 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			output = append(output, *arg1)
			i += 2
		} else if getDigitOfPos(input[i], 0, 2) == 5 {
			var arg1, arg2 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else {
				arg2 = &input[i+2]
			}
			if *arg1 != 0 {
				i = *arg2
			} else {
				i += 3
			}

		} else if getDigitOfPos(input[i], 0, 2) == 6 {
			var arg1, arg2 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else {
				arg2 = &input[i+2]
			}
			if *arg1 == 0 {
				i = *arg2
			} else {
				i += 3
			}
		} else if getDigitOfPos(input[i], 0, 2) == 7 {
			var arg1, arg2, arg3 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else {
				arg2 = &input[i+2]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else {
				arg3 = &input[i+3]
			}
			if *arg1 < *arg2 {
				*arg3 = 1
			}
			i += 4

		} else if getDigitOfPos(input[i], 0, 2) == 8 {
			var arg1, arg2, arg3 *int
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else {
				arg1 = &input[i+1]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else {
				arg2 = &input[i+2]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else {
				arg3 = &input[i+3]
			}
			if *arg1 == *arg2 {
				*arg3 = 1
			} else {
				*arg3 = 0
			}
			i += 4

		} else if getDigitOfPos(input[i], 0, 2) == 99 {

			return output
		} else {
			for _, v := range input {
				print(v, ", ")
			}
			panic("invalid command")
		}
	}
	return output
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	pattern := regexp.MustCompile("-?[0-9]+")
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
	curHighest := 0
	for i := 0; i < 4; i++ {
		var out []int
		for i2 := 0; i2 < 5; i2++ {
			for i3 := 0; i3 < 5; i3++ {
				for i4 := 0; i4 < 5; i4++ {
					for i5 := 0; i5 < 5; i5++ {
						for i6 := 0; i6 < 5; i6++ {
							if !anyequals(i2, i3, i4, i5, i6) {
								out = compute(matches, []int{i2, 0})
								out = compute(matches, []int{i3, out[0]})
								out = compute(matches, []int{i4, out[0]})
								out = compute(matches, []int{i5, out[0]})
								out = compute(matches, []int{i6, out[0]})
								if out[0] > curHighest {
									curHighest = out[0]
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("the flag is:", curHighest)

}
