package intcodecomputer

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

type Memory []int64

type Command struct {
	Handler  func(*Computer, []*int64) //computer pointer and args
	ArgCount int
}

type Computer struct {
	Data           Memory
	CurPos         int64
	relBase        int64
	Debug          bool
	CustomCommands map[int64]Command
	InputStack     []int64
	OutputStack    []int64
	EventFunctions map[string][]func(*Computer, map[string]interface{})
}

// list of current events:
// Output
// Input

func (c *Computer) AddEventFunction(eventName string, f func(*Computer, map[string]interface{})) {
	c.EventFunctions[eventName] = append(c.EventFunctions[eventName], f)
}

func (c *Computer) AddVirtualInput(in int64) {
	c.InputStack = append(c.InputStack, in)
}

func (c *Computer) GetNextOutput() int64 {
	out := c.OutputStack[0]
	c.OutputStack = c.OutputStack[1:]
	return out
}

func (c *Computer) AddCommand(name int64, nc Command) {
	c.CustomCommands[name] = nc
}

func (c *Computer) DebugPrint(formatter string, args ...interface{}) {
	if c.Debug {
		fmt.Printf("CMP DBG:"+strconv.FormatInt(c.CurPos, 10)+":"+formatter+"\n", args...)
	}
}

func (c *Computer) getArgPointers(count int) []*int64 {
	args := make([]*int64, count)
	for i := 0; i < count; i++ {
		switch getDigitOfPos(c.Data[c.CurPos], 2+int64(i), 1) {
		case 0:
			args[i] = &c.Data[c.Data[c.CurPos+1+int64(i)]]
		case 1:
			args[i] = &c.Data[c.CurPos+1+int64(i)]
		case 2:
			args[i] = &c.Data[c.relBase+c.Data[c.CurPos+1+int64(i)]]
		}
	}

	return args
}

func (m *Memory) LoadFromFile(filename string) {
	pattern := regexp.MustCompile("-?[0-9]+")
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	strMatches := pattern.FindAllString(string(str), -1)
	*m = make(Memory, len(strMatches)+5000)
	for k, v := range strMatches {
		(*m)[k], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Computer) callEvent(name string, data map[string]interface{}) {
	if v, ok := c.EventFunctions[name]; ok {
		for k, _ := range v {
			v[k](c, data)
		}
	}
}

func (c *Computer) Compute() {
	for {
		switch getDigitOfPos(c.Data[c.CurPos], 0, 2) {
		case 1:
			args := c.getArgPointers(3)
			c.DebugPrint("%d + %d", *args[0], *args[1])
			*args[2] = *args[0] + *args[1]
			c.CurPos += 4

		case 2:
			args := c.getArgPointers(3)
			c.DebugPrint("%d * %d", *args[0], *args[1])
			*args[2] = *args[0] * *args[1]
			c.CurPos += 4

		case 3:
			args := c.getArgPointers(1)
			evRET := int64(0)
			evUSERET := false
			c.callEvent("input", map[string]interface{}{
				"value":  &evRET,
				"enable": &evUSERET,
			})
			if evUSERET {
				*args[0] = evRET
			} else if len(c.InputStack) != 0 {
				*args[0] = c.InputStack[0]
				c.InputStack = c.InputStack[1:]
			} else {
				reader := bufio.NewReader(os.Stdin)
				fmt.Println("Please enter a value")
				in, _ := reader.ReadString('\n')
				inInt, err := strconv.ParseInt(in[:len(in)-1], 10, 64)
				if err != nil {
					panic(err)
				}
				*args[0] = inInt
			}
			c.CurPos += 2

		case 4:
			args := c.getArgPointers(1)
			//fmt.Printf("OUTPUT[%d]: %d\n", c.CurPos, *args[0])
			c.callEvent("output", map[string]interface{}{
				"value": *args[0],
			})
			c.OutputStack = append(c.OutputStack, *args[0])
			c.CurPos += 2

		case 5:
			args := c.getArgPointers(2)
			c.DebugPrint("jmp if true %d to %d", *args[0], *args[1])
			if *args[0] == 1 {
				c.DebugPrint("  jumped")
				c.CurPos = *args[1]
				continue
			}
			c.CurPos += 3

		case 6:
			args := c.getArgPointers(2)
			c.DebugPrint("jmp if false %d to %d", *args[0], *args[1])
			if *args[0] != 1 {
				c.DebugPrint("  jumped")
				c.CurPos = *args[1]
				continue
			}
			c.CurPos += 3

		case 7:
			args := c.getArgPointers(3)
			c.DebugPrint("%d less than %d ?", *args[0], *args[1])
			if *args[0] < *args[1] {
				c.DebugPrint("  true")
				*args[2] = 1
			} else {
				c.DebugPrint("  false")
				*args[2] = 0
			}
			c.CurPos += 4

		case 8:
			args := c.getArgPointers(3)
			c.DebugPrint("%d equal to %d ?", *args[0], *args[1])
			if *args[0] == *args[1] {
				c.DebugPrint("  true")
				*args[2] = 1
			} else {
				c.DebugPrint("  false")
				*args[2] = 0
			}
			c.CurPos += 4
		case 9:
			args := c.getArgPointers(1)
			c.DebugPrint("Change relBase to %d", *args[0])
			c.relBase += *args[0]
			c.CurPos += 2

		case 99:
			return

		default:
			if v, ok := c.CustomCommands[getDigitOfPos(c.Data[c.CurPos], 0, 2)]; ok {
				v.Handler(c, c.getArgPointers(v.ArgCount))
			}

			fmt.Println("unknown command \"" + strconv.FormatInt(getDigitOfPos(c.Data[c.CurPos], 0, 2), 10) + "\" at position " + strconv.FormatInt(c.CurPos, 10) + " ... dumping memory and exiting...")
			out := "CURRENT ID: " + strconv.FormatInt(c.CurPos, 10) + "\n"
			for i := 0; i < len(c.Data); i++ {
				if i%3 == 0 {
					tmp := strconv.FormatInt(int64(i), 16)
					out += "\n0x" + strconv.FormatInt(int64(i), 16) + ":"
					for j := 0; j < 6-len(tmp); j++ {
						out += " "
					}
				}
				tmp := strconv.FormatInt(c.Data[i], 10)
				spaces := 7 - len(tmp)
				out += tmp
				out += ","
				for j := 0; j < spaces; j++ {
					out += " "
				}

			}
			err := ioutil.WriteFile("dump.txt", []byte(out), 0644)
			if err != nil {
				panic("could not open dump.txt...")
			}
			fmt.Println("memory dumt at dump.txt...")
			os.Exit(1)

		}
	}
}

func NewComputer() Computer {
	var c Computer
	c.CustomCommands = make(map[int64]Command)
	c.EventFunctions = make(map[string][]func(*Computer, map[string]interface{}))
	return c
}

func GetComputerFromFile(filename string) Computer {
	c := NewComputer()
	c.Data.LoadFromFile(filename)
	return c
}
