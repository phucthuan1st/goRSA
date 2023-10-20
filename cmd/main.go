package main

import (
	"fmt"
	utils "goRSA/internal/utils"
	"math/big"
)

func main() {
	numberOfEquations := 3
	a := []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(4)}
	m := []*big.Int{big.NewInt(5), big.NewInt(7), big.NewInt(11)}

	x, N := utils.SolveCRT(uint(numberOfEquations), m, a)

	fmt.Printf("result of CRT equation is %v (mod %v)\n", x, N)

}
