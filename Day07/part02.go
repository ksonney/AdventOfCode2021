package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
	"sort"
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
	dlen := len(sourceData)
	var retVal []int

	for c := 0; c < dlen; c++ {
		crabData := strings.Split(sourceData[c], ",")
		for fd := 0; fd < len(crabData); fd++ {
			addTo, myBad := strconv.Atoi(crabData[fd])
			checkErr(myBad)
			retVal = append(retVal, addTo)
		}
	}
	sort.Ints(retVal)
	return retVal
}

func intArray(inVal []int, alen int) []int {
	var retVal []int
	for c := 0; c < alen; c++ {
		retVal = append(retVal, 0)
	}
	return retVal
}

func calculateFuel(inVal []int, targetPos int) int {
	sum := 0
	for _, pos := range inVal {
		fuel := 0
		distance := int(math.Abs(float64(targetPos) - float64(pos)))
		for d := 1; d <= distance; d++ {
			fuel += d
		}
		sum += fuel
	}
	return sum
}

func moveCrabs(inVal []int) int {

	target := 0
	prevVal := 0

	for {
		fuel := calculateFuel(inVal, target)
		if fuel > prevVal && prevVal != 0 {
			break
		}
		target += 1
		prevVal = fuel
	}

	return prevVal
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
