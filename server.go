package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("GET /{$}",handlerRoot)
	http.HandleFunc("GET /pi",handlerPI)
	http.HandleFunc("GET /circle/{radius}",handlerCircleArea)
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func handlerPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "3.14159265359")
}

func handlerCircleArea(w http.ResponseWriter, r *http.Request) {
	radiusStr := r.PathValue("radius")
	radius,err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid radius")
		return
	}
	fmt.Fprintf(w, "Area of the circle with radius %f is %f", radius, 3.14159265359 * radius * radius)
}
