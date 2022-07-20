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

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/gen/api"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/interfaces/handlers"
	"github.com/Code0716/go-vtm/app/registry"
	"github.com/Code0716/go-vtm/app/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	code := start()
	os.Exit(code)
}

func start() int {
	env := util.Env()

	dbConn, close, err := db.NewDBConn(env)
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
		return 1
	}
	db, err := db.NewDB(dbConn, env)
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
		return 1
	}

	defer func() {
		if err := close(); err != nil {
			log.Fatal(err)
		}
	}()

	reg := registry.New(db)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	e.Use(middleware.Recover())

	// helth check
	e.GET("/healthz", handlers.GetHealthz(reg))

	newHandlers := handlers.New(reg)

	e.POST("/admin/regist", newHandlers.RegistAdmin)
	e.POST("/admin/login", newHandlers.Login)

	// v1
	router := e.Group("/api/v1")

	router.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(env.Signingkey),
		TokenLookup: "header:authorization",
		Claims:      &domain.JwtCustomClaims{},
	}))

	api.RegisterHandlersWithBaseURL(router, newHandlers, "")

	if env.EnvCode == "local" {
		for _, r := range e.Routes() {
			if r.Path == "" || r.Path == "/api/v1" || r.Path == "/api/v1/*" {
				continue
			}
			fmt.Printf("[%v] %+v\n", r.Method, r.Path)
		}
	}
	addr := util.GetAPIPath(env)

	srv := &http.Server{
		Addr:              addr,
		Handler:           e,
		ReadHeaderTimeout: 20 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("unexpected shutting down: %w", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	defer stop()

	<-ctx.Done()
	log.Printf("server shutting down on %v", ctx.Err())

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Server forced to shutdown: %w", err)
	}

	return 0
}
