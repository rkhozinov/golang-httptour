package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const INT_RAND_RANGE int = 100
const HTTP_SERVER_PORT = ":80"
const BASE_TEMPLATE = "template.tpl"

var templates = template.Must(template.ParseFiles(BASE_TEMPLATE))

type Page struct {
	Title   string
	Content []string
}

func checkHttp(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", html.EscapeString(r.URL.Path[1:]))
}

func randIntHndlr(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "randInt",
		Content: []string{strconv.Itoa(rand.Intn(INT_RAND_RANGE))},
	}
	err := templates.ExecuteTemplate(w, BASE_TEMPLATE, &page)
	checkHttp(w, err)
}

func randFloatHndlr(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "randFloat",
		Content: []string{strconv.FormatFloat(rand.Float64(), 'f', -1, 64)},
	}
	err := templates.ExecuteTemplate(w, BASE_TEMPLATE, &page)
	checkHttp(w, err)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/randint", randIntHndlr)
	http.HandleFunc("/randfloat", randFloatHndlr)
	log.Fatal(http.ListenAndServe(HTTP_SERVER_PORT, nil))
}
