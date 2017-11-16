package main

import (
	"fmt"
    "net/http"
    // "text/template"
    // "time"
    // "math/rand"
    //  "io/ioutil"
    // "encoding/json"
    // "log"
)

type Message struct {  
    S string
    ChatMessages string
}
type test_struct struct {
    Test string
}




// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13


func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //.Path[1:])
}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}