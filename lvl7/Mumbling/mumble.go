// This time no story, no theory. The examples below show you how to write function accum:
//
// Examples:
//
// accum("abcd") // -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") // -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") // -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.
//
// package kata
//
// func Accum(s string) string {
//     // your code
// }
package main

import (
	d "strings"
)

func main() {
	var s string
	Accum(s)
}

func Accum(s string) string {
	var x string
	for i := 0; i < len(s); i++ {
		Up := d.ToUpper(string(s[i]))
		low := d.ToLower(string(s[i]))
		x += "-" + Up + d.Repeat(low, i)
	}
	return x[1:]
}
