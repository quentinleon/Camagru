package main

import (
	"net/http"
)

var contentDir = "client/"

func index(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		http.Redirect(w, r, "/gallery", 303)
	} else {
		http.ServeFile(w, r, contentDir+"index.html")
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		http.Redirect(w, r, "/gallery", 303)
	} else {
		http.ServeFile(w, r, contentDir+"signup.html")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		http.Redirect(w, r, "/gallery", 303)
	} else {
		http.ServeFile(w, r, contentDir+"login.html")
	}
}

func gallery(w http.ResponseWriter, r *http.Request) {
	//TODO serve as template
	http.ServeFile(w, r, contentDir+"gallery.html")
}

func editor(w http.ResponseWriter, r *http.Request) {
	if !isLoggedIn(r) {
		http.Redirect(w, r, "/signup", 303)
	} else {
		//TODO serve as template
		http.ServeFile(w, r, contentDir+"editor.html")
	}
}

func account(w http.ResponseWriter, r *http.Request) {
	if !isLoggedIn(r) {
		http.Redirect(w, r, "/", 303)
	} else {
		//TODO serve as template
		http.ServeFile(w, r, contentDir+"account.html")
	}
}
