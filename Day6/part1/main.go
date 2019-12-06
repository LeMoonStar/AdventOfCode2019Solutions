package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type planet struct {
	childs []string
	parent *planet
}

func (p planet) getOrbitCount(a *int) {
	if p.parent != nil {
		*a++
		p.parent.getOrbitCount(a)
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	pattern := regexp.MustCompile("([A-Z1-9]+)\\)([A-Z1-9]+)")
	str, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	strMatches := pattern.FindAllStringSubmatch(string(str), -1)

	orbits := make(map[string]*planet)
	for _, v := range strMatches {
		if val, ok := orbits[v[1]]; ok {
			val.childs = append(val.childs, v[2])
			orbits[v[1]] = val
		} else {
			newone := &planet{
				childs: []string{v[2]},
			}
			orbits[v[1]] = newone
		}
		if val, ok := orbits[v[2]]; ok {
			val.parent = orbits[v[1]]
			orbits[v[2]] = val
		} else {
			newone := &planet{
				parent: orbits[v[1]],
			}
			orbits[v[2]] = newone
		}
	}

	out := 0
	for _, v := range orbits {
		v.getOrbitCount(&out)
	}
	fmt.Println("the flag is: ", out)
}
