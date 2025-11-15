// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap provides a method to check whether the given year is leap or not.
package leap

// IsLeapYear returns boolean value for whether a year is leap or not.
func IsLeapYear(year int) bool {
	return (year % 100 == 0 && year % 400 == 0) || (year % 4 == 0 && year % 100 > 0)
}
