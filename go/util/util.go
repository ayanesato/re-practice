package main

import "fmt"

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	year := isLeapYear(2024)
	fmt.Println(year)
}
