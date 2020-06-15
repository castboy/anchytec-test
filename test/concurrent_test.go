package test

import (
	"os"
	"sync"
	"testing"
)

const Size = 1000000

func TestA(t *testing.T) {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	cache := make([]byte, 3)
	var rdOffset int64
	data := []byte{'a', 'b', 'c'}
	wg := &sync.WaitGroup{}
	wg.Add(Size)

	go func() {
		for {
			_, err := f.ReadAt(cache, rdOffset)
			if err != nil {
				//fmt.Println("rd err, ", err)
				continue
			}
			rdOffset += 3
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < Size; i++ {
			_, err := f.Write(data)
			if err != nil {
				//fmt.Println("wt err, ", err)
			}
		}
	}()

	wg.Wait()
}

func TestB(t *testing.T) {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	cache := make([]byte, 3)
	var rdOffset int64
	data := []byte{'a', 'b', 'c'}
	wg := &sync.WaitGroup{}
	wg.Add(Size)

	mux := sync.RWMutex{}

	go func() {
		for {
			mux.RLock()
			_, err := f.ReadAt(cache, rdOffset)
			if err != nil {
				//fmt.Println("rd err, ", err)
				mux.RUnlock()
				continue
			}
			rdOffset += 3
			wg.Done()
			mux.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < Size; i++ {
			mux.Lock()
			_, err := f.Write(data)
			if err != nil {
				//fmt.Println("wt err, ", err)
			}
			mux.Unlock()
		}
	}()

	wg.Wait()
}
