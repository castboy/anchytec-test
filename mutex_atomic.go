/*******************************************************************************
 * // Copyright AnchyTec Corp. All Rights Reserved.
 * // SPDX-License-Identifier: Apache-2.0
 * // Author: shaozhiming
 ******************************************************************************/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("service beginning...")
	runtime.GOMAXPROCS(2)

	len := 300
	total := 200000
	add := len * total
	t := make([]int32, len)
	wg := sync.WaitGroup{}
	wg.Add(add)
	begin := time.Now()
	if false {
		mux := sync.Mutex{}
		for i := 0; i < total; i++ {
			for j := 0; j < len; j++ {
				go func(j int) {
					mux.Lock()
					t[j]++
					mux.Unlock()
					wg.Add(-1)
				}(j)
			}
		}
	} else {
		for i := 0; i < total; i++ {
			for j := 0; j < len; j++ {
				go func(j int) {
					atomic.AddInt32(&t[j], 1)
					wg.Add(-1)
				}(j)
			}
		}
	}

	wg.Wait()
	fmt.Println(t, time.Now().Sub(begin))
}
