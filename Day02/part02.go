package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var posX, posZ, aim int = 0,0,0
	inputFile, myBad := os.Open("input.txt")
	checkErr(myBad)
	defer inputFile.Close()

	inFile := bufio.NewScanner(inputFile)
	for inFile.Scan() {
		direction := strings.Split(inFile.Text()," ")
		dist,myBad := strconv.ParseInt(direction[1],10,0)
		checkErr(myBad)
		switch dir := direction[0]; dir {
		case "forward":
			posX = posX + int(dist)
			posZ = posZ + (int(dist) * aim)
		case "down":
			aim = aim + int(dist)
		case "up":
			aim = aim - int(dist)
		}
	}
	
	fmt.Println("Depth is",posZ)
	fmt.Println("Distance is", posX)
	fmt.Println("Final Position is", (posX * posZ))
	myBad = inFile.Err();
	checkErr(myBad)

}
