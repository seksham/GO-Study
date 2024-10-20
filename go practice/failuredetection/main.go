package main

import (
	"fmt"
	"time"
)

func Heartbeat(heartbeatChan chan<- struct{}, interval time.Duration) {
	for {
		time.Sleep(interval)
		select {
		case heartbeatChan <- struct{}{}:
			fmt.Println("Heartbeat sent")
		default:
			fmt.Println("Heartbeat channel is full, skipping")
		}
	}
}

func Monitor(heartbeatChan <-chan struct{}, timeout time.Duration) {
	for {
		select {
		case <-heartbeatChan:
			fmt.Println("Heartbeat received")
		case <-time.After(timeout):
			fmt.Println("Heartbeat timeout, failure detected")
			return
		}
	}
}
func main() {
	heartbeatChan := make(chan struct{}, 1)
	heartbeatInterval := 1 * time.Second
	heartbeatTimeout := 3 * time.Second

	// Start the heartbeat goroutine
	go Heartbeat(heartbeatChan, heartbeatInterval)

	// Start the monitor goroutine
	go Monitor(heartbeatChan, heartbeatTimeout)

	// // Enhancement: Add a way to gracefully shut down the goroutines
	// stopChan := make(chan struct{})
	// go func() {
	// 	time.Sleep(time.Second * 5)
	// 	fmt.Println("Simulating Ctrl+C signal")
	// 	// Create a channel to receive OS signals
	// 	sigChan := make(chan os.Signal, 1)
	// 	// Notify the channel for SIGINT (Ctrl+C)
	// 	signal.Notify(sigChan, os.Interrupt)
	// 	// Wait for the signal
	// 	<-sigChan
	// 	fmt.Println("Ctrl+C received, initiating shutdown")
	// 	close(stopChan)
	// }()

	// <-stopChan
	time.Sleep(time.Second * 10)
	fmt.Println("Main: shutting down")
}
