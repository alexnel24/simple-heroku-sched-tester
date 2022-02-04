package app

import (
	"fmt"
	"net/http"
	"time"
)

func (s *Server) handleTimeGrab() http.HandlerFunc {
	fmt.Println("\nTIME-GRAB endpoint RUNNING")

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\nTIME-GRAB hit")
		now := time.Now()
		fmt.Println("Time is: " + now.String())
		_, err := w.Write([]byte("Time is: " + now.String()))
		if err != nil {
			fmt.Println("ERROR DURING TIME-GRAB")
		}
	}
}