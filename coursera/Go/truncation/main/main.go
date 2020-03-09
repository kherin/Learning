package main

import (
	"fmt"
	"math"
)


func main() {
	for {
		var truncated float64
		fmt.Println("Enter a floating-point number:")
		_, _ = fmt.Scan(&truncated)
		fmt.Println("Truncated value is: ", math.Trunc(truncated))
	}
}
