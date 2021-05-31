package main

import (
	"io/ioutil"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8000", nil))
}