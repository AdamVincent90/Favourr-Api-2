package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":80"

func main() {

	a := http.NewServeMux()

	fmt.Println("Hello World!")

	a.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write([]byte("Hello Geezer!"))
		if err != nil {
			log.Fatalln(err)
		}
	})

	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalln("Error")
		}
	}()

	if err := http.ListenAndServe(":4000", a); err != nil {
		log.Fatalln("Error")
	}

}
