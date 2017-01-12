package main
import (
    "fmt"
    "net/http"
    "strings"
    "github.com/stretchr/gomniauth"
)
type authHandler struct {
    next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    _,err := r.Cookie("auth")
    if err == http.ErrNoCookie  {
        w.Header().Set("Location","/login")
        w.WriteHeader(http.StatusTemporaryRedirect)
        return
    }
    if err != nil{
        // some other Error
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    //success - call the next Handler
    h.next.ServeHTTP(w, r)
}

func MustAuth(handler http.Handler) http.Handler {
    return &authHandler {next: handler}
}


// loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
    segs := strings.Split(r.URL.Path, "/")
    //TODO need to handle if segs has less then len 3
    action := segs[2]
    provider := segs[3]
    switch action{
    case "login":
        provider, err := gomniauth.Provider(provider)
        if err != nil{
            http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
            return
        }
        loginUrl, err := provider.GetBeginAuthURL(nil,nil)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s", provider, err), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", loginUrl)
        w.WriteHeader(http.StatusTemporaryRedirect)
        // w.WriteHeader(http.StatusNotImplemented)
        // log.Println("TODO handle login for", provider)
    default:
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "Auth action %s not supported", action)
    }
}