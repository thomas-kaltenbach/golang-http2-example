package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a server on port 8000
	// Exactly how you would run an HTTP/1.1 server
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// Start the server with TLS, since we are running HTTP/2 it must be
	// run with TLS.
	// Exactly how you would run an HTTP/1.1 server with TLS connection.
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection: %s", r.Proto)

	isWebsocket := false
	for headerName, headerValue := range r.Header {
		log.Print(headerName + ":" + headerValue[0])
		if headerName == "Upgrade" && headerValue[0] == "websocket" {
			isWebsocket = true
		}
	}

	// route to app ...
	// forward respond
	if isWebsocket {
		log.Print("is websocket... handle connection")
		//...
		w.WriteHeader(101)
	} else {
		log.Print("no websocket... handle request")
		w.WriteHeader(200)
		w.Write([]byte("Hello"))
	}
}
