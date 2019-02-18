// Return the number (count) of vowels in the given string.
//
// We will consider a, e, i, o, and u as vowels for this Kata.
//
// The input string will only consist of lower case letters and/or spaces.

package main

func main() {
	var str string
	var count int
	Lol := GetCount(str)
	x := GetCount(count)
}

func GetCount(str string) (count int) {
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "a" || string(str[i]) == "e" || string(str[i]) == "i" || string(str[i]) == "o" || string(str[i]) == "u" {
			count += 1
		}
	}
	return count
}
