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

type memory []int64

type Command struct {
	Handler  func(*Computer, []*int64) //computer pointer and args
	ArgCount int
}

type Computer struct {
	Data           memory
	CurPos         int64
	relBase        int64
	Debug          bool
	CustomCommands map[int64]Command
}

func (c *Computer) AddCommand(name int64, nc Command) {
	c.CustomCommands[name] = nc
}

func (c *Computer) debugPrint(formatter string, args ...interface{}) {
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

func (c *Computer) LoadFromFile(filename string) {
	pattern := regexp.MustCompile("-?[0-9]+")
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	strMatches := pattern.FindAllString(string(str), -1)
	c.Data = make(memory, len(strMatches)+5000)
	for k, v := range strMatches {
		c.Data[k], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Computer) Compute() {
	for {
		switch getDigitOfPos(c.Data[c.CurPos], 0, 2) {
		case 1:
			args := c.getArgPointers(3)
			c.debugPrint("%d + %d", *args[0], *args[1])
			*args[2] = *args[0] + *args[1]
			c.CurPos += 4

		case 2:
			args := c.getArgPointers(3)
			c.debugPrint("%d * %d", *args[0], *args[1])
			*args[2] = *args[0] * *args[1]
			c.CurPos += 4

		case 3:
			args := c.getArgPointers(1)
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Please enter a value")
			in, _ := reader.ReadString('\n')
			inInt, err := strconv.ParseInt(in[:len(in)-1], 10, 64)
			if err != nil {
				panic(err)
			}
			*args[0] = inInt
			c.CurPos += 2

		case 4:
			args := c.getArgPointers(1)
			fmt.Printf("OUTPUT[%d]: %d\n", c.CurPos, *args[0])
			c.CurPos += 2

		case 5:
			args := c.getArgPointers(2)
			c.debugPrint("jmp if true %d to %d", *args[0], *args[1])
			if *args[0] == 1 {
				c.debugPrint("  jumped")
				c.CurPos = *args[1]
				continue
			}
			c.CurPos += 3

		case 6:
			args := c.getArgPointers(2)
			c.debugPrint("jmp if false %d to %d", *args[0], *args[1])
			if *args[0] != 1 {
				c.debugPrint("  jumped")
				c.CurPos = *args[1]
				continue
			}
			c.CurPos += 3

		case 7:
			args := c.getArgPointers(3)
			c.debugPrint("%d less than %d ?", *args[0], *args[1])
			if *args[0] < *args[1] {
				c.debugPrint("  true")
				*args[2] = 1
			} else {
				c.debugPrint("  false")
				*args[2] = 0
			}
			c.CurPos += 4

		case 8:
			args := c.getArgPointers(3)
			c.debugPrint("%d equal to %d ?", *args[0], *args[1])
			if *args[0] == *args[1] {
				c.debugPrint("  true")
				*args[2] = 1
			} else {
				c.debugPrint("  false")
				*args[2] = 0
			}
			c.CurPos += 4
		case 9:
			args := c.getArgPointers(1)
			c.debugPrint("Change relBase to %d", *args[0])
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

func GetComputerFromFile(filename string) Computer {
	var c Computer
	c.LoadFromFile(filename)
	return c
}
