package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	port = os.Getenv("PORT")

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "200")
	})
	http.ListenAndServe(":"+port, nil)
}