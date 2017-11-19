package main

import (
	"fmt"
    "net/http"
    "regexp"
    // "text/template"
    // "time"
    // "math/rand"
    //  "io/ioutil"
    // "encoding/json"
    // "log"
    "encoding/json"
    "io/ioutil"
    "os"
)

type Message struct {  
    S string
    ChatMessages string
}
type test_struct struct {
    Test string
}

type Page struct {
    Reg    string    `json:"reg"`
    Resp string `json:"resp"`
    
}




// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13


func wheatleyResponse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")


 pages := getPages()
    for _, p := range pages {

        if ((regexp.MustCompile(p.Reg)).MatchString(r.URL.Query().Get("value")) ){
  	        fmt.Fprintf(w, "<li>%s</li>", p.Resp) //.Path[1:])
            return 
        }        
    }

    fmt.Fprintf(w,"hi" + r.URL.Query().Get("value"))

    fmt.Fprintf(w, "<li>Sorry, I don't quite get what you mean. Would you like me to look it up?</li>")


    
}



func (p Page) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}



func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", wheatleyResponse)
	http.ListenAndServe(":8080", nil)
}


func getPages() []Page {
    raw, err := ioutil.ReadFile("WheatleyDatabase.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Page
    json.Unmarshal(raw, &c)
    return c
}