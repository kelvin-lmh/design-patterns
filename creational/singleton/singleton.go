package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var (
	singleInstance *singleton
	mu             sync.Mutex
)

// Lazy initialization of a singleton instance in Go
func GetInstance() *singleton {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			singleInstance = &singleton{}
			fmt.Println("Creating single instance now.")
		} else {
			fmt.Println("Instance already created, returning existing instance.")
		}
	} else {
		fmt.Println("Instance already created, returning existing instance.")
	}
	return singleInstance
}

// Thread-safe singleton pattern with sync.Once in Go
var once sync.Once

func GetThreadSafeInstance() *singleton {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now in a thread-safe manner.")
			singleInstance = &singleton{}
		})
	} else {
		fmt.Println("Instance already created, returning existing instance in a thread-safe manner.")
	}
	return singleInstance
}

// Init function can be used to initialize the singleton if needed
// init will run only once when the package is imported
func init() {
	fmt.Println("Initializing singleton instance...")
	singleInstance = &singleton{}
}

// GetInstance returns the singleton instance
func GetInstanceFromInit() *singleton {
	return singleInstance
}
