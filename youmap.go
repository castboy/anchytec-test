
import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

package main

import (
"fmt"
"math/rand"
"strconv"
"sync"
"time"
)

type Person struct {
	Age int
}

func MapPointerValue() {
	m := make(map[string]*Person)
	mtx := sync.RWMutex{}

	m["500"] = &Person{Age: 1}

	var z *Person
	fmt.Printf("%p\n", m["500"])
	z = m["500"]
	fmt.Printf("%p\n", z)


	for j := 0; j < 10000; j++ {
		go func() {
			mtx.Lock()
			i := rand.Intn(10000)
			if i != 500 {
				m[strconv.Itoa(i)] = &Person{Age: i}
			}

			mtx.Unlock()
		}()
	}

	for i := 0; i < 10000; i++ {
		go func(i int) {
			z.Age = i
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 10)
	fmt.Printf("%p\n", m["500"])
	fmt.Println(m["500"], z, len(m))
}
