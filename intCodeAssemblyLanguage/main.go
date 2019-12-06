package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

var (
	outputFileName string
	filename       string
	dbg            bool
	commands       map[string]func([]Word) []int
	refs           map[string]int
)

func dbgPrint(out ...interface{}) bool {
	if dbg {
		fmt.Println(out...)
	}
	return dbg
}

func main() {
	Args := os.Args
	for i := 0; i < len(Args); i++ {
		v := Args[i]
		if v[0] == '-' {
			flag := v[1:]
			switch flag {
			case "dbg":
				dbg = true
			case "o":
				i++
				if len(Args) > i {
					outputFileName = Args[i]
				} else {
					fmt.Printf("-o needs a filename.\n")
					os.Exit(1)
				}

			default:
				fmt.Printf("Flag \"%s\" is not reconized.\n", flag)
				os.Exit(1)
			}
		} else {
			filename = v
		}
	}

	WordPattern := regexp.MustCompile("\\S+")
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ERROR: Could not open file \"%s\"\n", filename)
		os.Exit(1)
	}
	Words := WordPattern.FindAllString(string(fileData), -1)

	Output := make([]int, 0)
	setCommands()
	fmt.Print("Compiling file")
	for i := 0; i < len(Words); {
		fmt.Print(".")
		for k, _ := range commands {
			if k == Words[i] {
				Output = append(Output, commands[Words[i]](convStrSlicetoWordSlice(Words[i:]))...)
			}
		}
		i++
	}
	fmt.Println("\nDONE!")
	outStr := ""
	for k, v := range Output {
		outStr += strconv.Itoa(v)
		if k != len(Output)-1 {
			outStr += ", "
		}
	}
	ioutil.WriteFile(outputFileName, []byte(outStr), 0644)
}

func setCommands() {
	refs = make(map[string]int)
	commands = make(map[string]func([]Word) []int)
	commands["add"] = add
	commands["multi"] = multi
	commands["input"] = input
	commands["output"] = output
	commands["END"] = end
}
