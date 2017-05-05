package hamming

import (
	"errors"
	"strings"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Cannot compare these two strings")
	}

	distance := 0

	s1 := strings.Split(a, "")
	s2 := strings.Split(b, "")

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance, nil
}
