package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/viavarejo-internal/pre-separacao-api/cmd/web/routes/handlers"
	"github.com/viavarejo-internal/pre-separacao-api/kit/config"
	logger "github.com/viavarejo-internal/pre-separacao-api/kit/log"
)

var (
	l *logrus.Entry
)

func init() {
	config.Load()

	logConfig := logger.Config{
		Environment: config.ENVIRONMENT,
		LogLevel:    config.LOG_LEVEL,
	}

	logFields := logrus.Fields{
		"app": "pre-separacao",
	}

	l = logger.New(logConfig).WithFields(logFields)
}

func setupRouter(docHandler *handlers.Document, stkHandler *handlers.Stockist) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/pre-separacoes", docHandler.ListAll())
		v1.GET("/pre-separacoes/:filial/:document", docHandler.ListOne())
		v1.PUT("/pre-separacoes/alterar-situacao", docHandler.ChangeSituation())
		v1.PUT("/pre-separacoes/:filial/:document/:situation", docHandler.ChangeSituationPDV())

		v1.GET("/estoquistas", stkHandler.ListAll())
	}
	return router
}

func main() {
	docHandler := handlers.NewDocument(l)
	stkHandler := handlers.NewStockist(l)

	router := setupRouter(docHandler, stkHandler)

	httpAddr := fmt.Sprintf("%s:%s", config.HOST_NAME, config.HOST_PORT)
	w := l.Writer()
	defer w.Close()

	srv := http.Server{
		Addr:           httpAddr,          // configure the bind address
		Handler:        router,            // ch(r),             // set the default handler
		ReadTimeout:    5 * time.Second,   // max time to read request from the client
		WriteTimeout:   10 * time.Second,  // max time to write response to the client
		IdleTimeout:    120 * time.Second, // max time for connections using TCP Keep-Alive
		ErrorLog:       log.New(w, "", 0), // set the logger for the server
		MaxHeaderBytes: 1 << 20,           // 1 MB
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		l.WithFields(logrus.Fields{"bind_address": httpAddr}).
			Info("Starting server")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.WithError(err).Fatal("Error starting server")
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-wait
	l.WithFields(logrus.Fields{"signal": sig.String()}).
		Info("Shutting down server gracefully...")
	l.Info("Press Ctrl+C again to force")

	// Create a deadline to wait for. waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), config.GRACEFUL_TIMEOUT)
	defer cancel()

	// gracefully shutdown the server doesn't block if no connections,
	// but will otherwise wait until the timeout deadline.
	srv.Shutdown(ctx)
	if err := srv.Shutdown(ctx); err != nil {
		l.WithError(err).Fatal("Server forced to shutdown")
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	l.Info("Server exiting")
	os.Exit(0)
}
