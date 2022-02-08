package main

import (
	"heroku_sched/app"
	"net/http"
	"os"
)

func main(){
	// app.Sched_run()
	port := os.Getenv("PORT")
	server := app.NewServer()
	http.ListenAndServe(":" + port, server)
}