package utils

import (
	"math/big"
)

func SolveCRT(numberOfEquations uint, m []*big.Int, a []*big.Int) (*big.Int, *big.Int) {
	N := big.NewInt(1)

	for i := uint(0); i < numberOfEquations; i++ {
		N = N.Mul(N, m[i])
	}

	Ni := make([]*big.Int, numberOfEquations)
	for i := 0; i < int(numberOfEquations); i++ {
		Ni[i] = new(big.Int).Div(N, m[i])
	}

	result := new(big.Int)
	for i := 0; i < int(numberOfEquations); i++ {
		Mi := new(big.Int).ModInverse(Ni[i], m[i])

		prod := new(big.Int)
		prod = prod.Mul(a[i], Ni[i])
		prod = prod.Mul(prod, Mi)

		result = result.Add(result, prod)
		result = result.Mod(result, N)
	}

	return result, N
}
