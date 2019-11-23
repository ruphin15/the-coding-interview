package main

import (
	"fmt"
	"time"
)

func main() {
	f := debounce(talk, 1000)
	fmt.Println("1st time calling talk immediately")
	f() // blocked!
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2nd time calling talk after 2 secs")
	f() // Hello World!
	fmt.Println("3rd time calling talk immediately")
	f() // blocked@
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("4th time calling talk after 2 secs")
	f() // Hello World!
}

func debounce(f func(), timeoutMs time.Duration) func() {
	ticker := time.NewTicker(timeoutMs * time.Millisecond)

	return func() {
		select {
		case <-ticker.C:
			ticker = time.NewTicker(timeoutMs * time.Millisecond)
			f()
		default:
			fmt.Println("blocked!")
		}
	}
}

func talk() {
	fmt.Println("Hello World!")
}
