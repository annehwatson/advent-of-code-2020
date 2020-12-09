package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// CustomsResult is a representation of group Customs Form responses.
type CustomsResult struct {
	uniqueYesValues map[string]int
	respondents     int
}

func main() {
	lines := ReadRecords("input.txt")
	customsResults := RecordGroupYesResults(lines)
	fmt.Println("Total yes votes by anyone in a group:",TallyYesResults(customsResults))
  fmt.Println("Total unanimous yes CustomsResult answers:",TallyUnanimousResults(customsResults))
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

// RecordGroupYesResults creates CustomsResult records and returns them.
func RecordGroupYesResults(lines []string) []CustomsResult {
	var customsResults []CustomsResult
	currentCustomsResult := CustomsResult{}
	for i := 0; i < len(lines); i++ {
		line := string(lines[i])
		if line == "" {
			customsResults = append(customsResults, currentCustomsResult)
			currentCustomsResult = CustomsResult{}
		} else {
      currentCustomsResult.respondents++
			currentCustomsResult = ParseCustomsResults(currentCustomsResult, line)
		}
	}
	return customsResults
}

// ParseCustomsResults modifies the CustomsResult uniqueYesValues map to signify
// Yes answers on the Customs form.
func ParseCustomsResults(customsResult CustomsResult, resultLine string) CustomsResult {
	if customsResult.uniqueYesValues == nil {
		customsResult.uniqueYesValues = make(map[string]int)
	}

	for _, letter := range resultLine {
		customsResult.uniqueYesValues[string(letter)]++
	}
	return customsResult
}

// TallyYesResults sums up the total unique yes count for all CustomsResult records.
func TallyYesResults(customsResults []CustomsResult) int {
	totalYesCount := 0
	for _, customsResult := range customsResults {
		totalYesCount = totalYesCount + len(customsResult.uniqueYesValues)
	}
	return totalYesCount
}

// TallyUnanimousResults sums up the total unanimous yes counts seen in all
// CustomsResult records.
func TallyUnanimousResults(customsResults []CustomsResult) int {
  totalUnanimousYesCount := 0
  for _, customsResult := range customsResults {
    for _, mapping := range customsResult.uniqueYesValues {
      if mapping == customsResult.respondents {
        totalUnanimousYesCount++
      }
    }
  }
  return totalUnanimousYesCount
}
