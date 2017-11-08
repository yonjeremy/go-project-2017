package main

import (
    "net/http"
    "text/template"
    "time"
    "math/rand"
)

type Message struct {  
    S string
    ChatMessages string
}


func templateHandler(w http.ResponseWriter, r *http.Request){
    
    // h2 header for game
    h2header := "This is wheatley"     

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    liChat:= "<li>this is a list yoke</li>"

    // setup the template to be executed
    msg  := &Message{S:h2header,ChatMessages:liChat}


    // tells the page the location of the template files and executes them
    t, _ := template.ParseFiles("template/chat.html")
    t.Execute(w,msg)

    

    

}


func main() {
    // generate seed for random number
    rand.Seed(time.Now().UTC().UnixNano())

    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/chat", templateHandler)
    http.ListenAndServe(":8080", nil)
}