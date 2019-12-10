package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type vec struct {
	X int
	Y int
}

type field struct {
	Asteroid bool
}

type someStruct struct {
	Dist float64
	Pos  vec
}

func main() {
	args := os.Args
	if len(args) < 4 {
		panic("please enter a filename with the inputs followed by X and Y of your station")
	}
	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	statx, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}
	staty, err := strconv.Atoi(args[3])
	if err != nil {
		panic(err)
	}
	Station := vec{X: statx, Y: staty}

	pattern := regexp.MustCompile("(#|\\.)")
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	points := make(map[vec]field)
	curPos := vec{X: 0, Y: 0}
	size := vec{X: 0, Y: 0}
	for scanner.Scan() {

		matches := pattern.FindAllString(scanner.Text(), -1)
		for k, _ := range matches {
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

	anglesToDist := make(map[float64][]someStruct)
	for side, _ := range points {
		if side == Station || !points[side].Asteroid {
			continue
		}
		angle := math.Atan2(float64(side.Y-Station.Y), float64(side.X-Station.X))
		anglesToDist[angle] = append(anglesToDist[angle], someStruct{Dist: math.Abs(float64(side.X-Station.X)) + math.Abs(float64(side.Y-Station.Y)), Pos: side})
	}

	//SORT anglesToDis; Dist = 0 = lowest, 1 = highest
	for k, _ := range anglesToDist {
		sort.Slice(anglesToDist[k], func(i, j int) bool {
			return anglesToDist[k][i].Dist > anglesToDist[k][j].Dist
		})
	}

	i := 0
	for {
		for k, _ := range anglesToDist {
			if i == 200 {
				fmt.Println("the flag is:", anglesToDist[k][0].Pos.X*100+anglesToDist[k][0].Pos.Y)
				os.Exit(0)
			}
			anglesToDist[k] = anglesToDist[k][1:]
			i++
		}
	}

}
