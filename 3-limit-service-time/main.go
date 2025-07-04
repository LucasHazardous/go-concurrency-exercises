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
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  atomic.Uint64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	// if u.TimeUsed > 10 && !u.IsPremium {
	// 	return false
	// }

	ticker := time.NewTicker(1 * time.Second)

	completed := make(chan bool)
	
	go func() {
		process()
		completed<-true
	}()

	for {
		select {
		case <-ticker.C:
			u.TimeUsed.Add(1)
			if u.TimeUsed.Load() > 10 && !u.IsPremium {
				return false
			}
		case <-completed:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
