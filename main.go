package main

import (
	"fmt"

	"github.com/arxeiss/go-false-positive-coverage/binary"
	"github.com/arxeiss/go-false-positive-coverage/unary"
)

func main() {
	fmt.Println("Example go code for testing ???")

	fmt.Printf("abs(3 * -8 + 4) = %f\n", unary.Abs(binary.Add(binary.Mul(3, -8), 4)))
}
