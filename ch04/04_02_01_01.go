package main

import "fmt"

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		done <- true
	}()
	<-done
	fmt.Println("all tasks are finished")
}
