package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var myBad error

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

func parseData(sourceData []string) []int {
	var retVal []int
	dlen := len(sourceData)
	for c := 0; c < dlen; c++ {
		crabData := strings.Split(sourceData[c], ",")
		for fd := 0; fd < len(crabData); fd++ {
			addTo, myBad := strconv.Atoi(crabData[fd])
			checkErr(myBad)
			retVal = append(retVal, addTo)
		}
	}
	return retVal
}

func intArray(inVal []int) []int {
	var retVal []int
	for c := 0; c < len(inVal); c++ {
		retVal = append(retVal, 0)
	}
	return retVal
}

func moveCrabs(inVal []int) int {
	var retVal int
	var distCalc int
	var finalDist []int
	var crabDist []int

	crabDist = intArray(crabDist)
	for c := 0; c < len(inVal); c++ {
		distCalc = 0
		for d := 0; d < len(inVal); d++ {
			if inVal[c] > inVal[d] {
				distCalc = distCalc + (inVal[c] - inVal[d])
			} else if inVal[c] < inVal[d] {
				distCalc = distCalc + (inVal[d] - inVal[c])
			}
		}
		crabDist = append(crabDist, distCalc)
		sort.Ints(crabDist)
		finalDist = append(finalDist, crabDist[0])
	}
	sort.Ints(finalDist)
	retVal = finalDist[0]
	return retVal
}

func main() {

	var crabsPos []int
	var minDistance int
	var fname string
	fileList := []string{"sample", "input"}

	for fn := 0; fn < len(fileList); fn++ {
		fname = fileList[fn]
		sourceData := readData(fname + ".txt")
		crabsPos = parseData(sourceData)
		minDistance = moveCrabs(crabsPos)
		fmt.Println("Minimum Move Distance for", fname, "is", minDistance)
	}

}
