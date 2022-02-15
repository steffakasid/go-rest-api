package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/steffakasid/go-rest-api/pkg"
	"github.com/steffakasid/go-rest-api/pkg/middleware"
	"github.com/urfave/negroni"
)

const (
	PORT = 8080
)

func main() {

	n := negroni.New()
	setupRouter(n)

	log.Printf("Starting server at port %d", PORT)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), n))
}

func setupRouter(n *negroni.Negroni) {
	log.Println("Initializing Router!")
	router := pkg.NewRouter()
	n.Use(&middleware.Logger{})
	n.UseHandler(router)
}
