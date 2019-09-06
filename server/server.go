package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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
	} else {
		staticAsset(w, r)
	}
}

func redirectURL(w http.ResponseWriter, r *http.Request, url string, loggedIn bool) bool {
	if loggedIn == true {
		if url == "/signup.html" || url == "/login.html" || url == "/" {
			http.Redirect(w, r, "/gallery", 303)
			return true
		}
	} else {
		switch url {
		case "/editor.html":
			http.Redirect(w, r, "/signup", 303)
			return true
		case "/account.html":
			http.Redirect(w, r, "/", 303)
			return true
		}
	}
	return false
}

func staticAsset(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("loggedInAs")
	loggedIn := false
	if err == nil {
		loggedIn = verifyToken(c.Value)
	}

	//append .html to any sources
	if !strings.HasSuffix(r.URL.Path, ".html") && r.URL.Path != "/" {
		r.URL.Path += ".html"
	}

	if redirectURL(w, r, r.URL.Path, loggedIn) {
		return
	}

	//serve fileserver content
	if _, err := os.Stat(staticDir + r.URL.Path); err == nil {
		http.ServeFile(w, r, staticDir+r.URL.Path)
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

/*func main() {
	db, err := sql.Open("mysql", "root:strongpass@/dbname")
}*/
