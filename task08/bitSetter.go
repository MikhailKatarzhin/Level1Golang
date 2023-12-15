package main

import (
	"fmt"
)

func main() {
	var num int64 = 612

	bitPosition := 9 // Позиция бита, который нужно установить (начиная с 0)

	// Установка 0 в бит
	num = setBit(num, bitPosition, false)
	fmt.Printf("Число после установки 0 в бит %d:\t %d\n", bitPosition, num)

	bitPosition = 20

	// Установка 1 в бит
	num = setBit(num, bitPosition, true)
	fmt.Printf("Число после установки 1 в бит %d:\t %d\n", bitPosition, num)

}

// setBit устанавливает единицу или ноль в бит на указанную позицию числа
func setBit(number int64, pos int, setOne bool) int64 {
	if setOne {
		return setOneToBit(number, pos)
	} else {
		return setZeroToBit(number, pos)
	}
}

// setOneBit устанавливает единицу в бит на указанную позицию числа
//
//	через побитовое или с единицей, смещенной на указанную позицию
func setOneToBit(number int64, pos int) int64 {
	return number | int64(1<<pos)
}

// setZeroBit устанавливает ноль в бит на указанную позицию числа
//
//	через побитовое и с единицей, смещенной на указанную позицию и инвентированной
func setZeroToBit(number int64, pos int) int64 {
	return number & ^int64(1<<pos)
}
