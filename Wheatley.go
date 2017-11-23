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
    // "net/url"
    // "io"
    // "bytes"
)

type Page struct {
    Reg    string    `json:"reg"`
    Resp[] string `json:"resp"`
    Param bool `json:"param"`
    
}



func wheatleyResponse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    userInput := r.URL.Query().Get("value")

    // match := regexp.MustCompile("(?i)I need (.*)").MatchString("(?i)" + userInput)

    // fmt.Println(match);

    pages := getPages()

    for _, p := range pages {

        if ((regexp.MustCompile(p.Reg)).MatchString("(?i)" + userInput) ){
            response:= p.Resp[rand.Intn(len(p.Resp))]
            
            if(p.Param == true){
                match:= regexp.MustCompile(p.Reg).FindStringSubmatch("(?i)" + userInput)
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

                fmt.Fprintf(w, "<li>" + response + "</li>",updated) //.Path[1:])
            
            }else{
  	            fmt.Fprintf(w, "<li>" + response + "</li>") //.Path[1:])

            }
            return 
        }        
    }


    fmt.Fprintf(w, "<li>Sorry, I don't quite get what you mean. Would you like me to look it up?</li>")
//     fmt.Fprintf(w, "<li style='width:100%'>" +
//                         "<div class='msj-rta macro'>" +
//                             "<div class='text text-r'>" +
//                                 "<p>"+"Sorry, I don't quite get what you mean. Would you like me to look it up?"+"</p>" +
//                                 "<p><small>"+"23/11/17"+"</small></p>" +
//                             "</div>" +
//                         "<div class='avatar' style='padding:0px 0px 0px 10px !important'><img class='img-circle' style='width:100%;'/></div>" +                                
//                   "</li>")

// // src='+you.avatar+'"
    
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
    rand.Seed(time.Now().Unix())
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


// var baseUrl = "https://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"

// type Config struct {
// 	Language string
// 	Speak    string
// }

// type Speech struct {
// 	bytes.Buffer
// }

// func Speak(t Config) (*Speech, error) {
// 	req := fmt.Sprintf(baseUrl, url.QueryEscape(t.Speak), url.QueryEscape(t.Language))
// 	res, err := http.Get(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	speech := &Speech{}
// 	if _, err := io.Copy(&speech.Buffer, res.Body); err != nil {
// 		return nil, err
// 	}

// 	return speech, nil
// }