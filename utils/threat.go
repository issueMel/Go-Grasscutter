package utils

import (
	"sync"
)

// todo CHECK: need threat pool ?

func DoAllFunc(wait *sync.WaitGroup, fn ...func()) {
	l := len(fn)
	wait.Add(l)
	for _, f := range fn {
		go func() {
			f()
			wait.Done()
		}()
	}
}
