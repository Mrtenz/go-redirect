package main

import (
  "net/http"
  "log"
  "os"
)

func handler (w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, os.Getenv("TARGET_URL"), 301)
}

func main() {
  if os.Getenv("TARGET_URL") == "" {
    log.Fatal("TARGET_URL is not set")
  }

  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
