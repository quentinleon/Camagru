package main

import (
	"fmt"
	"net/http"
)

func stuff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Some stuff here.\n")
}

func main() {
	http.HandleFunc("/stuff", stuff)

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
