package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/ip", ip)
	log.Fatal(http.ListenAndServe(net.JoinHostPort(*host, strconv.Itoa(*port)), nil))
}