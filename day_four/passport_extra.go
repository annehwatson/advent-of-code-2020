package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
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
  validFieldCount int
}

// TODO: off by 1 right now. output = 141 but it should be 140. investigate.
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
  if passport.birthYear != "" && IsValidBirthYear(passport.birthYear) {
    passport.validFieldCount++
  }
	if passport.eyeColor != "" && IsValidEyeColor(passport.eyeColor) {
    passport.validFieldCount++
  }
	if passport.expirationYear != "" && IsValidExpirationYear(passport.expirationYear) {
    passport.validFieldCount++
  }
	if passport.hairColor != "" && IsValidHairColor(passport.hairColor) {
    passport.validFieldCount++
  }
	if passport.height != "" && IsValidHeight(passport.height) {
    passport.validFieldCount++
  }
	if passport.issueYear != "" && IsValidIssueYear(passport.issueYear) {
    passport.validFieldCount++
  }
	if passport.passportID != "" && IsValidPassportID(passport.passportID) {
		passport.validFieldCount++
	}
  if passport.validFieldCount >= 7 {
    fmt.Printf("%+v\n",passport)
  }
	return passport.validFieldCount >= 7
}

// IsValidBirthYear validates Passport birth year values.
func IsValidBirthYear(birthYear string) bool {
	numYear, err := strconv.Atoi(birthYear)
	if err != nil {
		log.Fatal(err)
	}
	if numYear >= 1920 && numYear <= 2002 {
		return true
	}
	return false
}

// IsValidEyeColor validates Passport eye color values.
func IsValidEyeColor(eyeColor string) bool {
  if len(eyeColor) != 3 {
    return false
  }
	switch eyeColor {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		// no op
	}
	return false
}

// IsValidExpirationYear validates Passport expiration year values.
func IsValidExpirationYear(expirationYear string) bool {
  isMatch, err := regexp.MatchString("(2[0-9]{3})", expirationYear)
	if err != nil {
		log.Fatal(err)
	}
  if isMatch {
    numYear, err := strconv.Atoi(expirationYear)
  	if err != nil {
  		log.Fatal(err)
  	}
    if numYear >= 2020 && numYear <= 2030 {
  		return true
  	}
  }
	return false
}

// IsValidHairColor validates Passport hair color values.
func IsValidHairColor(hairColor string) bool {
	isMatch, err := regexp.MatchString("#([0-9a-f]{6})", hairColor)
	if err != nil {
		log.Fatal(err)
	}
	return isMatch
}

// IsValidHeight validates Passport height values.
func IsValidHeight(height string) bool {
  isMatch, err := regexp.MatchString("([0-9]{2,3}(cm|in))", height)
	if err != nil {
		log.Fatal(err)
	}
  if !isMatch {
    return false
  }

	heightNum, err := strconv.Atoi(string(height[:len(height)-2]))
	if err != nil {
		log.Fatal(err)
	}

  measurementSystem := string(height[len(height)-2:])
  if measurementSystem == "cm" {
    return heightNum >= 150 && heightNum <= 193
  }
  if measurementSystem == "in" {
    return heightNum >= 59 && heightNum <= 76
  }
  return false
}

// IsValidIssueYear validates Passport issue year values.
func IsValidIssueYear(issueYear string) bool {
	numYear, err := strconv.Atoi(issueYear)
	if err != nil {
		log.Fatal(err)
	}
	if numYear >= 2010 && numYear <= 2020 {
		return true
	}
	return false
}

// IsValidPassportID validates Passport id values.
func IsValidPassportID(passportID string) bool {
	isMatch, err := regexp.MatchString("([0-9]{9})", passportID)
	if err != nil {
		log.Fatal(err)
	}
	return isMatch
}
