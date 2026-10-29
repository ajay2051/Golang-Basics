// To avoid race conditions

package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	lock sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.lock.Unlock()
		wg.Done()	
	}()
	p.lock.Lock()
	p.views += 1
	// p.lock.Unlock() 
}

func main() {
	var wg sync.WaitGroup
	myPost := Post{views: 0}
	
	for i := 0; i < 100; i++{
		wg.Add(1)
		go myPost.inc(&wg)
	}
	
	wg.Wait()
	fmt.Println(myPost.views)
}