package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/projectdiscovery/gologger"
	"srebootcamp/internal/config"
	"srebootcamp/internal/db"
	"srebootcamp/internal/handler"
	"srebootcamp/internal/router"
)

func main() {
	c := config.LoadConfig()
	dbConn := db.InitDB(c)
	defer dbConn.Close()

	h := &handler.Handler{DB: dbConn}
	r := router.SetupRouter(h)

	srv := &http.Server{
		Addr:    c.Port,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			gologger.Fatal().Msgf("server failed to start: %v", err)
		}
	}()
	gologger.Info().Msgf("listening on %s", c.Port)

	<-ctx.Done()
	gologger.Info().Msg("shutdown signal received, draining in-flight requests...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		gologger.Warning().Msgf("graceful shutdown did not complete cleanly: %v", err)
	}
}
