package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
  http.ServeFile(rw, r, "login.html")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.Handle("/", r)
  log.Println("Server running on :8080")
  err := http.ListenAndServe(":8080", r)
  
  if err != nil {
     log.Printf("Error: %s\n", err.Error())
  }
}