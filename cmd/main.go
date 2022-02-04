package main

import (
	"heroku_sched/app"
	"net/http"
)

func main(){
	app.Sched_run()
	server := app.NewServer()
	http.ListenAndServe(":8080", server)
}