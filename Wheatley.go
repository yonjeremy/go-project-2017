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
    "strings"
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
    Param bool `json:"param"`
    
}




// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13

var substitutions = map[string]string{
    "you":    "me",
    "your":   "my",
    "you're": "I am",
    "me":     "you",
    "I" : "you",
    "my": "your",
}


func wheatleyResponse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    userInput := r.URL.Query().Get("value")

    // match := regexp.MustCompile("(?i)I need (.*)").MatchString("(?i)" + userInput)

    // fmt.Println(match);

    pages := getPages()

    for _, p := range pages {

        if ((regexp.MustCompile(p.Reg)).MatchString("(?i)" + userInput) ){
            response:= p.Resp
            fmt.Println("hi")
            if(p.Param == true){
    boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(userInput, -1)

    for i := range tokens{
        fmt.Printf(tokens[i])
    }
	// List the reflections.
	reflections := [][]string{
		{`(?i)\bI\b`, `you`},
		{`(?i)\bme\b`, `you`},
		{`(?i)\byou\b`, `me`},
		{`(?i)\bmy\b`, `your`},
		{`(?i)\bam\b`, `are`},
        {`(?i)\bwas\b`, `were`},
        {`(?i)i'd`, `you would`},
        {`(?i)i'll`, `you will`},
        {`(?i)you've`, `I have`},
        {`(?i)you'll`, `I will`},
        {`(?i)\byours\b`, `mine`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	updated := strings.Join(tokens, ``)


                fmt.Fprintf(w, "<li>" + response + "</li>",updated) //.Path[1:])
            }else{
  	            fmt.Fprintf(w, "<li>" + response + "</li>") //.Path[1:])

            }
            return 
        }        
    }


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