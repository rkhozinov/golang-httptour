package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
)

func randInt() int {
	return rand.Intn(10)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", html.EscapeString(r.URL.Path[1:]))
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi")
}

func randIntHadler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Intn(10))
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hi", hiHandler)
	http.HandleFunc("/randint", randIntHadler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
