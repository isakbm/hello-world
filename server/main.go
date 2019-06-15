package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	fmt.Println("starting server ...")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/index", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world!"))
	})

	if err := http.Serve(lis, mux); err != nil {
		log.Fatal(err)
	}

}
