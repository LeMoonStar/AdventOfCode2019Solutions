package main

import (
	"bufio"
	"fmt"
	"math"
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

	crosses := make([]vec2, 0)
	beenHere := make(map[vec2]bool)

	for k, _ := range instructions {
		//fmt.Println("Wire change")
		var curPos vec2 = vec2{x: 0, y: 0}
		for _, v := range instructions[k] {
			if v.action == 0 { //up
				//fmt.Println("Up ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.y++
					if beenHere[curPos] && k != 0 {
						crosses = append(crosses, curPos)
					}
					beenHere[curPos] = true

				}
			} else if v.action == 1 { //down
				//fmt.Println("Down ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.y--
					if beenHere[curPos] && k != 0 {
						crosses = append(crosses, curPos)
					}
					beenHere[curPos] = true

				}
			} else if v.action == 2 { //right
				//fmt.Println("Right ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.x++
					if beenHere[curPos] && k != 0 {
						crosses = append(crosses, curPos)
					}
					beenHere[curPos] = true

				}
			} else if v.action == 3 { //left
				//fmt.Println("Left ", v.value)
				for i := 0; i < v.value; i++ {
					curPos.x--
					if beenHere[curPos] && k != 0 {
						crosses = append(crosses, curPos)
					}
					beenHere[curPos] = true

				}
			}
			//fmt.Println("newPos: ", curPos.x, ",", curPos.y)
		}
	}

	distances := make(map[vec2]float64)
	for _, v := range crosses {
		fmt.Println("cross: ", v.x, ",", v.y)
		distances[v] = math.Abs(float64(v.x)) + math.Abs(float64(v.y))
		fmt.Println("dis: ", distances[v])
	}

	cursmallest := float64(9999999)
	for _, v := range distances {
		if v < cursmallest {
			cursmallest = v
		}
	}
	fmt.Printf("smallest: %f\n", cursmallest)

}
