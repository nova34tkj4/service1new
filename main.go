package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "selamat malam semuanya & sehat selalu & selamat hari kamis & dan besok hari jumat")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 4000")
    http.ListenAndServe(":4000", nil)
}
