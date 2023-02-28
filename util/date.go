package util

import "time"

func SameDay(t1, t2 time.Time) bool {
	if t1.Year() == t2.Year() {
		if t1.Month() == t2.Month() {
			if t1.Day() == t2.Day() {
				return true
			}
		}
	}
	return false
}
