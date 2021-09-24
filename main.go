package main

import (
    "github.com/betNevS/build-own-web-framework/framework"
    "log"
    "net/http"
)

func main() {
    server := &http.Server{
        Handler: framework.NewCore(),
        Addr:":8888",
    }
    log.Fatal(server.ListenAndServe())
}
