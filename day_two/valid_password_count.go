package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// PasswordRecord is a representation of the elements of a password file row.
type PasswordRecord struct {
	password    string
	letter      string
	requiredMin int
	requiredMax int
}

func main() {

	records := ReadRecords("passwords_policies.csv")
	passwordRecords := CreatePasswordRecords(records)
	validPasswordCount := CountValidPasswords(passwordRecords)
	fmt.Println(validPasswordCount)
}

// ReadRecords returns records from the provided csv filename.
func ReadRecords(filename string) [][]string {
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

// CreatePasswordRecords converts the provided records into PasswordRecord elements.
func CreatePasswordRecords(records [][]string) []PasswordRecord {
	passwordRecords := make([]PasswordRecord, len(records))

	for i := 0; i < len(records); i++ {
		passwordRecords[i] = ConvertToPasswordRecord(records[i][0])
	}

	return passwordRecords
}

// ConvertToPasswordRecord converts a record from csv to a PasswordRecord.
// 2-7 p: pbhhzpmppb
func ConvertToPasswordRecord(record string) PasswordRecord {
	recordFields := strings.Fields(record)
	policyRange := strings.Split(recordFields[0], "-")
	min, err := strconv.Atoi(policyRange[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(policyRange[1])
	if err != nil {
		log.Fatal(err)
	}

	passwordRecord := PasswordRecord{
		password:    recordFields[2],
		letter:      string(recordFields[1][0]),
		requiredMin: min,
		requiredMax: max}
	return passwordRecord
}

// CountOccurrences returns an int representing the frequency of the letter in the password.
func CountOccurrences(password string, letter string) int {
	occurrences := 0
	for i := 0; i < len(password); i++ {
		if string(password[i]) == letter {
			occurrences++
		}
	}
	return occurrences
}

// IsValidPassword returns true if the given password contains the required frequency of the letter specified in the corporate password policy.
func IsValidPassword(password string, minCount int, maxCount int, letter string) bool {
	if len(password) == 0 || minCount < 0 || maxCount < 0 {
		log.Fatal("Invalid arguments provided")
	}

	letterFrequency := CountOccurrences(password, letter)
	if letterFrequency >= minCount && letterFrequency <= maxCount {
		return true
	}
	return false
}

// CountValidPasswords returns the total count of valid passwords.
func CountValidPasswords(passwordRecords []PasswordRecord) int {
	validCount := 0
	for i := 0; i < len(passwordRecords); i++ {
		if IsValidPassword(passwordRecords[i].password, passwordRecords[i].requiredMin, passwordRecords[i].requiredMax, passwordRecords[i].letter) {
			validCount++
		}
	}
	return validCount
}
