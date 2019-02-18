// Take a number: 56789. Rotate left, you get 67895.
//
// Keep the first digit in place and rotate left the other digits: 68957.
//
// Keep the first two digits in place and rotate the other ones: 68579.
//
// Keep the first three digits and rotate left the rest: 68597. Now it is over since keeping the first four it remains only one digit which rotated is itself.
//
// You have the following sequence of numbers
//
// 56789 -> 67895 -> 68957 -> 68579 -> 68597
//
// and you must return the greatest: 68957.
//
// Calling this function max_rot (or maxRot or ... depending on the language
//
// max_rot(56789) should return 68957
package main
import (
  "strconv"
  "fmt"
)

func main() {
var n int64 = 56789
Lol := MaxRot(n)
fmt.Println(Lol)
}

func MaxRot(n int64) int64 {
b := strconv.Itoa(int(n))
x := b[1:] + string(b[0])
y := b[:1] + string(b[2:]) + string(b[1])
z := b[:2] + b[3:] + string(b[2])
a := b[:3] + b[4:] + string(b[3])

my, _ := strconv.Atoi(b)
dick, _ := strconv.Atoi(x)
so, _ := strconv.Atoi(y)
fc, _ := strconv.Atoi(z)
big, _ := strconv.Atoi(a)

if int(my) > int(dick) && int(my) > int(so) && int(my) > int(fc) && int(my) > int(big){
return int64(my)
}else if int(dick) > int(my) && int(dick) > int(so) && int(dick) > int(fc) && int(dick) > int(big){
return int64(dick)
}else if int(so) > int(dick) && int(so) > int(my) && int(so) > int(fc) && int(so) > int(big){
return int64(so)
}else if int(fc) > int(dick) && int(fc) > int(so) && int(fc) > int(my) && int(fc) > int(big){
return int64(fc)
}else if int(big) > int(dick) && int(big) > int(so) && int(big) > int(fc) && int(big) > int(my){
return int64(big)
}
return 0
}
