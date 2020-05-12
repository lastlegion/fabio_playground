package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	host string
	port int
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from books service\n")
}

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "binding host")
	flag.IntVar(&port, "port", 8080, "binding port")
	flag.Parse()
}

func main() {
	log.Print("Starting grocery service...")
	http.HandleFunc("/book", handler)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("HTTP server up on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
