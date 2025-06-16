package bridge

import "fmt"

type Mac struct {
	Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.Printer = p
}
