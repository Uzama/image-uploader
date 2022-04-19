package bootstrap

import (
	"context"
	"imageUploader/http/router"
	"imageUploader/http/server"
	"imageUploader/util/cmd"
	"imageUploader/util/config"
	"imageUploader/util/container"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(ctx context.Context) {

	// reading param from command line
	err := cmd.ReadParam()
	if err != nil {
		log.Fatalln(err)
	}

	// parsing configurations
	conf, err := config.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	// initializing repositories and adapters
	ctr, err := container.Resolve(conf)
	if err != nil {
		log.Fatalln(err)
	}

	// initializing the router
	r := router.Init(ctr)

	server := server.NewHTTPServer(conf, r)

	// start the server
	go server.ListnAndServe(ctx)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-c

	// graceful shutdown
	Destruct(ctx, ctr, server)

	os.Exit(0)
}

func Destruct(ctx context.Context, ctr container.Containers, server server.HTTPServer) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	log.Println("closing database connections")
	ctr.Adapters.Db.Close()

	go server.Shutdown(ctx)

	<-ctx.Done()

	log.Println("service shutdown gracefully")
}
