package main

import (
	"fmt"
	"sync"
	"time"
)

func run() {
	time.Sleep(2 * time.Second)
	fmt.Println("run")
}

func run2() {
	time.Sleep(4 * time.Second)
	fmt.Println("run2")
}

func run3() {
	time.Sleep(6 * time.Second)
	fmt.Println("run3")
}

func addd(x int, y int, ch chan<- int, delay int) int {
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- x + y
	return 0
}

type Counter struct {
	value int
	lock  sync.Mutex
}

func count(counter *Counter, ch chan<- bool) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
	ch <- true
}

func count(counter *Counter, wg *sync.WaitGroup) {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.value++
	fmt.Println(counter.value)
	wg.Done()
}

func main() {
	go run()
	go run2()
	go run3()
	time.Sleep(7 * time.Second)
	fmt.Println("Done")

	ch := make(chan int)
	ch2 := make(chan int)
	go addd(1, 2, ch, 4)
	go addd(10, 20, ch2, 2)

	select {
	case x := <-ch:
		fmt.Println(x)
	case y := <-ch2:
		fmt.Println(y)
	}

	//	Buffer Channel
	bufCh := make(chan bool, 2)
	bufCh <- true
	<-bufCh

	//	Counter Channels
	counter := Counter{}
	cChan := make(chan bool)

	for i := 0; i < 10; i++ {
		go count(&counter, cChan)
	}

	for i := 0; i < 100; i++ {
		<-cChan
	}

	// Weight Group
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 10; i++ {
		go count(&counter, &wg)
	}
	wg.Wait()

}
