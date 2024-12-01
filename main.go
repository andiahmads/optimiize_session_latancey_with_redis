package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"gored/commons/exception"
	"gored/commons/helper"
	"gored/commons/logger"
	"gored/commons/response"
	"gored/infra"
	"gored/router"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	failOnError := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s : %s", msg, err)
		}
	}

	dir := helper.DynamicDir()
	envPath := fmt.Sprintf("%s.env", dir)
	err := godotenv.Load(envPath)
	failOnError(err, "error load .env")

	db, err := infra.NewMySQLConn()
	failOnError(err, "error create mysql connection")
	defer db.Close()

	rdb, err := infra.NewRedisConn(0)
	failOnError(err, "error create redis connection")
	defer rdb.Close()

	port := os.Getenv("SERVICE_PORT")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		response.JsonEncoder[string](w, http.StatusOK, exception.Success, true, "hello world")
	})

	r.Mount("/auth", router.UserRouters())

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, //
		TLSConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	serverErr := make(chan error, 1)
	server.SetKeepAlivesEnabled(true)
	go func() {
		logger.Slogger().Info("starting server", slog.String("port", port))
		serverErr <- server.ListenAndServe()
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT)
	select {
	case sig := <-shutdownChannel:
		log.Printf("get signal -> %s", sig)
		logger.Slogger().Info("[GRACEFUL_SHUTDOWN_SPAWNED]", slog.Any("signal type", sig))
		timewait := 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timewait)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			server.Close()
		}
	case err := <-serverErr:
		failOnError(err, "server error")
	}

}
