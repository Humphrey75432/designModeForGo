package observer

import "fmt"

type Item struct {
	ObserverList []Observer
	Name         string
	InStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func (i *Item) Register(observer Observer) {
	i.ObserverList = append(i.ObserverList, observer)
}

func (i *Item) Deregister(observer Observer) {
	i.ObserverList = removeFromSlice(i.ObserverList, observer)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.Name)
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
