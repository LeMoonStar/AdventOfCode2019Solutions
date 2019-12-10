package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strings"
)

type vec struct {
	X int
	Y int
}

type field struct {
	Asteroid bool
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	pattern := regexp.MustCompile("(#|\\.)")
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	points := make(map[vec]field)
	curPos := vec{X: 0, Y: 0}
	size := vec{X: 0, Y: 0}
	for scanner.Scan() {

		matches := pattern.FindAllString(scanner.Text(), -1)
		for k, _ := range matches {
			fmt.Println(curPos)
			if curPos.X > size.X {
				size.X = curPos.X
			}
			if curPos.Y > size.Y {
				size.Y = curPos.Y
			}
			if matches[k] == "#" {
				points[curPos] = field{
					Asteroid: true,
				}
			} else {
				points[curPos] = field{
					Asteroid: false,
				}
			}
			curPos.X++
		}

		curPos.Y++
		curPos.X = 0
	}

	/*for y := 0; y <= size.Y; y++ {
		for x := 0; x <= size.X; x++ {
			if points[vec{X: x, Y: y}].Asteroid {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}*/

	/*outerFields := make([]vec, 0)
	for k, _ := range points {
		if k.X == 0 || k.Y == 0 || k.X == size.X || k.Y == size.Y {
			outerFields = append(outerFields, k)
		}
	}*/

	canSee := make(map[vec]int)
	for ownPos, _ := range points {
		if !points[ownPos].Asteroid {
			continue
		}
		anglesToDist := make(map[float64]float64)
		for side, _ := range points {
			if side == ownPos || !points[side].Asteroid {
				continue
			}
			//fmt.Println(side)
			angle := math.Atan2(float64(side.Y-ownPos.Y), float64(side.X-ownPos.X))
			anglesToDist[angle] = math.Abs(float64(side.X-ownPos.X)) + math.Abs(float64(side.Y-ownPos.Y))
		}
		canSee[ownPos] = len(anglesToDist)
		//fmt.Println(len(anglesToDist))
	}

	//lastLargestKey := vec{X: 0, Y: 0}
	lastLargestVal := 0
	for k, _ := range canSee {
		if lastLargestVal < canSee[k] {
			//lastLargestKey = k
			lastLargestVal = canSee[k]
		}
	}
	fmt.Println(lastLargestVal)
}
