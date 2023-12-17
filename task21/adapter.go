package main

// DeviceInterface который реализует адаптер
type DeviceInterface interface {
	request() string
}

// LegacyObjectAdapter адаптер, реализующий интерфейс
type LegacyObjectAdapter struct {
	legacyObj *LegacyObject
}

func (a *LegacyObjectAdapter) request() string {
	response := a.legacyObj.specificRequest()
	// Какие-либо дополнительные обработки response полученные от legacyObj
	return response
}
