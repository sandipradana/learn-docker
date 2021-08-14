package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	instanceId := os.Getenv("INSTANCE_ID")
	if instanceId == "" {
		log.Fatal("INSTANCE_ID env is required")
	}

	httpHost := os.Getenv("HOST")
	if httpHost == "" {
		log.Fatal("HOST env is required")
	}

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		log.Fatal("PORT env is required")
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "instance %v, http server running on %v:%v", instanceId, httpHost, httpPort)
	})

	httpServer := http.Server{
		Addr:    fmt.Sprintf("%v:%v", httpHost, httpPort),
		Handler: httpMux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
