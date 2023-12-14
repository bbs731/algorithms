package one_day_exercise

import (
	"fmt"
	"testing"
)

/*
n = 1000000000  a = 40000  b= 40000
 */
func Test878(t *testing.T) {
	//ans := nthMagicalNumber(4, 2, 3)
	ans := nthMagicalNumber(1000000000, 40000, 40000)
	fmt.Println(ans)
}
