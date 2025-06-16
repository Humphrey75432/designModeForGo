package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single2 struct {
}

var singleInstance2 *Single2

func GetInstance2() *Single2 {
	if singleInstance2 == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance2 == nil {
			fmt.Println("Creating single instance now.")
			singleInstance2 = &Single2{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance2
}
