package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", router)
	//http.HandleFunc("/subpath", hdlSub)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var numericPath,_ = regexp.Compile("[0-9]")

var numericSuffix,_ = regexp.Compile("[0-9]$")

var numericPrefix,_ = regexp.Compile("^/[0-9]")

var nestedPath,_ = regexp.Compile("^/.*/[a-z]")

func router(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch {
	case numericPrefix.MatchString(r.URL.Path):
		hdlNumPre(w, r)
	case numericSuffix.MatchString(r.URL.Path):
		hdlNumSuf(w, r)
	case numericPath.MatchString(r.URL.Path):
		hdlNum(w, r)
	case nestedPath.MatchString(r.URL.Path):
		hdlNested(w, r)
	default:
		hdlMain(w, r)
	}
}

func hdlNum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on numeric path handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlNested(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on nested path handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlNumPre(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on numeric prefix path handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlNumSuf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on numeric suffix path handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on default handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlSub(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on sub handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}