package main

import (
	"strings"

	valid "github.com/asaskevich/govalidator"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "fco"
	password = "pass"
)

func main() {

}

func normalize(number string) string {
	var normalized_number string

	noSpaceNumber := strings.ReplaceAll(number, " ", "")
	phone_number := strings.Split(noSpaceNumber, "")

	for _, digit := range phone_number {
		if valid.IsInt(digit) {
			normalized_number += digit
		}
	}

	return normalized_number
}

// re := regexp.MustCompile("\\D")
// return re.ReplaceAllString(number, "")
