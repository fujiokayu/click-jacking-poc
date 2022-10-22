package main

import (
	"log"
	"net/http"
)

func main() {
	dir := http.FileServer(http.Dir(""))
	http.Handle("/", dir)
	log.Print("Open http://localhost:3000/")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
