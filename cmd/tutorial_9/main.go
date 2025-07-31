package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
 * Goroutines are a way to lunch
 * multiple functions and have them executed concurently
 * Concurrency is not same as parallel execution
 * Concurrency means that having multiple tasks executing at the same time
 * One way of doing this is jumping back and forth from one task to another
 * In concurrent programming, while we wait for for example Database call,
 * CPU can move on to working on the other subtasks in the meantime
 *
 * But if we have 2 CPU Cores that are executing different tasks in parallel
 * Then the execution can happen simultaneously
 * The execution here is still concurrent because we have multiple tasks
 * executing at the same time but also these tasks are running in parallel
 * So this is true parallelysm
 *
 * In Go we have some level of parallelys as long as we have a multicore CPU
 */

// wait groups are basically just counters
var waitGroup = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
var mutex = sync.Mutex{} // Mutual exclusion

// Use RWMutex when we have Concurrent READ + WRITE
var RWmutex = sync.RWMutex{} // Readability, reader/writer mutual exclusion lock

// Alternatively we can define variables in groups like this:
// var (
// 	data   = "GENESIS"
// 	mutex  sync.Mutex
// 	result []string
// )

func dbCall(i int) {
	// Simulate DB Call delay
	var delay float32 = rand.Float32() * 2
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("The result from DB is: ", dbData[i], " with delay: ", delay)
	// One drawback is using normal mutex lock is that
	// it completely locks out other go routines
	// to accessing a result slice even for reading
	// So we are writing here but in case any other
	// go routine wants to read from results array,
	// it won't be able to read due to the normal mutex usage
	mutex.Lock()
	results = append(results, dbData[i])
	mutex.Unlock()
	waitGroup.Done() // Done, decrement the counter
}

func main() {
	// t0 := time.Now()
	// for i := range dbData {
	// 	// dbCall(i) // This will wait each and then output

	// 	waitGroup.Add(1) // Start, increment the counter
	// 	go dbCall(i)     // Concurrently works
	// }
	// waitGroup.Wait() // waits for counter to go back down to zero
	// fmt.Println("\nTotal execution time: %v", time.Since(t0))
	// fmt.Println("\nresults: ", results)

	// Better implementation with RW Mutex
	for range 10 {
		test()
		fmt.Println("\n\n")
	}

	var waitGroupConstant = sync.WaitGroup{}
	t_constant := time.Now()
	for range 1000 {
		waitGroupConstant.Add(1)
		go constantTimeFunc(&waitGroupConstant)
	}
	waitGroupConstant.Wait()
	fmt.Println("Total execution constantTimeFunc time: ", time.Since(t_constant))

	var waitGroupTimeConsuming = sync.WaitGroup{}
	t_time_consuming := time.Now()
	for range 1000 {
		waitGroupTimeConsuming.Add(1)
		go timeConsumingFunc(&waitGroupTimeConsuming)
	}
	waitGroupTimeConsuming.Wait()
	fmt.Println("Total execution timeConsumingFunc time: ", time.Since(t_time_consuming))

	// Shows the real thread count
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}

func constantTimeFunc(waitGroupConstant *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	waitGroupConstant.Done()
}

func timeConsumingFunc(waitGroupTimeConsuming *sync.WaitGroup) {
	var count int
	for range 100_000_000 {
		count += 1
	}
	waitGroupTimeConsuming.Done()
}

// ############################
// ############################
// ############################
// ############################
// ############################
//
//// RWMutex Performance Benefits in This Code:
// 1. MAIN PURPOSE: Prevents log() from reading while save() is writing
// 2. PERFORMANCE BOOST: Multiple log()/read calls can run simultaneously
// 3. SAFETY: Prevents reading half-written results2 data
//
// Basically,
// RWMutex allows multiple "READERS" to access the shared resource simultaneously,
// while only one writer can access it at a time.
// This improves performance
// by allowing concurrent reads while still ensuring data consistency.
//
// Normal Mutex: log() calls would wait in line (slower)
// RWMutex: Multiple log() calls run together simultaneously (faster)
//
// Example with 5 goroutines calling log():
// - Normal Mutex: 5 × 100ms = 500ms total (sequential)
// - RWMutex: 100ms total (parallel reading)

// wait groups are basically just counters
var waitGroupDB = sync.WaitGroup{}
var waitGroupLOG = sync.WaitGroup{}
var dbData2 = []string{"id1", "id2", "id3", "id4", "id5"}
var results2 = []string{}

// Use RWMutex when we have Concurrent READ + WRITE
var RWmutex2 = sync.RWMutex{} // Readability, reader/writer mutual exclusion lock

func dbCall2(i int) {
	// Simulate DB Call delay
	var delay float32 = rand.Float32() * 1
	time.Sleep(time.Duration(delay) * time.Second)
	// fmt.Println("The result from DB is: ", dbData2[i], " with delay: ", delay)
	save(dbData2[i]) // Write
	// log()             // Read
	waitGroupDB.Done() // Done, decrement the counter
}

func save(result string) {
	// Lock locks rw for Writing, not Reading!
	RWmutex2.Lock()
	results2 = append(results2, result)
	fmt.Println("Write Done")
	// Unlocks the locks rw for Writing
	RWmutex2.Unlock()
}

func log() {
	// This will always run even right after the return statement
	// multiple defers = LIFO (Last In, First Out) - like Stack
	// defer fmt.Println("3. works last")
	// defer fmt.Println("2. works in the middle")
	// defer fmt.Println("1. works first")
	defer waitGroupLOG.Done()

	// If we don't use waitGroup.Done()
	// It will DEADLOCK, we will get:
	// fatal error: all goroutines are asleep - deadlock!

	// RLock locks rw for reading
	RWmutex2.RLock()
	fmt.Println("Read: ", results2)
	// RLock unlocks rw for reading
	RWmutex2.RUnlock()

	// Alternatively, we can do it here
	// waitGroupLOG.Done()
}

func test() {
	t0 := time.Now()
	for i := range dbData2 {
		waitGroupDB.Add(1) // Start, increment the counter
		go dbCall2(i)      // Start Goroutines, Concurrently works

		waitGroupLOG.Add(1)
		// go keyword starts parallel work just like "async" keyword
		go log() // Read
	}

	waitGroupLOG.Wait()
	waitGroupDB.Wait() // waits for counter to go back down to zero

	fmt.Println("Total execution time: %v", time.Since(t0))
	fmt.Println("results2: ", results2, "\n\n\n\n")
	// Reset results
	results2 = []string{}
}
