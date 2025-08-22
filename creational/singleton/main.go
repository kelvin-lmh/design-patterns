package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// ----- Singleton (Lazy Initialization) -----
	fmt.Println("=== Singleton (Lazy Initialization) Demo ===")
	a := GetInstance()
	b := GetInstance()
	fmt.Printf("Same instance? %v\n\n", a == b)

	// Check concurrent access
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			GetInstance()
		}()
	}
	wg.Wait()

	// ----- Thread-safe Singleton with sync.Once -----
	fmt.Println("\n=== Thread-safe Singleton with sync.Once Demo ===")
	x := GetThreadSafeInstance()
	y := GetThreadSafeInstance()
	fmt.Printf("Same instance? %v\n\n", x == y)

	// Check concurrent access
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			GetThreadSafeInstance()
		}()
	}
	wg.Wait()

}
