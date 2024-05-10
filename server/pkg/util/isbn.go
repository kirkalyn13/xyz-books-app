package util

import (
	"regexp"
	"strconv"
)

// IsValidISBN10 to check if a string is a valid ISBN-13
func IsValidISBN13(isbn string) bool {
	isbn = regexp.MustCompile(`[-]`).ReplaceAllString(isbn, "")

	if len(isbn) != 13 {
		return false
	}

	digits := make([]int, 13)
	for i := 0; i < 13; i++ {
		digit, err := strconv.Atoi(string(isbn[i]))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	checksum := 0
	for i := 0; i < 12; i++ {
		if i%2 == 0 {
			checksum += digits[i]
		} else {
			checksum += 3 * digits[i]
		}
	}

	checkDigit := (10 - (checksum % 10)) % 10

	lastDigit, err := strconv.Atoi(string(isbn[12]))
	if err != nil {
		return false
	}

	return checkDigit == lastDigit
}

// IsValidISBN10 to check if a string is a valid ISBN-10
func IsValidISBN10(isbn string) bool {
	isbn = regexp.MustCompile(`[-\s]`).ReplaceAllString(isbn, "")

	if len(isbn) != 10 {
		return false
	}

	var sum int
	for i, digit := range isbn {
		if i == 9 && digit == 'X' {
			sum += 10 * (10 - i)
		} else if digit < '0' || digit > '9' {
			return false
		} else {
			sum += int(digit-'0') * (10 - i)
		}
	}

	return sum%11 == 0
}
