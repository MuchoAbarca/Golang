package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type read0p struct {
	key int
	resp chan int
}

type write0p struct {
	key, val int
	resp chan bool
}

func main() {
	var read0ps uint64 = 0
	var write0ps uint64 = 0

	reads := make(chan *read0p)
	writes := make(chan *write0p)

	go func() {
		var state = make(map[int]int)
		for{
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &read0p{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&read0ps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &write0p{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&write0ps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	read0psFinal := atomic.LoadUint64(&read0ps)
	fmt.Println("read0ps:", read0psFinal)
	write0psFinal := atomic.LoadUint64(&write0ps)
	fmt.Println("write0ps:", write0psFinal)
}
