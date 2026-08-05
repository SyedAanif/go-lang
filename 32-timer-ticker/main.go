package main

import (
	"fmt"
	"math/rand"
	"time"
)

// timer: is a way to signal after a specified duration

// ticker: runs after every n units of time interval
func main() {
	// // TIMER
	// timer := time.NewTimer(5 * time.Second)

	// a := <-timer.C // blocking operation

	// fmt.Println("Timer elapsed:", a)

	// // TICKER
	// ticker := time.NewTicker(time.Second) // run a ticker every second
	// go func() {
	// 	for {
	// 		// select { // this yields 5 ticks
	// 		// case t := <-ticker.C:
	// 		// 	fmt.Println("Ticker:", t)
	// 		// }
	// 		t := <-ticker.C // this yields 4 ticks
	// 		fmt.Println("Ticker:", t)
	// 	}
	// }()

	// time.Sleep(time.Second * 5) // to simulate processing
	// ticker.Stop()

	// CPU USAGE tracker
	// keeps track of cpu-usage and alerts if the cpu is greater than specified threshold for more than a specified duration
	ticker := time.NewTicker(time.Second) // check every second
	defer ticker.Stop()

	var alertTimerChan <-chan time.Time

	alertTimerActive := false // to send the alert
	go func() {
		for {
			select {
			case <-ticker.C: // whenever the ticker ticks
				// cpuUsage := rand.Intn(100) // simulate cpu usage
				cpuUsage := rand.Intn(100) + 100 // simulate cpu usage, high load
				fmt.Println("current CPU usage", cpuUsage)
				if cpuUsage > 80 { // beyond threshold
					if !alertTimerActive { // if alert timer is not active
						fmt.Println("High CPU Usage Detected starting alert timer", cpuUsage)
						alertTimer := time.NewTimer(10 * time.Second)
						alertTimerActive = true
						alertTimerChan = alertTimer.C
					}
				} else { // if cpu usage is normal
					if alertTimerActive {
						fmt.Println("cpu usage returned normal, stopping time", cpuUsage)
						alertTimerChan = nil
						alertTimerActive = false
					}
				}
			case <-alertTimerChan: // receive values from alert-timer
				if alertTimerActive {
					fmt.Println("ALERT: High CPU USage sustained for 10s")
					// reset the signal
					alertTimerActive = false
					alertTimerChan = nil
				}
			}
		}

	}() // this runs in separate go-routine

	select {} // block main thread forever
}
