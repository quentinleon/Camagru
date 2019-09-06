package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var staticDir = "client/"

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
	http.ServeFile(w, r, "test_client/tokentest.html")
}

func auth(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "loggedInAs",
		Value:    generateToken("hello"),
		Path:     "/",
		MaxAge:   60 * 60,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func deauth(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "loggedInAs",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func stuff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Some stuff here.\n")
}

func signupsubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Redirect(w, r, "/stuff", 303)
	}
}

func staticAsset(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".html") && r.URL.Path != "/" {
		r.URL.Path += ".html"
	}
	if _, err := os.Stat(staticDir + r.URL.Path[1:]); err == nil {
		http.ServeFile(w, r, staticDir+r.URL.Path[1:])
	} else if os.IsNotExist(err) {
		w.WriteHeader(404)
		http.ServeFile(w, r, staticDir+"404.html")
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

	http.HandleFunc("/signupsubmit", signupsubmit)

	http.HandleFunc("/", staticAsset)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
