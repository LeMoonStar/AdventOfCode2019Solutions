package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	action uint8
	value  int
}
type vec2 struct {
	x int
	y int
}
type Crossing struct {
	pos       vec2
	travelled []int
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	pattern := regexp.MustCompile("([A-Z][0-9]+)")
	f, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	instructions := make([][]instruction, 0)

	for scanner.Scan() {
		strMatches := pattern.FindAllString(string(scanner.Text()), -1)
		localInstructions := make([]instruction, 0)
		for _, v := range strMatches {
			value, err := strconv.Atoi(v[1:])
			if err != nil {
				panic(err)
			}
			var action uint8 = 255
			if v[0] == 'U' {
				action = 0
			} else if v[0] == 'D' {
				action = 1
			} else if v[0] == 'R' {
				action = 2
			} else if v[0] == 'L' {
				action = 3
			}
			localInstructions = append(localInstructions, instruction{
				action: action,
				value:  value,
			})
		}
		instructions = append(instructions, localInstructions)
	}

	crosses := make([]Crossing, 0)
	beenHere := make(map[vec2]int)
	blockCrossing := make(map[vec2]bool)

	for k, _ := range instructions {
		fmt.Println("Wire change")
		travelled := 0
		var curPos vec2 = vec2{x: 0, y: 0}
		for _, v := range instructions[k] {
			if v.action == 0 { //up
				fmt.Println("Up ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.y++
					travelled++
					if beenHere[curPos] != 0 && k != 0 && !blockCrossing[curPos] {
						crosses = append(crosses, Crossing{pos: curPos, travelled: []int{beenHere[curPos], travelled}})
						blockCrossing[curPos] = true
					}
					beenHere[curPos] = travelled

				}
			} else if v.action == 1 { //down
				fmt.Println("Down ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.y--
					travelled++
					if beenHere[curPos] != 0 && k != 0 && !blockCrossing[curPos] {
						crosses = append(crosses, Crossing{pos: curPos, travelled: []int{beenHere[curPos], travelled}})
						blockCrossing[curPos] = true
					}
					beenHere[curPos] = travelled

				}
			} else if v.action == 2 { //right
				fmt.Println("Right ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.x++
					travelled++
					if beenHere[curPos] != 0 && k != 0 && !blockCrossing[curPos] {
						crosses = append(crosses, Crossing{pos: curPos, travelled: []int{beenHere[curPos], travelled}})
						blockCrossing[curPos] = true
					}
					beenHere[curPos] = travelled

				}
			} else if v.action == 3 { //left
				fmt.Println("Left ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.x--
					travelled++
					if beenHere[curPos] != 0 && k != 0 && !blockCrossing[curPos] {
						crosses = append(crosses, Crossing{pos: curPos, travelled: []int{beenHere[curPos], travelled}})
						blockCrossing[curPos] = true
					}
					beenHere[curPos] = travelled

				}
			}
			fmt.Println("newPos: ", curPos.x, ",", curPos.y, " travelled: ", travelled)
		}
	}

	combSteps := make(map[vec2]int)
	for _, v := range crosses {
		combSteps[v.pos] = v.travelled[0] + v.travelled[1]
	}

	cursmallest := 999999999
	for _, v := range combSteps {
		if v < cursmallest {
			cursmallest = v
		}
	}
	fmt.Printf("smallest: %d\n", cursmallest)

}
