# Wheatley Chatbot
Wheatley is a web-app chatbot created using GoLang.

## About
Name: Jeremy Yon
ID: G00330435

This is an adaptation of the Eliza AI chat bot coded in [Go](http://golang.org). This bot was created as part of my Data Representation and Querying Module in GMIT. The code here is adapted from [smallsurething.com](https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/) as well as this example provided by [Ian McLoughlin](https://github.com/data-representation/eliza).

## Wheatley
Whetley is an artificial intelligence robot that works in Apperture studios. He's a decent fella, just to annoy him too much!

### How it works
Wheatley is a  goLang web-app. It serves a HTML page as the root resource. 

He uses Ajax to communicate with the golang server. It parses all the user input using the json database file. 

Wheatley operates by using [RegeX](https://golang.org/pkg/regexp/) to recognise key words or phrases from the user input.

#### The program can be broken into the following steps:
          1. The user enters input.
          2. This input is then prepared for processing.
          3. The input is searched for keywords.
          4. Pronouns are swapped.
          5. The response is displayed in the chat window.
          
### Code Design
Instead of using a dat file, I decided to use a json file instead for the database. This proved to be a hassle, but I think it is a better way of coding. It uses objects instead of just reading the entire dat file, therefore sort of using object oriented methodologies.
Everything else works the same as the referenced eliza.go file from ian mcgloughlin's github.

### Extras
I incorporated a text to speech functionality, the user just has to check the text to speech box on the bottom right corner.

## Compilation
[Go](https://golang.org) must be installed to run the code.

### Clone this repo
```bash
git clone https://github.com/yonjeremy/go-project-2017
```
### Navigate to the folder
```bash
cd go-project-2017
```
### Build and the web app:
```go
go build Wheatley.go
```
### Run the exe:
```bash
./wheatley.exe
```
After running the exe open a browser and type in localhost and the port number (8080)
```bash
e.g localhost:8080/chat.html
```