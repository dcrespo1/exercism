package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var s []string

	if number%3 == 0 {
		s = append(s, "Pling")
	}

	if number%5 == 0 {
		s = append(s, "Plang")
	}

	if number%7 == 0 {
		s = append(s, "Plong")
	}

	if len(s) > 0 {
		return strings.Join(s, "")
	}
	return strconv.Itoa(number)
}
