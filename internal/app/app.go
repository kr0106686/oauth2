package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kr0106686/oauth2/v2/config"
	"github.com/kr0106686/oauth2/v2/internal/controller/grpc"
	"github.com/kr0106686/oauth2/v2/internal/controller/http"
	"github.com/kr0106686/oauth2/v2/internal/repo/provider"
	"github.com/kr0106686/oauth2/v2/internal/repo/user"
	"github.com/kr0106686/oauth2/v2/internal/usecase/oauth"
	"github.com/kr0106686/oauth2/v2/pkg/gormx"
	"github.com/kr0106686/oauth2/v2/pkg/grpcserver"
	"github.com/kr0106686/oauth2/v2/pkg/httpserver"
)

func Run(cfg *config.Config) {

	db, err := gormx.New(gormx.Config(cfg.DB))
	if err != nil {
		log.Fatalf("%v", err)
	}

	providerRepo := provider.New(cfg.Provider)
	userRepo := user.New(db)

	oauthUseCase := oauth.New(providerRepo, userRepo, cfg.JWT)

	// gRPC Server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	grpc.NewRouter(grpcServer.App, oauthUseCase)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	http.NewRouter(httpServer.App, oauthUseCase)

	// Start servers
	grpcServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-grpcServer.Notify():
		log.Println(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = grpcServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}
}
