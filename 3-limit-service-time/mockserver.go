//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

// RunMockServer pretends to be a video processing service. It
// simulates user interacting with the Server.
func RunMockServer() {
	u1 := User{ID: 0, IsPremium: false}
	u2 := User{ID: 1, IsPremium: true}

	wg.Add(5)

	go createMockRequest(1, shortProcess, &u1)
	time.Sleep(1 * time.Second)

	go createMockRequest(2, longProcess, &u2)
	time.Sleep(2 * time.Second)

	go createMockRequest(3, shortProcess, &u1)
	time.Sleep(1 * time.Second)

	go createMockRequest(4, longProcess, &u1)
	go createMockRequest(5, shortProcess, &u2)
	wg.Wait()
	time.Sleep(10 * time.Second)
}

func createMockRequest(pid int, fn func(), u *User) {
	start := time.Now()
	fmt.Println("UserID:", u.ID, "\tProcess", pid, "started.")
	res := HandleRequest(fn, u)
	used := time.Since(start).Seconds()
	fmt.Println(runtime.NumGoroutine())
	if res {
		fmt.Println("UserID:", u.ID, "\tProcess", pid, "done.", " using time ", used)
	} else {
		fmt.Println("UserID:", u.ID, "\tProcess", pid, "killed. (No quota left)"," using time ", used)
	}

	wg.Done()
}

func shortProcess() {
	time.Sleep(6 * time.Second)
}

func longProcess() {
	time.Sleep(11 * time.Second)
	fmt.Println("still in backgroud")
}
