// Let us begin with an example:
//
// A man has a rather old car being worth $2000.He saw a secondhand car being worth $8000.He wants to keep his old car until he can buy the secondhand one.
//
// He thinks he can save $1000 each month but the prices of his oldcar and of the new one decrease of 1.5 percent per month.Furthermore this percent of loss increases by0.5 percentat the end of every two months.Our man finds it difficult to make all these calculations.
//
// Can you help him?
//
// How many months will it take him to save up enough money to buy the car he wants, and how much money will he have left over?
//
// Parameters and return of function:
//
// parameter (positive int, guaranteed) startPriceOld (Old car price)
// parameter (positive int, guaranteed) startPriceNew (New car price)
// parameter (positive int, guaranteed) savingperMonth
// parameter (positive float or int, guaranteed) percentLossByMonth
//
// nbMonths(2000, 8000, 1000, 1.5) should return [6, 766] or (6, 766)
// where6 is the number of months atthe end of which he can buy the new car and766 is the nearest integer to766.158 (rounding766.158 gives766).
//
// Note:
//
// Selling, buying and saving are normally done at end of month.Calculations are processed at the end of each considered monthbut if, by chance from the start, the value of the old car is bigger than the value of the new one or equal there is no saving to be made, no need to wait so he can at the beginning of the month buy the new car:
//
// nbMonths(12000, 8000, 1000, 1.5) should return [0, 4000]
// nbMonths(8000, 8000, 1000, 1.5) should return [0, 0]
// We don't take care of a deposit of savings in a bank:-)
package main

import (
	"fmt"
	"math"
)

func main() {
	startPriceOld := 2000
	startPriceNew := 8000
	savingperMonth := 1000
	var percentLossByMonth float64 = 1.5
	NbMonths(startPriceOld, startPriceNew, savingperMonth, percentLossByMonth)
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) (int, float64) {
	x := 0
	var b float64 = 1
	z := true
	for float64(startPriceNew)-float64(startPriceNew)*percentLossByMonth/100*float64(x)-float64(startPriceNew)*float64(0.5/100*math.Gamma(b)) > float64(startPriceOld)-float64(startPriceOld)*percentLossByMonth/100*float64(x)+float64(startPriceOld)*float64(0.5/100*math.Gamma(b))+float64(savingperMonth*x) {
		x += 1
		// fmt.Println(x)
		if z == false {
			b += 1
			z = true
			continue
		}
		z = false
		fmt.Println(b)
	}
	y := float64(startPriceOld) - float64(startPriceOld)*percentLossByMonth/100*float64(x) + float64(startPriceOld)*float64(0.5/100*math.Gamma(b)) + float64(savingperMonth*x) - float64(startPriceNew) + float64(startPriceNew)*percentLossByMonth/100*float64(x) + float64(startPriceNew)*float64(0.5/100*math.Gamma(b))
	fmt.Println(x, y)
	return x, y
}
