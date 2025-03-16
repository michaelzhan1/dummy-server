package runtime

import (
	"log"
	"net/http"

	"github.com/michaelzhan1/dummyserver/internals/handlers"
)

func SetupServer() *http.Server {
	port := "3000"
	log.Printf("Server is running on port %s", port)
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", handlers.Home)

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	return server
}
