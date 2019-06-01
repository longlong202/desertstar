package main

import (
        "fmt"
        "net/http"
)

func rootGateWay(w http.ResponseWriter, r *http.Request) {
        println("Welcome to Jacob's homepage! ")
}

func defaultGateWay(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to Jessie&Jacob's homepage!")
        println("Welcome to Jacob's homepage! ")
}

func main() {
        http.HandleFunc("/", defaultGateWay)
        http.ListenAndServe(":8080", nil)
}

