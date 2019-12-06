package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type planet struct {
	name   string
	childs []string
	parent *planet
}

type pathPart struct {
	name     string
	distance int
}

type path []pathPart

func (p planet) getOrbitCount(a *int) {
	if p.parent != nil {
		*a++
		p.parent.getOrbitCount(a)
	}
}

func (p planet) getPathToCom() path {
	out := make(path, 0)
	cur := &p
	i := 0
	for cur.parent != nil {
		fmt.Println(i, cur.parent.name)
		out = append(out, pathPart{name: cur.parent.name, distance: i})
		i++
		cur = cur.parent
	}
	fmt.Println(p.name, " : ", out)
	return out
}

func (p planet) getDistance(otherOne *planet) int {
	path2 := p.getPathToCom()
	for _, v := range otherOne.getPathToCom() {
		for _, v2 := range path2 {
			fmt.Println(v, v2)
			if v.name == v2.name {
				return v.distance + v2.distance
			}
		}
	}
	return -1
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
				name:   v[2],
				childs: []string{v[2]},
			}
			orbits[v[1]] = newone
		}
		if val, ok := orbits[v[2]]; ok {
			val.parent = orbits[v[1]]
			orbits[v[2]] = val
		} else {
			newone := &planet{
				name:   v[2],
				parent: orbits[v[1]],
			}
			orbits[v[2]] = newone
		}
	}

	out := orbits["YOU"].getDistance(orbits["SAN"])
	fmt.Println("the flag is: ", out)
}
