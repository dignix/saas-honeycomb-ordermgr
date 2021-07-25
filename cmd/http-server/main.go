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

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/dignix/saas-honeycomb-ordermgr/cmd/internal"
	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/handler/rest"
	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/repository/mongodb"
	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/service"
)

const (
	defaultAddress  = "8091"
	shutdownTimeout = time.Second * 10
)

type serverConfig struct {
	Address     string
	Logger      *zap.Logger
	MongoDB     *mongo.Client
	Metrics     http.Handler
	Middlewares []mux.MiddlewareFunc
}

func main() {
	address, ok := os.LookupEnv("PORT")
	if !ok {
		address = defaultAddress
	}

	errC, err := run(":" + address)
	if err != nil {
		log.Fatalf("Failed to run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(address string) (<-chan error, error) {
	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("zap.NewProduction %w", err)
	}

	db, err := internal.NewMongoDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("internal.NewMongoDB %w", err)
	}

	srv, err := newServer(serverConfig{
		Address:     address,
		Logger:      logger,
		MongoDB:     db,
		Metrics:     nil,
		Middlewares: nil,
	})
	if err != nil {
		return nil, fmt.Errorf("newServer %w", err)
	}

	errC := make(chan error)

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

		defer func() {
			logger.Sync()
			db.Disconnect(ctxTimeout)
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		logger.Info("Shutdown completed")
	}()

	go func() {
		logger.Info("Listening and serving", zap.String("address", address))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, nil
}

func newServer(conf serverConfig) (*http.Server, error) {
	r := mux.NewRouter()

	for _, mwf := range conf.Middlewares {
		r.Use(mwf)
	}

	orderRepository := mongodb.NewMongoDBOrderRepository(conf.MongoDB)
	orderService := service.NewOrderService(orderRepository)
	orderEndpoint := rest.NewOrderRestEndpoint(orderService)
	orderEndpoint.BuildRoutes(r)

	return &http.Server{
		Addr:    conf.Address,
		Handler: r,
	}, nil
}
