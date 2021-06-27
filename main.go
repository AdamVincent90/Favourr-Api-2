package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

const port string = ":80"

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

	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatalln("Error")
		}
	}()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		_, e := rw.Write([]byte("Hello From GO API"))
		if e != nil {
			log.Println(e)
		}
	})

	if err := http.ListenAndServe(":4000", c); err != nil {
		log.Fatalln("Error")
	}

}
