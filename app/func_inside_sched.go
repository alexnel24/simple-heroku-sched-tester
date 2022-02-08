package app

import "fmt"

func FakeFinder()(error) {
	fmt.Println("RUNNING THE FAKE FINDER FUNCTION INSIDE THE SCHEDULER")
	return nil
}