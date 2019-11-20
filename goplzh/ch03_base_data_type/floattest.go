package main

import (
	"fmt"
	"math"
)

func main()  {
	var f float32 = 16777216
	fmt.Println(f == f + 1)

	const Avogadro = 6.02214e32
	const Planck = 6.2323e-32
	fmt.Println(Avogadro)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f %[2]g\n", x, math.Exp(float64(x)))
	}

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}