package server

import "net/http"

import "log"

func Serve() {
	log.Println("Started at port 5000!")
	log.Fatal(http.ListenAndServe(":5000", NewRouter()))
}
