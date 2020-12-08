package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// BoardingPass represents the element of an airline boarding pass.
type BoardingPass struct {
	row    int
	column int
	seatID int
}

func main() {
	lines := ReadRecords("input.txt")
	boardingPasses := GenerateBoardingPasses(lines)
	highestSeatID := FindHighestSeatID(boardingPasses)
	fmt.Println("Highest seatID:", highestSeatID)

	boardingPasses = SortBoardingPassesByRow(boardingPasses)
	boardingPasses = SortBoardingPassesBySeatID(boardingPasses)
	mySeatID := FindMySeat(boardingPasses, highestSeatID)
	fmt.Println("My seatID:", mySeatID)
}

// ReadRecords returns records from the provided text file.
func ReadRecords(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// FindRow returns the row assignment for the input line using the first 7
// characters of the input string.
func FindRow(line string) int {
	// first 7 characters used to binary search for row on plane
	rowString := line[:7]
	low := 0
	high := 127
	mid := 0
	idx := 0
	for low <= high && idx < len(rowString) {
		mid = low + (high-low)/2
		// F means to take the lower half
		// B means to take the upper half
		currentCharacter := string(rowString[idx])
		if currentCharacter == "F" {
			high = mid - 1
		}
		if currentCharacter == "B" {
			low = mid + 1
		}
		idx++
	}
	return low
}

// FindColumn returns the column assignment for the input line using the last
// 3 characters of the provided input string.
func FindColumn(line string) int {
	// last 3 characters used to binary search for column on plane
	colString := line[7:]
	low := 0
	high := 7
	mid := 0
	idx := 0
	for low <= high && idx < len(colString) {
		mid = low + (high-low)/2
		// L means to keep the lower half
		// R means to keep the upper half
		currentCharacter := string(colString[idx])
		if currentCharacter == "R" {
			low = mid + 1
		}
		if currentCharacter == "L" {
			high = mid - 1
		}
		idx++
	}
	return low
}

// FindSeatID calculates the seatID for a BoardingPass.
func FindSeatID(row int, column int) int {
	return (row * 8) + column
}

// CreateBoardingPass generates a BoardingPass from string input.
func CreateBoardingPass(line string) BoardingPass {
	row := FindRow(line)
	column := FindColumn(line)
	seatID := FindSeatID(row, column)
	boardingPass := BoardingPass{
		row:    row,
		column: column,
		seatID: seatID}
	return boardingPass
}

// GenerateBoardingPasses returns BoardingPass elements from strings.
func GenerateBoardingPasses(lines []string) []BoardingPass {
	var boardingPasses []BoardingPass
	for _, line := range lines {
		boardingPass := CreateBoardingPass(line)
		boardingPasses = append(boardingPasses, boardingPass)
	}
	return boardingPasses
}

// FindHighestSeatID returns the highest BoardingPass seatID value given a
// collection of BoardingPass elements.
func FindHighestSeatID(boardingPasses []BoardingPass) int {
	highestSeatID := 0
	for _, boardingPass := range boardingPasses {
		highestSeatID = Max(highestSeatID, boardingPass.seatID)
	}
	return highestSeatID
}

// Max returns the higher of two integers.
func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// SortBoardingPassesByRow modifies the boardingPasses and returns them sorted by row ascending.
func SortBoardingPassesByRow(boardingPasses []BoardingPass) []BoardingPass {
	sort.Slice(boardingPasses, func(i, j int) bool {
		return boardingPasses[i].row < boardingPasses[j].row
	})
	return boardingPasses
}

// SortBoardingPassesBySeatID modifies the boardingPasses and returns them sorted by seatID ascending.
func SortBoardingPassesBySeatID(boardingPasses []BoardingPass) []BoardingPass {
	sort.Slice(boardingPasses, func(i, j int) bool {
		return boardingPasses[i].seatID < boardingPasses[j].seatID
	})
	return boardingPasses
}

// FindMySeat returns the seatID corresponding to the missing boarding pass.
func FindMySeat(sortedBoardingPasses []BoardingPass, highestSeatID int) int {
	low := 1
	high := highestSeatID
	mid := low + (high-low)/2

	// sliding window of size 2 down
	windowStart := mid - 2
	for windowEnd := windowStart + 1; windowEnd >= 2; windowEnd-- {
		if sortedBoardingPasses[windowStart].seatID+1 != sortedBoardingPasses[windowEnd].seatID {
			return sortedBoardingPasses[windowStart].seatID + 1
		}
		windowStart--
	}

	//sliding window of size 2 up
	windowStart = mid + 1
	for windowEnd := windowStart + 1; windowEnd < len(sortedBoardingPasses); windowEnd++ {
		if sortedBoardingPasses[windowStart].seatID+1 != sortedBoardingPasses[windowEnd].seatID {
			return sortedBoardingPasses[windowStart].seatID + 1
		}
		windowStart++
	}

	return 1

}
