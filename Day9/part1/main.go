package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
)

func getDigitOfPos(in int64, pos int64, digitCout int64) int64 {
	if pos == 0 {
		return in % int64(math.Pow(10, float64(digitCout)))
	}
	return in / int64(math.Pow(10, float64(pos))) % int64(math.Pow(10, float64(digitCout)))
}

func compute(in []int64) int64 {
	relBase := int64(0)
	input := make([]int64, len(in)+5000)
	copy(input, in)
	i := int64(0)
	for i < int64(len(input)) {
		if getDigitOfPos(input[i], 0, 2) == 1 {
			var arg1, arg2, arg3 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else if getDigitOfPos(input[i], 3, 1) == 1 {
				arg2 = &input[i+2]
			} else {
				arg2 = &input[input[i+2]+relBase]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else if getDigitOfPos(input[i], 4, 1) == 1 {
				arg3 = &input[i+3]
			} else {
				arg3 = &input[input[i+3]+i]
			}
			*arg3 = *arg1 + *arg2
			i += 4
		} else if getDigitOfPos(input[i], 0, 2) == 2 {
			var arg1, arg2, arg3 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else if getDigitOfPos(input[i], 3, 1) == 1 {
				arg2 = &input[i+2]
			} else {
				arg2 = &input[input[i+2]+relBase]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else if getDigitOfPos(input[i], 4, 1) == 1 {
				arg3 = &input[i+3]
			} else {
				arg3 = &input[input[i+3]+i]
			}

			*arg3 = *arg1 * *arg2
			i += 4
		} else if getDigitOfPos(input[i], 0, 2) == 3 {
			var arg1 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			reader := bufio.NewReader(os.Stdin)
			in, _ := reader.ReadString('\n')
			inInt, err := strconv.ParseInt(in[:len(in)-1], 10, 64)
			if err != nil {
				panic(err)
			}
			*arg1 = inInt
			i += 2
		} else if getDigitOfPos(input[i], 0, 2) == 4 {
			var arg1 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}

			fmt.Println("OUTPUT[", i, "]: ", *arg1)
			i += 2
		} else if getDigitOfPos(input[i], 0, 2) == 5 {
			var arg1, arg2 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else if getDigitOfPos(input[i], 3, 1) == 1 {
				arg2 = &input[i+2]
			} else {
				arg2 = &input[input[i+2]+relBase]
			}
			if *arg1 != 0 {
				i = *arg2
			} else {
				i += 3
			}

		} else if getDigitOfPos(input[i], 0, 2) == 6 {
			var arg1, arg2 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else if getDigitOfPos(input[i], 3, 1) == 1 {
				arg2 = &input[i+2]
			} else {
				arg2 = &input[input[i+2]+relBase]
			}
			if *arg1 == 0 {
				i = *arg2
			} else {
				i += 3
			}
		} else if getDigitOfPos(input[i], 0, 2) == 7 {
			var arg1, arg2, arg3 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else if getDigitOfPos(input[i], 3, 1) == 1 {
				arg2 = &input[i+2]
			} else {
				arg2 = &input[input[i+2]+relBase]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else if getDigitOfPos(input[i], 4, 1) == 1 {
				arg3 = &input[i+3]
			} else {
				arg3 = &input[input[i+3]+i]
			}
			if *arg1 < *arg2 {
				*arg3 = 1
			} else {
				*arg3 = 0
			}
			i += 4

		} else if getDigitOfPos(input[i], 0, 2) == 8 {
			var arg1, arg2, arg3 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			if getDigitOfPos(input[i], 3, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg2 = &input[input[i+2]]
			} else if getDigitOfPos(input[i], 3, 1) == 1 {
				arg2 = &input[i+2]
			} else {
				arg2 = &input[input[i+2]+relBase]
			}
			if getDigitOfPos(input[i], 4, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg3 = &input[input[i+3]]
			} else if getDigitOfPos(input[i], 4, 1) == 1 {
				arg3 = &input[i+3]
			} else {
				arg3 = &input[input[i+3]+i]
			}
			if *arg1 == *arg2 {
				*arg3 = 1
			} else {
				*arg3 = 0
			}
			i += 4

		} else if getDigitOfPos(input[i], 0, 2) == 9 {
			var arg1 *int64
			if getDigitOfPos(input[i], 2, 1) == 0 {
				//fmt.Println("MODE ", i, " 0")
				arg1 = &input[input[i+1]]
			} else if getDigitOfPos(input[i], 2, 1) == 1 {
				arg1 = &input[i+1]
			} else {
				arg1 = &input[input[i+1]+i]
			}
			relBase += *arg1
			i += 2

		} else if getDigitOfPos(input[i], 0, 2) == 99 {

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
	pattern := regexp.MustCompile("-?[0-9]+")
	str, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	strMatches := pattern.FindAllString(string(str), -1)
	matches := make([]int64, len(strMatches))
	for k, v := range strMatches {
		matches[k], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	compute(matches)
}
