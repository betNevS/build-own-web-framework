package main

import (
	"log"
	"net/http"

	"github.com/betNevS/build-own-web-framework/framework"
)

func main() {
	core := framework.NewCore()
	RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	log.Fatal(server.ListenAndServe())
}
