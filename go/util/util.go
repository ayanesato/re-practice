package util

func isLeapYear(year int) bool {
	//return year%4 == 0 && (year%100 != 0 || year%400 == 0)
	if year%4 == 0 {
		if year%400 == 0 {
			return true
		}
		if year%100 != 0 {
			return true
		}
	}
	return false
}
