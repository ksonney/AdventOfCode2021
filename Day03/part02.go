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
var o2genCount, co2scrubCount [16]int

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func o2Rating(readData [][]string) int64 {

	return 0
}

func co2Rating(readData [][]string) int64 {
	
	return 0
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
	// <LaurenCooper>Ist this bovvered face thou
	// seest before thee?</LaurenCooper>
	rlen := 12

}