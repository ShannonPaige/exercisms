package leap

const TestVersion = 1

// Take a year and report if it is a leap year.
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 == 0 && year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}
