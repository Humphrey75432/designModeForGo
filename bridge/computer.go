package bridge

type Computer interface {
	Print()
	SetPrinter(printer Printer)
}
