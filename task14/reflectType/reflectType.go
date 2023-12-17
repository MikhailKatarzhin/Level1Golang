package main

import (
	"fmt"
)

func main() {
	var intValue int
	var stringValue string
	var boolValue bool
	var chAValue chan any
	var chIValue = make(chan int)
	var chSValue chan string
	var chBValue chan bool

	getType(intValue)
	getType(stringValue)
	getType(boolValue)
	getType(chAValue)
	getType(chIValue)
	getType(chSValue)
	getType(chBValue)
}

// При таком способе определения типа необходимо указывать тип в точности до типа данных канала
func getType(v interface{}) {
	fmt.Printf("\nЭто %T\n", v)
	switch v.(type) {
	case int:
		fmt.Println("Это int")
	case string:
		fmt.Println("Это string")
	case bool:
		fmt.Println("Это bool")
	case chan any:
		fmt.Println("Это channel any")
	case chan int:
		fmt.Println("Это channel int")
	case chan string:
		fmt.Println("Это channel string")
	case chan bool:
		fmt.Println("Это channel bool")
	}
}

/* Вывод

Это int
Это int

Это string
Это string

Это bool
Это bool

Это chan interface {}
Это channel any

Это chan int
Это channel int

Это chan string
Это channel string

Это chan bool
Это channel bool

*/
