package main

import (
	r "github.com/michaelzhan1/dummyserver/internals/runtime"
)

func main() {
	server := r.SetupServer()
	server.ListenAndServe()
}
