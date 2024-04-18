package util

func isLeapYear(year int) bool {
	//return year%4 == 0 && (year%100 != 0 || year%400 == 0)
	var isLeap bool
	if year%4 == 0 {
		if year%400 == 0 {
			isLeap = true
		} else if year%100 != 0 {
			isLeap = true
		}
	}
	return isLeap
}
