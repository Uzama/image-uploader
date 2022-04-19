package server

import (
	"context"
	"imageUploader/util/config"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// server
type HTTPServer struct {
	server  *http.Server
	address string
}

// creating new server
func NewHTTPServer(config config.Config, r *mux.Router) HTTPServer {

	address := config.App.Host + ":" + strconv.Itoa(config.App.Port)

	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,

		Handler: r,
	}

	httpServer := HTTPServer{
		server:  server,
		address: address,
	}

	return httpServer
}

// up and running the server
func (srv HTTPServer) ListnAndServe(ctx context.Context) {

	log.Printf("server listening on %s", srv.address)

	err := srv.server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// gracefully shutdown the server
func (srv HTTPServer) Shutdown(ctx context.Context) {

	log.Println("stropping HTTP server")

	srv.server.SetKeepAlivesEnabled(false)

	err := srv.server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}
