package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}

// func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	h.l.Println("Hello World")

// 	d, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(rw, "Oops", http.StatusBadRequest)
// 		return
// 	}
// 	// Format print F
// 	fmt.Fprintf(rw, "Hello %s", d)
// }
