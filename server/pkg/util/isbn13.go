package util

import (
	"regexp"
	"strconv"
)

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
