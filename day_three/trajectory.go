// Main implements part 1 of day_3 puzzle.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Slope represents the pattern of movement that can be made.
type Slope struct {
	movesRight int
	movesDown  int
}

// Coordinate represents a location on the map.
type Coordinate struct {
	row    int
	column int
}

func main() {
	mapInput := ReadRecords("terrain.csv")
	slope := Slope{3, 1}
	fmt.Println(CountTrees(mapInput, slope, "#"))
}

// ReadRecords returns records from the provided csv filename.
func ReadRecords(filename string) [][]string {
	if len(filename) == 0 {
		log.Fatal("Invalid arguments provided")
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	return records
}

// CountTrees returns the total count of trees encountered given a mapInput.
func CountTrees(mapInput [][]string, slope Slope, treeSymbol string) int {
	if mapInput == nil || len(mapInput) == 0 || len(mapInput[0][0]) == 0 || (Slope{}) == slope || len(treeSymbol) == 0 {
		log.Fatal("Invalid arguments provided")
	}

	treeCount := 0
	currentCoordinate := Coordinate{0, 0}
	for startRow, stopRow := 0, 1; startRow < len(mapInput)-1; startRow, stopRow = startRow+1, stopRow+1 {
		currentCoordinate = Move(mapInput, slope, currentCoordinate, startRow)
		if IsTree(mapInput, currentCoordinate, treeSymbol) {
			treeCount++
		}
	}
	return treeCount
}

// Move returns an updated coordinate after a move down the slope has been made.
func Move(mapInput [][]string, slope Slope, currentCoordinate Coordinate, startRow int) Coordinate {
	numberOfColumns := len(mapInput[0][0])
	if currentCoordinate.column+slope.movesRight >= numberOfColumns {
		currentCoordinate.column = (currentCoordinate.column + slope.movesRight) % (numberOfColumns)
	} else {
		currentCoordinate.column = currentCoordinate.column + slope.movesRight
	}
	currentCoordinate.row = startRow + 1
	return currentCoordinate
}

// IsTree returns true if a cell contains the treeSybmol, else false.
func IsTree(mapInput [][]string, coordinate Coordinate, treeSymbol string) bool {
	if len(treeSymbol) == 0 {
		log.Fatal("Invalid arguments provided")
	}

	if string(mapInput[coordinate.row][0][coordinate.column]) == treeSymbol {
		return true
	}
	return false
}
