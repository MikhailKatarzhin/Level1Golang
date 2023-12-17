package dataChannel

type DataChanel interface {
	SendData() chan any
	ReceiveData(chan any)
}
