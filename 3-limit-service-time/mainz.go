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
type User1 struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest1(process func(), u *User) bool {
	ticker := time.NewTicker(1 * time.Second)

	completed := make(chan bool)
	var start, end *time.Time
	
	go func() {
		*start = time.Now()
		process()
		*end = time.Now()
		completed<-true
	}()

	for {
		select {
		case <-ticker.C:
			if !u.IsPremium {
				return false
			}
		case <-completed:
			return true
		}
	}
}

func main1() {
	RunMockServer()
}
