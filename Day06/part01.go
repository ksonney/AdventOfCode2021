package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fish struct {
	untilSpawn int
}

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

func newFish(spawnTime int) fish {
	var retVal fish
	retVal.untilSpawn = spawnTime
	return retVal
}

func parseData(sourceData []string) []fish {
	var retVal []fish
	dlen := len(sourceData)
	for c := 0; c < dlen; c++ {
		fishData := strings.Split(sourceData[c], ",")
		for d := 0; d < len(fishData); d++ {
			fishAge, myBad := strconv.Atoi(fishData[d])
			checkErr(myBad)
			readFish := newFish(fishAge)
			retVal = append(retVal, readFish)
		}
	}
	return retVal
}

func spawn(curSchool []fish) []fish {
	slen := len(curSchool)
	toAdd := 0
	for st := 0; st < slen; st++ {
		if curSchool[st].untilSpawn == 0 {
			toAdd++
			curSchool[st].untilSpawn = 6
		} else {
			curSchool[st].untilSpawn--
		}
	}
	for nf := 0; nf < toAdd; nf++ {
		curSchool = append(curSchool, newFish(8))
	}
	return curSchool
}
func main() {
	spawnCycle := 80

	// fname := "sample.txt"
	fname := "input.txt"
	sourceData := readData(fname)
	school := parseData(sourceData)
	// fmt.Println(school)
	schoolSize := len(school)
	fmt.Println("Initial School Size is",schoolSize)
	for sc := 0; sc < spawnCycle; sc ++ {
		// fmt.Println("Day",sc)
		school = spawn(school)
	}
	schoolSize = len(school)
	fmt.Println("Final Schools Size is", schoolSize)
}
