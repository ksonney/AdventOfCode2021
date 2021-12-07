package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
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

func parseData(sourceData []string, retVal []int) []int {
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

func intArray(inVal []int, alen int) []int {
	var retVal []int
	for c := 0; c < alen; c++ {
		retVal = append(retVal, 0)
	}
	return retVal
}

func linearUpdate(numSteps float64) int {
	en := math.Abs(float64(numSteps))
	fmt.Println("Number of Steps",en)
	distCalc := float64(0)
	for st := float64(1); st <= en; st++ {
		distCalc = distCalc + st
		fmt.Println("Current Fuel Spent",distCalc)
	}
	fmt.Println("distance",distCalc)
	return int(distCalc)
}

func getAvg(inVal []int) float64 {
	retVal := float64(0)
	arLen := float64(len(inVal))
	fullSum := float64(0)
	
	for ga:=float64(0); ga < arLen; ga++ {
		fullSum += float64(inVal[int(ga)])
	}
	retVal = fullSum/arLen
	return retVal
}

func moveCrabs(inVal []int) int {
	var retVal int
	var startPos float64
	var endPos float64

	endPos = math.Round(getAvg(inVal))
	for c := 0; c < len(inVal); c++ {
		numSteps := float64(0)
		fmt.Println("Target Position is",endPos)
		fmt.Println("Source Position is",inVal[c])
		startPos= float64(inVal[c])
		if startPos > endPos {
			numSteps = startPos - endPos
		} else if startPos < endPos {
			numSteps = endPos - startPos
		}
		fmt.Println("Num Steps",numSteps)
		retVal += linearUpdate(numSteps)
}
	fmt.Println("Total Fuel Spent",retVal)
	return retVal
}

func main() {

	var crabsPos []int
	var minDistance int
	var fname string
	fileList := []string{"sample"} //, "input"}

	for fn := 0; fn < len(fileList); fn++ {
		fname = fileList[fn]
		sourceData := readData(fname + ".txt")
		crabsPos = parseData(sourceData, crabsPos)
		minDistance = moveCrabs(crabsPos)
		fmt.Println(crabsPos)
		fmt.Println("Minimum Move Distance for", fname, "is", minDistance)
	}

}
