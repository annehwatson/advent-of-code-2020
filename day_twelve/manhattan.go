package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Direction represents a cardinal direction
type Direction int

const (
	// North represents the cardinal direction north.
	North Direction = iota
	// East represents the cardinal direction east.
	East
	// South represents the cardinal direction south.
	South
	// West represents the cardinal direction west.
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// Instruction represents an Action Value pair.
type Instruction struct {
	Action string
	Value  int
}

func main() {
	instructions := ReadRecords("input.txt")
	directionCountMap := FollowInstructions(instructions)
	fmt.Println("Manhattan distance part 1:", CalculateManhattanDistance(directionCountMap))
}

// ReadRecords takes a filename input and returns the new-line delimited collection
// of strings from the file.
func ReadRecords(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(rawBytes), "\n")
}

// FollowInstructions populates a direction frequency map.
func FollowInstructions(instructions []string) map[Direction]int {
	directionCountMap := make(map[Direction]int)
	directionFacing := East
	for _, stringInstruction := range instructions {
		value, err := strconv.Atoi(stringInstruction[1:])
		if err != nil {
			log.Fatal(err)
		}
		currentInstruction := Instruction{string(stringInstruction[0]), value}
		switch currentInstruction.Action {
		case "N":
			directionCountMap[North] += currentInstruction.Value
		case "S":
			directionCountMap[South] += currentInstruction.Value
		case "E":
			directionCountMap[East] += currentInstruction.Value
		case "W":
			directionCountMap[West] += currentInstruction.Value
		case "F":
			directionCountMap[directionFacing] += currentInstruction.Value
		case "L":
			ninetyDegreeTurnsCount := currentInstruction.Value / 90
			currentDirectionIndex := (ConvertDirectionToIndex(directionFacing) - ninetyDegreeTurnsCount)
			if currentDirectionIndex < 0 {
				currentDirectionIndex += 4
			}
			directionFacing = ConvertIndexToDirection(currentDirectionIndex)
		case "R":
			ninetyDegreeTurnsCount := currentInstruction.Value / 90
			currentDirectionIndex := (ConvertDirectionToIndex(directionFacing) + ninetyDegreeTurnsCount) % 4
			directionFacing = ConvertIndexToDirection(currentDirectionIndex)
		}
	}
	return directionCountMap
}

// CalculateManhattanDistance returns the sum of the absolute values of a ship's
// east/west and north/south positions.
func CalculateManhattanDistance(directionCountMap map[Direction]int) int {
	eastWest := math.Abs(float64(directionCountMap[East]) - float64(directionCountMap[West]))
	northSouth := math.Abs(float64(directionCountMap[North]) - float64(directionCountMap[South]))
	return int(eastWest + northSouth)
}

// ConvertIndexToDirection returns a Direction from the corresponding int index.
func ConvertIndexToDirection(index int) Direction {
	indexDirectionMap := map[int]Direction{
		0: North,
		1: East,
		2: South,
		3: West}
	return indexDirectionMap[index]
}

// ConvertDirectionToIndex returns an int index from the corresponding Direction.
func ConvertDirectionToIndex(direction Direction) int {
	directionIndexMap := map[Direction]int{
		North: 0,
		East:  1,
		South: 2,
		West:  3}
	return directionIndexMap[direction]
}
