package http

import (
	"log"
	"net/http"
)

func RunFileServer() {
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello blog!"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
