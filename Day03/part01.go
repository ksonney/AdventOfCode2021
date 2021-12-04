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
var gammaRate, epsilonRate int64
var gammaCount, epsilonCount [16]int

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	inputFile, myBad := os.Open("input.txt")
	checkErr(myBad)
	defer inputFile.Close()
	inFile := bufio.NewScanner(inputFile)
	linecount := 0
	// Read the file in
	for inFile.Scan() {
		readingCount = append(readingCount, strings.Split(inFile.Text(), ""))
		linecount++
	}
	// This is terrible and I am ashamed of myself
	// Don't judge me
	// number of digits in the source data
	rlen := 12
	for c := 0; c < linecount; c++ {
		readings := readingCount[c]
		// You're judging me 
		for r := 0; r < rlen; r++ {
			switch q := readings[r]; q {
			case "1":
				gammaCount[r+4] = gammaCount[r+4] + 1
			case "0":
				epsilonCount[r+4] = epsilonCount[r+4] + 1
			}
		}
	}
	// <LaurenCooper>I'm not bovvered</LaurenCooper>
	for c := 15; c >= 0; c-- {
		if c > 3 {
			if gammaCount[c] > epsilonCount[c] {
				gammaCount[c] = 1
				epsilonCount[c] = 0
			} else {
				gammaCount[c] = 0
				epsilonCount[c] = 1
			}
		}
	}
	// <LaurenCooper>Do I even LOOK bovvered</LaurenCooper>
	var gammaBin bytes.Buffer
	var epsilonBin bytes.Buffer
	for c := 0; c < 16; c++ {
		gammaBin.WriteString(strconv.Itoa(gammaCount[c]))
		epsilonBin.WriteString(strconv.Itoa(epsilonCount[c]))
	}
	gammaRate, myBad = strconv.ParseInt(gammaBin.String(), 2, 16)
	checkErr(myBad)
	epsilonRate, myBad = strconv.ParseInt(epsilonBin.String(), 2, 16)
	checkErr(myBad)
	fmt.Println("Gamma is   :", gammaRate)
	fmt.Println("Epsilon is :", epsilonRate)
	fmt.Println("Power Consumption is :", (gammaRate * epsilonRate))
}
