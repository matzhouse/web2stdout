package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "port for the webserver to listen on")
}

func main() {

	flag.Parse()

	serverport := fmt.Sprintf(":%d", port)

	h := &Serveall{}

	m := http.NewServeMux()
	m.Handle("/", h)

	s := &http.Server{
		Addr:         serverport,
		Handler:      m,
		ReadTimeout:  10 * time.Second, // don't want to be waiting forever!
		WriteTimeout: 10 * time.Second, // don't want to be waiting forever!
	}

	log.Fatal(s.ListenAndServe())

}

type Serveall struct {
}

func (s *Serveall) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fmt.Println("####################")
	fmt.Println("Method: ", r.Method)

	fmt.Println("URL: ", r.RequestURI)

	fmt.Println("User Agent: ", r.UserAgent())

	for k, v := range r.Form {
		fmt.Printf("Form - %s = %s \n", k, v)
	}

	fmt.Println("")
}
