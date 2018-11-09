// rectprops.go
//package rectangle
package oauth

import (
	"fmt"
	"math"
)

/*
* init funciton added
 */
func init() {
	fmt.Println("rectangle package initialized at === OAUTH ===")
}

func Area(len, wid float64) float64 {
	area := len * wid
	fmt.Println("=== Area ===")
	return area
}

func Diagonal(len, wid float64) float64 {
	fmt.Println("=== DIAGONAL ===")
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
