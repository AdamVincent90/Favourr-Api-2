package main

import (
	"log"
	"net/http"
)

const port string = ":4000"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write([]byte("Hello World!"))
		if err != nil {
			log.Fatalln(err)
		}
	})

	if err := http.ListenAndServe(port, http.NewServeMux()); err != nil {
		log.Fatalln("Error")
	}
}
