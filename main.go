package main

import (
	"designModeForGo/adapter"
	"designModeForGo/bridge"
	"designModeForGo/chain"
	"designModeForGo/command"
	"designModeForGo/composite"
	"designModeForGo/decorator"
	"designModeForGo/facade"
	"designModeForGo/factory"
	"designModeForGo/iterator"
	"designModeForGo/medium"
	"designModeForGo/memento"
	"designModeForGo/observer"
	"designModeForGo/prototype"
	"designModeForGo/proxy"
	"designModeForGo/singleton"
	"fmt"
	"log"
)

func main() {
	//	适配器模式
	client := &adapter.Client{}
	mac := &adapter.Mac{}
	client.InsertLightingConnectorIntoComputer(mac)
	windowsMachine := &adapter.Windows{}
	windowsAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}
	client.InsertLightingConnectorIntoComputer(windowsAdapter)

	// 桥接模式
	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &bridge.Windows{}
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()
	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()

	// 工厂模式
	ak47, _ := factory.GetGun("ak47")
	fmt.Printf("Gun: %s", ak47.GetName())
	fmt.Printf("Power: %d", ak47.GetPower())
	fmt.Println()

	musket, _ := factory.GetGun("musket")
	fmt.Printf("Gun: %s", musket.GetName())
	fmt.Printf("Power: %d", musket.GetPower())
	fmt.Println()

	//	原型模式
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("Printing hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("Printing hierarchy for clone folder")
	cloneFolder.Print(" ")

	//	单例模式
	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
		go singleton.GetInstance2()
	}

	//	组合模式
	file11 := &composite.File{Name: "file1"}
	file22 := &composite.File{Name: "file2"}
	file33 := &composite.File{Name: "file3"}
	folder11 := &composite.Folder{
		Name: "Folder1",
	}
	folder11.Add(file11)

	folder22 := &composite.Folder{
		Name: "Folder2",
	}
	folder22.Add(file22)
	folder22.Add(file33)
	folder22.Add(folder11)

	folder22.Search("rose")

	//	装饰器模式
	pizza := &decorator.VeggieMania{}
	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}
	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrince())

	//	门面模式
	walletFacade := facade.NewWalletFacade("abc", 1234)
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error:%s\n", err.Error())
	}
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error:%s\n", err.Error())
	}

	//	代理模式
	nginxServer := proxy.NewNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	//	责任链模式
	cashier := &chain.Cashier{}
	medical := &chain.Medical{}
	medical.SetNext(cashier)
	doctor := &chain.Doctor{}
	doctor.SetNext(medical)
	reception := &chain.Reception{}
	reception.SetNext(doctor)
	patient := &chain.Patient{Name: "abc"}
	reception.Execute(patient)

	//	命令模式
	tv := &command.Tv{}
	onCommand := &command.OnCommand{
		Device: tv,
	}
	offCommand := &command.OffCommand{
		Device: tv,
	}
	onButton := &command.Button{
		Command: onCommand,
	}
	onButton.Press()

	onButton = &command.Button{
		Command: offCommand,
	}
	onButton.Press()

	//	迭代器模式
	user1 := &iterator.User{
		Name: "a",
		Age:  30,
	}
	user2 := &iterator.User{
		Name: "b",
		Age:  20,
	}
	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}
	createIterator := userCollection.CreateIterator()
	for createIterator.HasNext() {
		user := createIterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

	// 中介者模式
	stationManager := medium.NewStationManager()
	passengerTrain := &medium.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &medium.FreightTrain{
		Mediator: stationManager,
	}
	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()

	// 备忘录模式
	caretaker := &memento.CareTaker{
		MementoArray: make([]*memento.Memento, 0),
	}
	originator := &memento.Originator{State: "A"}
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(2))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	// 观察者模式
	shirtItem := observer.NewItem("Nike Shirt")
	observeFirst := &observer.Customer{Id: "abc@email.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}
	shirtItem.Register(observeFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
}
