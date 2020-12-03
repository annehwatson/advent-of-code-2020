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
	password       string
	letter         string
	firstPosition  int
	secondPosition int
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
	positions := strings.Split(recordFields[0], "-")
	first, err := strconv.Atoi(positions[0])
	if err != nil {
		log.Fatal(err)
	}
	second, err := strconv.Atoi(positions[1])
	if err != nil {
		log.Fatal(err)
	}

	passwordRecord := PasswordRecord{
		password:       recordFields[2],
		letter:         string(recordFields[1][0]),
		firstPosition:  first,
		secondPosition: second}
	return passwordRecord
}

// IsValidPassword returns true if the given password contains the required frequency of the letter specified in the corporate password policy.
// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// non zero indexed
func IsValidPassword(password string, firstPosition int, secondPosition int, letter string) bool {
	if len(password) == 0 || firstPosition <= 0 || secondPosition <= 0 {
		log.Fatal("Invalid arguments provided")
	}

	firstPositionLetter := string(password[firstPosition-1])
	secondPositionLetter := string(password[secondPosition-1])

	if firstPositionLetter == secondPositionLetter {
		return false
	}
	if firstPositionLetter == letter || secondPositionLetter == letter {
		return true
	}
	return false
}

// CountValidPasswords returns the total count of valid passwords.
func CountValidPasswords(passwordRecords []PasswordRecord) int {
	validCount := 0
	for i := 0; i < len(passwordRecords); i++ {
		if IsValidPassword(passwordRecords[i].password, passwordRecords[i].firstPosition, passwordRecords[i].secondPosition, passwordRecords[i].letter) {
			validCount++
		}
	}
	return validCount
}
