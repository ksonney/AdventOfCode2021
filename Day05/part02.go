package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	X int
	Y int
}
type vents struct {
	startPoints []coord
	endPoints   []coord
	maxX        int
	maxY        int
	fullMap     [][]int
}

var myBad error
var myMap vents

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

func printMap(mapData vents) {
	for m := 0; m < mapData.maxX+1; m++ {
		fmt.Println(mapData.fullMap[m])
	}

}

func parseData(sourceData []string) vents {
	var retVal vents
	var newCoord coord
	dlen := len(sourceData)
	for c := 0; c < dlen; c++ {
		lineData := strings.Split(sourceData[c], " -> ")
		for d := 0; d < 2; d++ {
			bovvered := strings.Split(lineData[d], ",")
			newX, myBad := strconv.Atoi(bovvered[0])
			checkErr(myBad)
			newY, myBad := strconv.Atoi(bovvered[1])
			checkErr(myBad)
			newCoord.X = newX
			newCoord.Y = newY
			if newX > retVal.maxX {
				retVal.maxX = newX
			}
			if newY > retVal.maxY {
				retVal.maxY = newY
			}
			switch d {
			case 0:
				retVal.startPoints = append(retVal.startPoints, newCoord)
			case 1:
				retVal.endPoints = append(retVal.endPoints, newCoord)
			}
		}
	}
	return retVal
}

func getOverLap(mapData vents) int {
	retVal := 0
	for i := 0; i < len(mapData.fullMap); i++ {
		for j := 0; j < len(mapData.fullMap[i]); j++ {
			if mapData.fullMap[i][j] >= 2 {
				retVal++
			}
		}
	}
	return retVal
}

func mapHorizVents(mapData vents) vents {
	retVal := mapData

	for i := 0; i < len(retVal.endPoints); i++ {
		startX := retVal.startPoints[i].X
		startY := retVal.startPoints[i].Y
		endX := retVal.endPoints[i].X
		endY := retVal.endPoints[i].Y
		if startX == endX {
			// verticle Line
			if startY > endY {
				startY = retVal.endPoints[i].Y
				endY = retVal.startPoints[i].Y
			}
			for c := startY; c <= endY; c++ {
				retVal.fullMap[c][startX]++
			}
		}
	}
	return retVal
}

func mapVerticVents(mapData vents) vents {
	retVal := mapData

	for i := 0; i < len(retVal.endPoints); i++ {
		startX := retVal.startPoints[i].X
		startY := retVal.startPoints[i].Y
		endX := retVal.endPoints[i].X
		endY := retVal.endPoints[i].Y
		if startY == endY {
			// horizontal line
			if startX > endX {
				startX = retVal.endPoints[i].X
				endX = retVal.startPoints[i].X
			}
			for c := startX; c <= endX; c++ {
				retVal.fullMap[startY][c]++
			}
		}
	}
	return retVal
}

func mapDiagVents(mapData vents) vents {
	retVal := mapData

	for i := 0; i < len(retVal.endPoints); i++ {
		startX := retVal.startPoints[i].X
		startY := retVal.startPoints[i].Y
		endX := retVal.endPoints[i].X
		endY := retVal.endPoints[i].Y
		if startX != endX && startY != endY {
			d := startY
			c := startX
			for c != endX && d != endY {
				retVal.fullMap[d][c]++
				if c < endX {
					c++
				} else {
					c--
				}
				if d < endY {
					d++
				} else {
					d--
				}
			}
			retVal.fullMap[d][c]++
		}
	}
	return retVal
}

func blankLine(size int) []int {
	var retVal []int
	for x := 0; x <= size; x++ {
		retVal = append(retVal, 0)
	}
	return retVal
}

func blankMap(mapData vents) vents {
	retVal := mapData
	for y := 0; y <= mapData.maxY; y++ {
		freshLine := blankLine(mapData.maxX)
		retVal.fullMap = append(retVal.fullMap, freshLine)
	}

	return retVal
}

func main() {
	// fname := "sample.txt"
	fname := "input.txt"
	sourceData := readData(fname)
	myMap = parseData(sourceData)
	myMap = blankMap(myMap)
	myMap = mapHorizVents(myMap)
	myMap = mapVerticVents(myMap)
	myMap = mapDiagVents(myMap)
	// printMap(myMap)
	overlapping := getOverLap(myMap)
	fmt.Println("Number of Overlapping Vents :", overlapping)
}
