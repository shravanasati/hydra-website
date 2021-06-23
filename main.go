package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func loadPage(page string) ([]byte, error) {
	content, e := ioutil.ReadFile("./templates/" + page)
	if e != nil {
		return []byte(strconv.Itoa(http.StatusInternalServerError)), e
	}
	return content, nil
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		c, e := loadPage("index.html")
		if e != nil {
			fmt.Println(e)
		}
		w.Write(c)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 { port = "8080" }
	fmt.Println("Listening on port:", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}