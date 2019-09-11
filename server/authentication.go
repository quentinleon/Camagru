package main

import (
	"net/http"
)

func isLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("loggedInAs")
	if err == nil && verifyToken(c.Value) {
		return true
	}
	return false
}

func loggedInAs(r *http.Request) string {
	c, err := r.Cookie("loggedInAs")
	if err == nil {
		return getTokenContent(c.Value)
	}
	return "error"
}

func setAuthCookie(uID string, w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "loggedInAs",
		Value:    generateToken(uID),
		Path:     "/",
		MaxAge:   60 * 60,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func removeAuthCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "loggedInAs",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
