package main

import (
	"flag"
	"github.com/go_chat/trace"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
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
	t.templ.Execute(w, r)
}


func main() {
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags
	gomniauth.SetSecurityKey("gMmSLu5Veqvei2sAYVUSoTwb")
	gomniauth.WithProviders(
		facebook.New("462137771096-q3uqmav00bi5d3atn42f4g2fr47jfc4a.apps.googleusercontent.com","gMmSLu5Veqvei2sAYVUSoTwb",
		"http://localhostL8080/auth/callback/facebook"),
		github.New("key","secret",
		"http://localhostL8080/auth/callback/github"),
		google.New("key","secret",
			"http://localhostL8080/auth/callback/google"),
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
