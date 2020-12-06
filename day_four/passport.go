package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Passport is a representation of the elements of a passport.
type Passport struct {
	birthYear      string
	countryID      string
	expirationYear string
	eyeColor       string
	hairColor      string
	height         string
	issueYear      string
	passportID     string
}

func main() {
	lines := ReadLinesFromFile("input.txt")
	passports := CreatePasswords(lines)
	fmt.Println(CountValidPassports(passports))
}

// ReadLinesFromFile reads a file and returns the output.
func ReadLinesFromFile(filename string) []string {
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

// CreatePasswords creates Passport records and returns them.
func CreatePasswords(lines []string) []Passport {
	var passports []Passport
	currentPassport := Passport{}
	for i := 0; i < len(lines); i++ {
		line := string(lines[i])
		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = Passport{}
		} else {
			currentPassport = ParsePassportDetails(currentPassport, line)
		}
	}
	return passports
}

// ParsePassportDetails assigns details to a Passport record from string input.
func ParsePassportDetails(passport Passport, detailLine string) Passport {
	properties := strings.Split(strings.Split(detailLine, ",")[0], " ")
	for i := 0; i < len(properties); i++ {
		property := strings.Split(properties[i], ":")
		key := property[0]
		value := property[1]
		switch key {
		case "byr":
			passport.birthYear = value
		case "cid":
			passport.countryID = value
		case "ecl":
			passport.eyeColor = value
		case "eyr":
			passport.expirationYear = value
		case "hcl":
			passport.hairColor = value
		case "hgt":
			passport.height = value
		case "iyr":
			passport.issueYear = value
		case "pid":
			passport.passportID = value
		default:
			// no op
		}
	}
	return passport
}

// CountValidPassports returns the number of valid passports.
func CountValidPassports(passports []Passport) int {
	validPassportCount := 0
	for _, passport := range passports {
		if IsValidPassport(passport) {
			validPassportCount++
		}
	}
	return validPassportCount
}

// IsValidPassport checks if Passport record is valid.
func IsValidPassport(passport Passport) bool {
	if passport.birthYear == "" ||
		passport.eyeColor == "" ||
		passport.expirationYear == "" ||
		passport.hairColor == "" ||
		passport.height == "" ||
		passport.issueYear == "" ||
		passport.passportID == "" {
		return false
	}
	return true
}

func IsValidBirthYear(birthYear string) bool {
  return true
}
