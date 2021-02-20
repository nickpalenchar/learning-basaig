package main

import (
  "log"
  "net/http"
  "fmt"
)

func main() {
  log.Println("This is not a web server")

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
  })

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("unable to start web server", err)
  }
}
