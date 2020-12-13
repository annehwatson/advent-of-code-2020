package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := ReadInput("input.txt")
	invalidNumber := FindInvalidNumber(numbers, 25)
	fmt.Println("invalidNumber:", invalidNumber)
	contiguousRange := FindContiguousRange(numbers, invalidNumber)
	fmt.Println(FindEncryptionWeakness(contiguousRange))
}

// ReadInput takes a filename input and converts the new-line delimited collection
// of strings from the file into ints.
func ReadInput(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rawBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	strings := strings.Split(string(rawBytes), "\n")
	var numbers []int
	for i := 0; i < len(strings); i++ {
		if strings[i] != "" {
			number, err := strconv.Atoi(strings[i])
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)
		}
	}
	return numbers
}

// FindInvalidNumber returns the invalid number.
func FindInvalidNumber(numbers []int, preambleSize int) int {
	for i, j := preambleSize-1, preambleSize+1; i < len(numbers)-1; i, j = i+1, j+1 {
		if !IsValidNumber(numbers[i+1], numbers[i-preambleSize+1:i+1]) {
			return numbers[i+1]
		}
	}
	return -1
}

// IsValidNumber returns true if the input number is a sum of two numbers from the
// preamble, else false.
func IsValidNumber(number int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := 0; j < len(preamble); j++ {
			if i == j {
				continue
			} else {
				if preamble[i]+preamble[j] == number {
					return true
				}
			}
		}
	}
	return false
}

// FindContiguousRange returns the contiguous numbers that add up to the
// invalidNumber provided as input.
func FindContiguousRange(numbers []int, invalidNumber int) []int {
	windowStart := 0
	currentTotal := 0
	var result []int

	for windowEnd := 0; windowEnd < len(numbers); windowEnd++ {
		currentTotal += numbers[windowEnd]

		for currentTotal > invalidNumber {
			currentTotal -= numbers[windowStart]
			windowStart++
		}
		if currentTotal == invalidNumber {
			result = numbers[windowStart : windowEnd+1]
			return result
		}
	}
	return result
}

// FindEncryptionWeakness returns the sum of the smallest and largest numbers
// from the provided continguousRange.
func FindEncryptionWeakness(contiguousRange []int) int {
	sort.Ints(contiguousRange)
	return contiguousRange[0] + contiguousRange[len(contiguousRange)-1]
}
