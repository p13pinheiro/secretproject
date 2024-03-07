package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("pod hit")
			fmt.Fprintf(w, "hello\n")
		},
	)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Panic(err)
	}
}
