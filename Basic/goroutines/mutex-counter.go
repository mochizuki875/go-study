package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Int(key string) {
	c.mu.Lock() // 他のスレッドが同時にmap(c.v)にアクセスしないようにLock(排他制御:このスレッドしか触れない)

	c.v[key]++
	c.mu.Unlock() // (排他制御解除)
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock() // 他のスレッドが同時にmap(c.v)にアクセスしないようにLock(排他制御:このスレッドしか触れない)

	defer c.mu.Unlock() // Return直前に排他制御を解除
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Int("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
