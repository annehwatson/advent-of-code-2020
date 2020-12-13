package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := ReadRecords("input.txt")
	fmt.Println(RunInstructions(instructions))
}

// ReadRecords takes a filename input and returns a new-line delimited collection
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

// RunInstructions executes the instructions until either we are going to execute
// an already executed instruction or we have reached the end of our instructions.
func RunInstructions(instructions []string) int {
	visited := make(map[int]int)
	total := 0

	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		// check if we have executed this instruction already
		if visited[i] != 0 || instruction == "" {
			return total
		}
		// mark the instruction index as visited
		visited[i] = 1
		// execute the instruction
		command := instruction[:3]
		sign := string(instruction[4])
		value, err := strconv.Atoi(instruction[5:])
		if err != nil {
			log.Fatal(err)
		}

		switch command {
		case "nop":
			continue
		case "acc":
			if sign == "+" {
				total += value
			} else {
				total -= value
			}
		case "jmp":
			if sign == "+" {
				i += value - 1
			} else {
				i -= value + 1
			}
		}
	}
	fmt.Println(visited)
	return total
}
