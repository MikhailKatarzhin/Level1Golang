package main

import (
	"fmt"
)

func main() {
	legacyObj := &LegacyObject{}
	adapter := &LegacyObjectAdapter{legacyObj}

	device := &newDevice{}
	device.someDeviceCode(adapter)
}

// newDevice структура устройства, взаимодействующего через адаптер со старым объектом
type newDevice struct{}

// Функция, использующая адаптер
func (d newDevice) someDeviceCode(target DeviceInterface) {
	fmt.Println(target.request())
}
