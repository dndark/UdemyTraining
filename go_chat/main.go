package main

import (
	"flag"
	"github.com/UdemyTraining/go_chat/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/stretchr/objx"
	"fmt"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	fmt.Println(r.Host)
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
	t.templ.Execute(w, r)
}

func main() {
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags
	gomniauth.SetSecurityKey("gMmSLu5Veqvei2sAYVUSoTwb")
	gomniauth.WithProviders(
		google.New("462137771096-q3uqmav00bi5d3atn42f4g2fr47jfc4a.apps.googleusercontent.com", "gMmSLu5Veqvei2sAYVUSoTwb",
			"http://localhost:8080/auth/callback/google"),
		github.New("key", "secret",
			"http://localhost:8080/auth/callback/github"),
		facebook.New("key", "secret",
			"http://localhost:8080/auth/callback/facebook"),
	)

	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
