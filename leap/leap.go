package leap

// IsLeapYear: Given a year, reports if it is a leap year.
func IsLeapYear(year int) bool {

	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
