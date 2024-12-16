package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	tlsFlag := flag.Bool("tls", false, "Enable TLS (default: false)")
	port := flag.String("port", "443", "Port to listen on (default: 443)")
	certFile := flag.String("cert", "server.crt", "Path to TLS certificate file")
	keyFile := flag.String("key", "server.key", "Path to TLS key file")
	flag.Parse()

	// Define a handler function for the root URL path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Alive!")
	})

	useTLS := *tlsFlag
	address := "0.0.0.0:" + *port

	// Start the HTTP server
	if useTLS {
		fmt.Printf("Server is listening on https://%s\n", address)
		err := http.ListenAndServeTLS(address, *certFile, *keyFile, nil)
		if err != nil {
			fmt.Printf("Error starting TLS server: %s\n", err)
		}
	} else {
		fmt.Printf("Server is listening on http:// %s\n", address)
		err := http.ListenAndServe(address, nil)
		if err != nil {
			fmt.Printf("Error starting simple server: %s\n", err)
		}
	}
}
