package main

import (
	"fmt"
	"math"
	"math/big"
)

// Для работы с действительно большими числами необходимо использовать big.Int или big.float
func main() {
	bigIntA := big.NewInt(int64(math.Pow(2, 21)))
	bigIntB := big.NewInt(int64(math.Pow(2, 22)))

	fmt.Println("a = ", bigIntA)
	fmt.Println("b = ", bigIntB)

	fmt.Println("Sum [add]\t\t = ", add(bigIntA, bigIntB))

	fmt.Println("Subtraction [sub]\t = ", subtract(bigIntA, bigIntB))

	fmt.Println("Multiplication [mul]\t = ", multiply(bigIntA, bigIntB))

	fmt.Println("Dividing [div] \t\t = ", divide(bigIntA, bigIntB))

}

// add сложение a и b
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// subtract вычитание b из a
func subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// multiply умножение a и b
func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// divide деление a на b
func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}
