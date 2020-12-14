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
	adapters := ReadInput("input.txt")
	sort.Ints(adapters)
	joltageDifferenceFrequencies := RecordJoltageDifferences(adapters)
	fmt.Println("Part 1 result:", joltageDifferenceFrequencies[1]*joltageDifferenceFrequencies[3])

	adapters = append(adapters, adapters[len(adapters)-1]+3)
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	fmt.Println("Part 2 result:", FindDistinctArrangements(adapters))
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

// RecordJoltageDifferences returns the frequencies of 1-jolt and 3-jolt
// differences among a collection of adapters.
func RecordJoltageDifferences(adapters []int) map[int]int {
	joltageDifferenceFrequencies := make(map[int]int)
	// charging outlet
	joltageDifferenceFrequencies[adapters[0]]++

	for i, j := 0, 1; i < len(adapters)-1; i, j = i+1, j+1 {
		diff := adapters[j] - adapters[i]
		joltageDifferenceFrequencies[diff]++
	}
	// built-in adapter
	joltageDifferenceFrequencies[3]++
	return joltageDifferenceFrequencies
}

// FindDistinctArrangements returns the count of valid arrangements of adapters
// that connect the charging outlet to the device.
func FindDistinctArrangements(adapters []int) int {
	dp := make(map[int]int)
	dp[0] = 1
	diffs := []int{1, 2, 3}

	for _, adapter := range adapters {
		for _, diff := range diffs {
			connectingAdapter := adapter + diff
			if Contains(adapters, connectingAdapter) {
				dp[connectingAdapter] += dp[adapter]
			}
		}
	}
	return dp[adapters[len(adapters)-1]]
}

// Contains returns true if the number is found in the numbers collection, else
// false.
func Contains(numbers []int, number int) bool {
	for i := 0; i < len(numbers); i++ {
		if number == numbers[i] {
			return true
		}
	}
	return false
}
