package main

import "net/http"

func signupsubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Redirect(w, r, "/stuff", 303)
	} else {
		index(w, r)
	}
}

//login handler

//post handler

//comment handler

//like handler
