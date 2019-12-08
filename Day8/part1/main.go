package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type layer struct {
	Data string
	Size int
}

func (l layer) getAmountOfDigits(digit int) int {
	pattern := regexp.MustCompile(strconv.Itoa(digit))
	result := pattern.FindAllString(l.Data, -1)
	fmt.Println(len(result))
	return len(result)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("please enter a filename.")
	}
	width := 25
	height := 6

	fileDataBytes, err := ioutil.ReadFile(args[1])
	fileData := string(fileDataBytes)
	if err != nil {
		panic(err)
	}

	layerCount := len(fileData) / (width * height)

	layers := make([]layer, layerCount)
	for curLayer := 0; curLayer < layerCount; curLayer++ {
		curData := fileData[(curLayer * width * height):((curLayer + 1) * width * height)]
		layers[curLayer] = layer{
			Data: curData,
			Size: (width * height),
		}
	}

	lastSmallest := 99999999
	curLayer := -1
	for k, _ := range layers {
		if layers[k].getAmountOfDigits(0) < lastSmallest {
			lastSmallest = layers[k].getAmountOfDigits(0)
			curLayer = k
		}
	}

	fmt.Println("the flag is:", layers[curLayer].getAmountOfDigits(1)*layers[curLayer].getAmountOfDigits(2))

}
