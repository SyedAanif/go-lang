package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SyedAanif/students-api/internal/config"
	"github.com/SyedAanif/students-api/internal/http/handlers/student"
	"github.com/SyedAanif/students-api/internal/storage/sqlite"
)

func main() {
	fmt.Println("Welcome to students api!!!")
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialised", slog.String("env", cfg.Env))

	// router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

	router.HandleFunc("GET /api/students", student.GetList(storage))

	// server setup

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// for graceful shutdown we serve ina go-routine, because we want control of program after ListenAndServe()
	// this program will move ahead if no channel/synchronisation

	slog.Info("server started", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)                                    // channel to catch signals of interrupt
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // notify the channel of interrupt signals

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done // getting blocking call till we receive a signal

	// graceful shutdown
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // empty context with 5s timeout
	defer cancel()                                                          // abandon the work once done
	// server.Shutdown() // direct shutdown, but we should control the timeout context

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown gracefully")
}
