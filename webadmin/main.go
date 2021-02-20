package main

import (
  "log"
  "net/http"
  "fmt"
  "time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "only GET method allowed", http.StatusMethodNotAllowed)
  }
  w.Write([]byte(time.Now().String()))
}

func isEven(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if time.Now().Second() % 2 == 0 {
      h.ServeHTTP(w, r)
    }
    http.Error(w, "Current time second is odd, cannot serve", 503)
  })
}

func main() {
  log.Println("This is not a web server")

  http.Handle("/iseven", isEven(http.HandlerFunc(getTime)))

  http.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
    w.Write([]byte("hello world"))
  })

  http.HandleFunc("/time", getTime)

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
  })

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("unable to start web server", err)
  }
}
