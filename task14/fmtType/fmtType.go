package main

import (
	"fmt"
	"reflect"
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

// При определении типа через рефлексию тип данных,
// можно округлить определение типа до канала, игнорируя тип данных канала
func getType(v interface{}) {

	fmt.Printf("\nЭто %s\n", reflect.TypeOf(v))
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("Это int")
	case reflect.String:
		fmt.Println("Это string")
	case reflect.Bool:
		fmt.Println("Это bool")
	case reflect.Chan:
		fmt.Println("Это channel")
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
Это channel

Это chan int
Это channel

Это chan string
Это channel

Это chan bool
Это channel

*/
