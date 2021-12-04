package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var readingCount [][]string
var o2genRate, co2scrubRate int64
var zeroCount, oneCount [16]int
var o2genArray, co2scrubArray [][]string

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BinToInt(inVal []string) int64 {
	var retVal int64
	var destBin bytes.Buffer
	for c := 0; c < len(inVal); c++ {
		destBin.WriteString(inVal[c])
	}
	retVal, myBad := strconv.ParseInt(destBin.String(), 2, 16)
	checkErr(myBad)
	return retVal
}

func scrubValues(inData [][]string, position int, detectValue string) [][]string {
	var retVal [][]string
	var slen, rvlen int

	slen = len(inData)
	if position >= len(inData[0]) {
		return inData
	}
	for q := 0; q < slen; q++ {
		if inData[q][position] == detectValue {
			retVal = append(retVal, inData[q])
		}
	}
	rvlen = len(retVal)

	if rvlen < 1 {
		retVal = append(retVal, inData[slen-1])
	}
	return retVal
}

func countVals(inData [][]string, countVal string) [16]int {
	linecount := len(inData)
	rlen := 12
	var countArray [16]int
	for c := 0; c < linecount; c++ {
		readings := inData[c]
		// Amest I bovveréd?
		for r := 0; r < rlen; r++ {
			if readings[r] == countVal {
				countArray[r] = countArray[r] + 1
			}
		}
	}
	return countArray
}

func main() {
	inputFile, myBad := os.Open("input.txt")
	checkErr(myBad)
	defer inputFile.Close()
	inFile := bufio.NewScanner(inputFile)
	for inFile.Scan() {
		readingCount = append(readingCount, strings.Split(inFile.Text(), ""))
	}
	// Ist this bovvered face thou seest before thee?
	rlen := 12
	o2genArray = readingCount
	co2scrubArray = readingCount

	// My Liege, I be not bovveréd, forsooth.
	for r := 0; r < rlen; r++ {
		oneCount = countVals(o2genArray, "1")
		zeroCount = countVals(o2genArray, "0")
		if len(o2genArray) > 1 {
			if (oneCount[r] > zeroCount[r]) || (oneCount[r] == zeroCount[r]) {
				o2genArray = scrubValues(o2genArray, r, "1")
			} else {
				o2genArray = scrubValues(o2genArray, r, "0")
			}
		}
	}
	// Regardez mon visage. Suis-je bovvered?
	for r := 0; r < rlen; r++ {
		oneCount = countVals(co2scrubArray, "1")
		zeroCount = countVals(co2scrubArray, "0")
		if len(co2scrubArray) > 1 {
			if (oneCount[r] > zeroCount[r]) || (oneCount[r] == zeroCount[r]) {
				co2scrubArray = scrubValues(co2scrubArray, r, "0")
			} else {
				co2scrubArray = scrubValues(co2scrubArray, r, "1")
			}
		}
	}

	o2genRate = BinToInt(o2genArray[0])
	co2scrubRate = BinToInt(co2scrubArray[0])
	fmt.Println("O2 Generation is       :", o2genRate)
	fmt.Println("CO2 Srubbing is        :", co2scrubRate)
	fmt.Println("Life Support Rating is :", (o2genRate * co2scrubRate))
}
