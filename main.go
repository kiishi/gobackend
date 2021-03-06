package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "ok")
	})

	BootStrapValidator()
	BootstrapHandlers()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.Use(cors.Default())
	n.UseHandler(http.DefaultServeMux)

	log.Println("Starting server")
	err := http.ListenAndServe(":"+PORT, n)

	log.Fatal(err)

}
