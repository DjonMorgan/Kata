// System A : buy a ticket (15 dollars) every time
// System B : buy a card (500 dollars) and every time
//     buy a ticket the price of which is 0.90 times the price he paid for the previous one.
// #Example: If John goes to the cinema 3 times:
//
// System A : 15 * 3 = 45
// System B : 500 + 15 * 0.90 + (15 * 0.90) * 0.90 + (15 * 0.90 * 0.90) * 0.90 ( = 536.5849999999999, no rounding for each ticket)
// John wants to know how many times he must go to the cinema so that the final result of System B, when rounded up to the next dollar, will be cheaper than System A.
//
// The function movie has 3 parameters: card (price of the card), ticket (normal price of a ticket), perc (fraction of what he paid for the previous ticket) and returns the first n such that
//
// ceil(price of System B) < price of System A.
// More examples:
//
// movie(500, 15, 0.9) should return 43
//     (with card the total price is 634, with tickets 645)
// movie(100, 10, 0.95) should return 24
//     (with card the total price is 235, with tickets 240)
package main

import (
	"math"
)

func main() {
	var card, ticket int
	var perc float64
	Movie(card, ticket, perc)

func Movie(card, ticket int, perc float64) int {
	x := 1
	A := float64(ticket) * float64(x)
	B := float64(card) + float64(ticket)*math.Pow(float64(perc), float64(x))
	for A < B {
		A = float64(ticket) * float64(x)
		B = float64(card) + float64(ticket)*math.Pow(float64(perc), float64(x))
		x = x + 1
	}
	return x
}
// //func Movie(card, ticket int, perc float64) int {
//   var b float64 = float64(card)
//   var i int = 1
//   for i = 1; true; i++{
//     var a int = ticket * i
//     b = b + float64(ticket) * math.Pow(perc,float64(i))
//     if int(int(math.Ceil(b))) < a {
//       break
//     }
//   }
//
//   return i
// }
