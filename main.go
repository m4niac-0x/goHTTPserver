package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	fmt.Printf("URL : http://localhost:%s \nPress CTRL+C for interrupt...\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
