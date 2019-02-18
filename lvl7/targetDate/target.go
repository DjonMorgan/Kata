// You have an amount of money a0 > 0 and you deposit it with a constant interest rate of p% > 0 per year on the 1st of January 2016. You want to have an amount a >= a0.
//
// Task:
// The function date_nb_days (or dateNbDays...) with parameters a0, a, p will return, as a string, the date on which you have just reached a.
//
// Example:
// date_nb_days(100, 101, 0.98) --> "2017-01-01" (366 days)
//
// date_nb_days(100, 150, 2.00) --> "2035-12-26" (7299 days)
//
// Notes:
// The return format of the date is "YYYY-MM-DD"
// The deposit is always on the "2016-01-01"
// If p% is the rate for a year the rate for a day is p divided by 36000 since banks consider that there are 360 days in a year.
// You have: a0 > 0, p% > 0, a >= a0
package main

import (
	"fmt"
	"time"
)

func main() {
	var a0 int = 1000
	var a int = 2000
	var p int = 5.00
	fmt.Println(start)
}

func DateNbDays(a0 int, a int, p int) string {
	var x string
	payday := a0 * float64(p) / 36000
	var y float64 = a0 + payday/a
	dateStart := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	x = start.Add(time.Hour * 24 * 10 * y)
	start := x.String()
}
