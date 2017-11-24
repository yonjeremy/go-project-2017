package main

import (
	"fmt"
    "net/http"
    "regexp"
    "time"
    "math/rand"
    "encoding/json"
    "io/ioutil"
    "os"
    "strings"

)

// create a struct that holds whatever that is parsed in from the json database file
type Page struct {
    Reg    string    `json:"reg"`
    Resp[] string `json:"resp"`
    Param bool `json:"param"`
    
}


// wheatley's response from the ajax request
func wheatleyResponse(w http.ResponseWriter, r *http.Request) {

    // seed for the random number
    rand.Seed(time.Now().Unix())

    // gets the value from the ajax query
    userInput := r.URL.Query().Get("value")

    // makes the json request
    pages := getPages()

    // for each object in the json file
    for _, p := range pages {

        // compare the regex of each json object and trigger it if it matches
        if ((regexp.MustCompile(p.Reg)).MatchString("(?i)" + userInput) ){

            // get the response from the json file
            response:= p.Resp[rand.Intn(len(p.Resp))]


            if(strings.Compare(p.Resp[0], "search") == 0){
                match:= regexp.MustCompile(p.Reg).FindStringSubmatch("(?i)" + userInput)

                fmt.Fprintf(w,"search" + match[1])
                return
            }
            
            // if there are parameters from the user input string that have been captured by the regex and needs to be passed through
            if(p.Param == true){
                // capture the required input
                match:= regexp.MustCompile(p.Reg).FindStringSubmatch("(?i)" + userInput)
                // split the captured words
                boundaries := regexp.MustCompile(`\b`)
                tokens := boundaries.Split(match[1], -1)


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
                    {`(?i)\byour\b`, `my`},
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

                // print the response
                fmt.Fprintf(w,response,updated)
            
            }else{
                // [rint the response]
  	            fmt.Fprintf(w,response)

            }
            return 
        }        
    }

    // error checking message
    fmt.Fprintf(w, "<li>Sorry, I don't quite get what you mean. Would you like me to look it up?</li>")
    
}




func main() {
    
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", wheatleyResponse)
	http.ListenAndServe(":8080", nil)
}

// gets access to wheatley's database
func getPages() []Page {
    // read in from database
    raw, err := ioutil.ReadFile("WheatleyDatabase.json")
    // check if file exists
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    // gets the information from json
    var c []Page
    json.Unmarshal(raw, &c)
    return c
}

