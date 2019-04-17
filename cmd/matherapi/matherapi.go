package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"mather/internal/api"
	"mather/internal/healthcheck"
	"mather/pkg/interceptors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	apiPortKey = "PORT"
	apiKeyKey  = "API_KEY"
)

func main() {
	loadConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt(apiPortKey)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := logrus.New()
	opt := interceptors.New(logger, viper.GetString(apiKeyKey))
	grpcServer := grpc.NewServer(opt)
	grpcServicer := &api.MatherAPI{}
	api.RegisterMatherServer(grpcServer, grpcServicer)

	httpServer := &http.Server{
		Addr: ":8000",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: healthcheck.New(grpcServicer),
	}

	sigs := make(chan os.Signal)
	errs := make(chan error)

	// Real API
	go func () {
		log.Println("RPC server listening on port", viper.GetInt(apiPortKey))
		if err := grpcServer.Serve(lis); err != nil {
			errs <- err
		}
	}()
	// Health Checker API
	go func () {
		if err := httpServer.ListenAndServe(); err != nil {
			errs <- err
		}
	}()

	select {
	case sig := <-sigs:
		logger.Infoln("sig", sig)
	case err := <-errs:
		logger.Errorln("error", err)
	}

	grpcServer.GracefulStop()
	_ = httpServer.Shutdown(context.Background())
}

func loadConfig() {
	// TODO: Could customise these as flags
	viper.AddConfigPath(".")
	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // Can customise config via Env
	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			break // If we are in the env var only case
		default:
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
