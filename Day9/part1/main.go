package main

import (
	"os"

	"github.com/LeMoonStar/AdventOfCode2019Solutions/intcodecomputer"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename with the inputs")
	}
	computer := intcodecomputer.GetComputerFromFile(args[1])
	computer.Compute()
}
