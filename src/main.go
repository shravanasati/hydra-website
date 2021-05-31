package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func loadPage(page string) ([]byte, error) {
	content, e := ioutil.ReadFile("./templates/" + page)
	if e != nil {
		return nil, e
	}
	return content, nil
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		c, e := loadPage("index.html")
		if e != nil {
			panic(e)
		}
		w.Write(c)
	})
	
	port := os.Getenv("PORT")
	if len(port) == 0 { port = "8080" }
	log.Fatal(http.ListenAndServe(":"+port, nil))
}