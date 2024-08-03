# Mutex and RWMutex

## 1. Mutex

A Mutex is a mutual exclusion lock. It is used to protect shared data from being simultaneously accessed by multiple goroutines.

- Lock(): Acquires the lock. If the lock is already in use, the calling goroutine blocks until the lock is available.

- Unlock(): Releases the lock. If any other goroutines are waiting for the lock, one of them will be woken up and will acquire the lock.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var y = 0

func worker(wg *sync.WaitGroup, mu *sync.Mutex){
	mu.Lock()
	x = x + 1
	mu.Unlock()
	time.Sleep(time.Millisecond * 400)
	fmt.Println("x : ",x)
	wg.Done()
}

func worker2(wg *sync.WaitGroup, mu *sync.Mutex){
	mu.Lock()
	y = y + 1
	mu.Unlock()
	time.Sleep(time.Millisecond * 600)
	fmt.Println("y : ",y)
	wg.Done()
}

func main(){
	var w sync.WaitGroup

	var mu sync.Mutex

	for i:=0; i<10; i++{
		w.Add(2)
		go worker(&w,&mu)
		go worker2(&w,&mu)
	}

	w.Wait()
}
```

## 2. RWMutex

An RWMutex is a reader/writer mutual exclusion lock. It allows multiple readers or a single writer, but not both.

- RLock(): Acquires a read lock. Multiple readers can hold this lock simultaneously.
- RUnlock(): Releases a read lock.
- Lock(): Acquires a write lock. Only one writer can hold this lock, and no readers can hold it at the same time.
- Unlock(): Releases a write lock.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type safeMap struct {
	mu sync.RWMutex
	m map[int]int
}

func (s *safeMap) Read(key int)(int,bool){
	s.mu.RLock()
	val, ok := s.m[key]
	s.mu.RUnlock()
	return val, ok
}

func (s *safeMap) Write(key, value int){
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

func main(){
	s := safeMap{m : map[int]int{}}
	var wg sync.WaitGroup

	for i:=0; i<5;i++{
		wg.Add(1)
		go func (i int)  {
			s.Write(i, i*10)
			fmt.Printf("Written key: %d, value: %d\n", i, i*10)
			wg.Done()
		}(i)
	}

		for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			val, ok := s.Read(i)
			if ok {
				fmt.Printf("Read key: %d, value: %d\n", i, val)
			} else {
				fmt.Printf("Read key: %d not found\n", i)
			}
			wg.Done()
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
}
```
