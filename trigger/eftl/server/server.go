package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		mime := r.Header.Get("Content-Type")
		log.Println("Request URI :", r.RequestURI)
		log.Println(mime)
		log.Println(string(body))
		w.Header().Set("Content-Type", mime)
		_, err = w.Write(body)
		if err != nil {
			log.Println(err)
		}
	}
	http.HandleFunc("/a", handler)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
