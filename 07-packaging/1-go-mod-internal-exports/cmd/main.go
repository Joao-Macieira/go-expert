package main

import (
	"fmt"

	"packaging-internal-exports/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
