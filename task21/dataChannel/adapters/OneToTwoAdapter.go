package adapters

import (
	"github.com/MikhailKatarzhin/Level1Golang/task21/dataChannel"
	"github.com/MikhailKatarzhin/Level1Golang/task21/dataChannel/structs"
)

type OneToTwoAdapter struct {
	Bd *dataChannel.TypeOneBD
}

func (a OneToTwoAdapter) SendData() chan structs.TypeTwoStruct {
	at := make(chan structs.TypeTwoStruct)

	go func() {
		for data := range a.Bd.SendData() {
			at <- structs.TypeTwoStruct{
				Two:   data.Two,
				One:   data.One,
				Three: data.Three,
			}
		}

		close(at)
	}()

	return at
}

func (a OneToTwoAdapter) ReceiveData(at chan structs.TypeTwoStruct) {
	ao := make(chan structs.TypeOneStruct)

	go func() {
		for data := range at {
			ao <- structs.TypeOneStruct{
				One:   data.One,
				Two:   data.Two,
				Three: data.Three,
			}
		}

		close(ao)
	}()
}
