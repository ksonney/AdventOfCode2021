package main

import (
	"bufio"
	"fmt"
	"os"

	// "bytes"
	"strconv"
	"strings"
)

type board struct {
	winner     bool
	fullBoard  [][]string
	markerData [][]string
}

var boardSet []board
var myBad error

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func populateBoards(boardData []string) []board {
	var retVal []board
	var lineSource board

	blen := len(boardData)
	blankLine := ".,.,.,.,."
	boardline := 0

	for c := 0; c < blen; c++ {
		if boardline == 5 {
			retVal = append(retVal, lineSource)
			boardline = 0
			lineSource = board{winner: false}
		} else {
			newLine := strings.Split(strings.ReplaceAll(boardData[c], "  ", " "), " ")
			if len(newLine) == 6 {
				newLine = newLine[1:]
			}
			lineSource.fullBoard = append(lineSource.fullBoard, newLine)
			lineSource.markerData = append(lineSource.markerData, strings.Split(blankLine, ","))
			boardline++
		}
	}
	return retVal
}

func markWinner(inVal board) board {
	var x, y int

	retVal := inVal
	//check rows
	winningRow := false
	winningCol := false
	for x = 0; x < 5; x++ {
		countX := 0
		for y = 0; y < 5; y++ {
			if inVal.markerData[x][y] == "X" {
				countX++
			}
		}
		if countX == 5 {
			winningRow = true
		}
	}
	for y = 0; y < 5; y++ {
		countY := 0
		for x = 0; x < 5; x++ {
			if inVal.markerData[x][y] == "X" {
				countY++
			}
		}
		if countY == 5 {
			winningCol = true
		}
	}
	if winningRow || winningCol {
		retVal.winner = true
	}
	return retVal
}

func markBoard(inVal board, calledNum string) board {
	var x, y int

	retVal := inVal
	for x = 0; x < 5; x++ {
		for y = 0; y < 5; y++ {
			if inVal.fullBoard[x][y] == calledNum {
				inVal.markerData[x][y] = "X"
			}
		}
	}
	return retVal
}

func tallyWinner(inVal board) int {
	var x, y, retVal int
	retVal = 0
	for x = 0; x < 5; x++ {
		for y = 0; y < 5; y++ {
			curVal, myBad := strconv.Atoi(inVal.fullBoard[x][y])
			checkErr(myBad)
			marker := inVal.markerData[x][y]
			if marker == "." { 
				retVal = retVal + curVal
			}
		}
	}
	return retVal
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

func main() {
	fname := "input.txt"
	sourceData := readData(fname)
	drawData := strings.Split(sourceData[0], ",")
	boardSet := populateBoards(sourceData[2:])
	totalBoards := len(boardSet)
	foundWinner := false
	winningBoard := 0
	winningCall := 0
	for c := 0; c < len(drawData); c++ {
		if !foundWinner {
			for d := 0; d < totalBoards; d++ {
				markBoard(boardSet[d], drawData[c])
				boardSet[d] = markWinner(boardSet[d])
				if boardSet[d].winner {
					fmt.Println("Board", d, "is a winner")
					winningBoard = d
					winningCall, myBad = strconv.Atoi(drawData[c])
					checkErr(myBad)
					foundWinner = true
				}
			}
		}
	}
	finalTally := tallyWinner(boardSet[winningBoard])
	finalScore := finalTally * winningCall
	fmt.Println("Total Boards", totalBoards)
	fmt.Println("Winning Board", winningBoard)
	fmt.Println("Winning Number", winningCall)
	fmt.Println("Winning Tally", finalTally)
	fmt.Println("Final Score", finalScore)
}
