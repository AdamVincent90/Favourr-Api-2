package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

type Response struct {
	Title string
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	// log the request out to the console to see what is going on
	r.Use(middleware.Logger)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // front-end dev
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*/*"},
		AllowCredentials: true,
	}).Handler(r)

	fmt.Println("Hello World!")

	// thread this reflex server away from the main routine to not block the chi graphql handler server
	go func() {
		if err := http.ListenAndServe(":80", nil); err != nil {
			log.Println(err)
		}
	}()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {

		res := Response{
			Title: "Hello World!",
		}

		re, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
		}

		_, e := rw.Write(re)
		if e != nil {
			log.Println(e)
		}
	})

	if err := http.ListenAndServe(":4000", c); err != nil {
		log.Fatalln("Error")
	}

}
