package server

import "net/http"

func Start() error {
	server := http.Server{Addr: "localhost:8080"}

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hei der!"))
	}

	server.Handler = handler

	return server.ListenAndServe()
}
//
//func getTrack {
//
//}
