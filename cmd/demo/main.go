package main

import (
	"net/http"
	// "os"
	// "os/signal"
	"time"
	// "log"
	// "context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "github.com/lesismal/nbio/nbhttp"
	// "github.com/lesismal/nbio/nbhttp/websocket"

	"github.com/orsanawwad/htmxdemo/internal/handlers"
	"github.com/orsanawwad/htmxdemo/internal/services"
)

func main() {

	ts := services.New()
	mainHandler := handlers.NewMainHandler()
	assetsHandler := handlers.NewAssetsHandler()
	todosHandler := handlers.NewTodosHandler(ts)
	// wsHandler := handlers.NewWSHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Mount("/", mainHandler.Routes())
	r.Mount("/assets", assetsHandler.Routes())
	r.Mount("/todos", todosHandler.Routes())
	// r.Mount("/ws", wsHandler.Routes())

	// server := nbhttp.NewServer(nbhttp.Config{
	// 	Network: "tcp",
	// 	Addrs:         []string{"localhost:3000"},
	// 	Handler:      r,
	// 	// ReadTimeout:  time.Second * 10,
	// 	WriteTimeout: time.Second * 10,
	// })

	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	server.ListenAndServe()
	// err := server.Start()
	// if err != nil {
	// 	log.Fatalf("nbio.Start failed: %v\n", err)
	// }

	// log.Println("serving [go-chi/chi] on [nbio]")

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)
	// <-interrupt

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()
	// server.Shutdown(ctx)
}
