package kata

import (
	"math/big"
)

func AmIWilson(n int) bool {

	if n <= 1 {
		return false
	}

	factorialResult := new(big.Int)

	quotient := new(big.Int)
	remainder := new(big.Int)

	_, remainder = quotient.DivMod(factorialResult.Add(factorial(n-1), big.NewInt(1)), big.NewInt(int64(n*n)), remainder)

	return remainder.Cmp(big.NewInt(0)) == 0
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
