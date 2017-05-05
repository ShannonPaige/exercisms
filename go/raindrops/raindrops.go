package raindrops

import "strconv"

const testVersion = 3

func Convert(x int) string {
	rain := ""

	if x%3 == 0 {
		rain += "Pling"
	}
	if x%5 == 0 {
		rain += "Plang"
	}
	if x%7 == 0 {
		rain += "Plong"
	}
	if rain == "" {
		rain = strconv.Itoa(x)
	}

	return rain
}
