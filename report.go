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

// IdentifyResult returns the product of two problematic records.
func IdentifyResult(records [][]string) int {
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records) && i != j; j++ {
			isProblem, product := IsProblematicRecordSet(records[i][0], records[j][0])
			if isProblem {
				return product
			}
		}
	}
	return -1
}

// IsProblematicRecordSet returns true if two records sum to 2020, else returns false.
func IsProblematicRecordSet(firstRecord string, secondRecord string) (bool, int) {
	first, _ := strconv.Atoi(firstRecord)
	second, _ := strconv.Atoi(secondRecord)

	if first+second == 2020 {
		fmt.Println(first, second)
		return true, first * second
	}
	return false, 0
}
