package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type Single struct {
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
