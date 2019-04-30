package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "test")
	})
}
