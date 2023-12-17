package main

import (
	"github.com/MikhailKatarzhin/Level1Golang/task21/dataChannel"
	"github.com/MikhailKatarzhin/Level1Golang/task21/dataChannel/adapters"
	"github.com/MikhailKatarzhin/Level1Golang/task21/dataChannel/structs"
	"sync"
)

func main() {
	// Создаем экземпляр TypeOneBD
	typeOneBD := dataChannel.TypeOneBD{}

	// Создаем экземпляр OneToTwoAdapter
	adapter := adapters.OneToTwoAdapter{
		Bd: &typeOneBD,
	}

	// Отправляем данные из БД через адаптер в канал типа TypeTwoStruct
	sendChannel := adapter.SendData()

	var wg *sync.WaitGroup

	// Отправляем данные TypeTwoStruct в канал БД типа TypeOneStruct через адаптер
	typeTwoChannel := make(chan structs.TypeTwoStruct)
	wg.Add(1)

	go func() {
		for data := range sendChannel {
			println("Received TypeTwoStruct data on dst: ", data)
			typeTwoChannel <- data
			println("Send TypeOneStruct data to BD: ", data)
		}
		close(typeTwoChannel)
		wg.Done()
	}()

	adapter.ReceiveData(typeTwoChannel)

	wg.Wait()
}
