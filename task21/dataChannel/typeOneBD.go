package dataChannel

import (
	"github.com/MikhailKatarzhin/Level1Golang/task21/dataChannel/structs"
)

type TypeOneBD struct{}

func (t TypeOneBD) SendData() chan structs.TypeOneStruct {
	ch := make(chan structs.TypeOneStruct)

	go func() { // Пример данных для отправки
		for i := 0; i < 10; i++ {
			dataToSend := structs.TypeOneStruct{
				One:   1,
				Two:   2.0,
				Three: "three",
			}
			ch <- dataToSend
			println("BD send TypeOneStruct data:", dataToSend)
		}

		close(ch)
	}()
	return ch
}

func (t TypeOneBD) ReceiveData(ch chan structs.TypeOneStruct) {
	for data := range ch {
		print("BD received data:", data)
	}
}
