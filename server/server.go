package main

import (
	"fmt"
	"net/http"
)

var staticDir = "client"

func ttinfo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("loggedInAs")
	if err != nil {
		fmt.Fprintf(w, "Not Authenticated")
	} else {
		if verifyToken(c.Value) {
			fmt.Fprintf(w, "Authenticated as : "+getTokenContent(c.Value))
		} else {
			fmt.Fprintf(w, "Invalid Token")
		}
	}

}

func tokenTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	http.ServeFile(w, r, "test_client/tokentest.html")
}

func auth(w http.ResponseWriter, r *http.Request) {
	setAuthCookie("hello", w, r)
}

func deauth(w http.ResponseWriter, r *http.Request) {
	removeAuthCookie(w, r)
}

func stuff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Some stuff here.\n")
}

func fallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		index(w, r)
	} else {
		http.ServeFile(w, r, "client/404.html")
	}
}

func main() {
	//Dev Stuff
	http.HandleFunc("/tt", tokenTest)
	http.HandleFunc("/ttinfo", ttinfo)
	http.HandleFunc("/a", auth)
	http.HandleFunc("/d", deauth)
	http.HandleFunc("/stuff", stuff)
	//end Dev stuff

	//content handlers
	http.HandleFunc("/index", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/editor", editor)
	http.HandleFunc("/account", account)
	//http.HandleFunc("/gallerycontent", gallerycontent)

	//post handlers
	http.HandleFunc("/signupsubmit", signupsubmit)

	//fallback
	http.HandleFunc("/", fallbackHandler)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
