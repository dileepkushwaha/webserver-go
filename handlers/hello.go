package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//Adding method which satisfies the interface http.Handler, it has no return parameters
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")
	// Adding response writer and request variables to the http handler.
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//rw.WriteHeader(http.StatusBadRequest)
		//rw.Write([]byte("Bazinga!"))
		//same things can be written replacing the above two lines and using the
		//http error package.
		http.Error(rw, "Bazinga!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
