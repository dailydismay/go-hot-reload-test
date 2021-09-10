package main

import (
	"context"
	"fmt"
	"gofirstapp/internal/config"
	"gofirstapp/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	r := mux.NewRouter()

	
	errCh := make(chan error)
	
	cfg, err := config.NewConfig()

	if err != nil {
		errCh <- err
	}

	s := &http.Server{
		Handler: r,
		Addr: fmt.Sprintf(":%s", cfg.Port),
	}

	h := handlers.NewHandlers()
	r.HandleFunc("/kanye", h.Kanye)
	
	go func ()  {
		log.Println("Serving kanye quotes")
		errCh <- s.ListenAndServe()
	}()

	select {
	case e := <- errCh:	
		log.Fatal(e)
	case <- ch:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancel()

		err := s.Shutdown(ctx)

		if err != nil {
			log.Fatal(err)
			err = s.Close()
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Shutting down...")
	}
}
