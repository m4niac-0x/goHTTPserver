package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage : ./simpleHTTPserver <PORT>")
		os.Exit(1)
	}

	port := flag.String("p", os.Args[1], "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	fmt.Printf("URL : http://localhost:%s \nPress CTRL+C for interrupt...\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
