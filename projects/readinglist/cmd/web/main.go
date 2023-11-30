package main

import (
	"flag"
	"net/http"
)

type application struct {
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")

	srv := &http.Server{
		Addr: *addr,
	}

	srv.ListenAndServe()
}
