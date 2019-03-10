package main

import (
	"fmt"
	"net"
	"net/http"
)

func ip(w http.ResponseWriter, r *http.Request) {
	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		fmt.Fprint(w, host)
	}
}