// Package main analyzes expense reports for validity.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	records := ReadRecords("elves_expense.csv")
	fmt.Println(IdentifyResult(records))
}

// ReadRecords returns the expense report entries.
func ReadRecords(expenseReportFileName string) [][]string {
	f, err := os.Open(expenseReportFileName)
	if err != nil {
		log.Fatal(err)
	}

	records, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

// IdentifyResult returns the product of three problematic records in problematic time complexity.
func IdentifyResult(records [][]string) int {
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records) && i != j; j++ {
			for k := 0; k < len(records) && i != k && k != j; k++ {
				isProblem, product := IsProblematicRecordSet(records[i][0], records[j][0], records[k][0])
				if isProblem {
					return product
				}
			}
		}
	}
	return -1
}

// IsProblematicRecordSet returns true if three records sum to 2020, else returns false.
func IsProblematicRecordSet(firstRecord string, secondRecord string, thirdRecord string) (bool, int) {
	first, _ := strconv.Atoi(firstRecord)
	second, _ := strconv.Atoi(secondRecord)
	third, _ := strconv.Atoi(thirdRecord)

	if first+second+third == 2020 {
		fmt.Println(first, second, third)
		return true, first * second * third
	}
	return false, 0
}
