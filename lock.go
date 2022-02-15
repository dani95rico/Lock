// EXAMPLE OF LOCK FUNCTIONALITY IN GO LENGUAGE BY DANI95RICO

package main

import (
	"fmt"
	"sync"
	"time"
)

type cuentaSegura struct {
	i int
	sync.Mutex
}

func main() {
	cs := new(cuentaSegura)

	for i := 0; i < 100; i++ {
		go cs.Increment()
		go cs.Decrement()
		fmt.Println(cs.GetValue())
	}
	time.Sleep(15 * time.Second)
}

func (cs *cuentaSegura) Increment() {
	cs.Lock()
	cs.i++
	cs.Unlock()
}

func (cs *cuentaSegura) Decrement() {
	cs.Lock()
	cs.i--
	cs.Unlock()
}

func (cs *cuentaSegura) GetValue() int {
	cs.Lock()
	v := cs.i
	cs.Unlock()
	return v
}
