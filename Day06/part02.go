package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var myBad error
var fishSchool []int

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readData(inputFile string) []string {
	var retVal []string
	readFile, myBad := os.Open(inputFile)
	checkErr(myBad)
	defer readFile.Close()
	inFile := bufio.NewScanner(readFile)
	for inFile.Scan() {
		retVal = append(retVal, inFile.Text())
	}
	return retVal
}

func parseData(sourceData []string, retVal []int) []int {
	dlen := len(sourceData)
	for c := 0; c < dlen; c++ {
		fishData := strings.Split(sourceData[c], ",")
		for fd := 0; fd < len(fishData); fd++ {
			addTo, myBad := strconv.Atoi(fishData[fd])
			checkErr(myBad)
			retVal[addTo] = retVal[addTo] + 1
		}
	}
	return retVal
}

func spawn(inVal []int) []int {
	var retVal []int
	retVal = initSchool(retVal)
	for fishCount := 0; fishCount < 9; fishCount++ {
		if fishCount == 0 {
			retVal[8] = inVal[0]
			inVal[7] = inVal[0]+inVal[7]
		} else {
			retVal[fishCount-1] = inVal[fishCount]
		}
	}
	return retVal
}

func initSchool(retVal []int) []int {
	for sl := 0; sl < 9; sl++ {
		retVal = append(retVal, 0)
	}
	return retVal
}

func schoolSum(inVal []int) int {
	retVal := 0
	for i := 0; i < 9; i++ {
		retVal = retVal + inVal[i]
	}
	return retVal
}

func main() {
	spawnCycle := 256

	// fname := "sample.txt"
	fname := "input.txt"

	fishSchool := initSchool(fishSchool)
	sourceData := readData(fname)
	fishSchool = parseData(sourceData, fishSchool)
	fmt.Println("Initial School Size is", schoolSum(fishSchool))
	for sc := 0; sc < spawnCycle; sc++ {
		fishSchool = spawn(fishSchool)
	}
	fmt.Println("Final Schools Size is", schoolSum(fishSchool))
}
