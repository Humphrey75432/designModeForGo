package adapter

import "fmt"

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lighting signal to USB.")
	w.WindowsMachine.insertIntoUSBPort()
}
