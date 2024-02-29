//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	done := make(chan struct{})
	ticker := time.NewTicker(1 * time.Millisecond)
	go func(done chan struct{}) {
		process()
		done <- struct{}{}
	}(done)
	for {
		select {
		case <-done:
			return true
		case <-time.After(10 * time.Second):
			if !u.IsPremium {
				return false
			}
		case <-ticker.C:
			u.TimeUsed += int64(1)
			if u.TimeUsed > 10*1000 && !u.IsPremium {
				return false
			}
		}
	}
}

func main() {
	RunMockServer()
}
