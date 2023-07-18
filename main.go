package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pribadi1st/learn-go/handlers"
)

// Bad practice
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		data, err := ioutil.ReadAll(r.Body)
// 		log.Printf("Received: %s", err)
// 		if err != nil {
// 			http.Error(w, "Oops", http.StatusBadRequest)
// 			return
// 		}
// 		fmt.Fprintf(w, "Hello, %s", data)
// 	})

// 	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
// 		log.Printf("Goodbye")
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	// Handle custom timeout
	server := http.Server{
		Addr:        ":8080",
		Handler:     serveMux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Add gracefully shutdown
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Println("Received terminate, graceful shutdown", sig)

	timeOutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeOutContext)
}
