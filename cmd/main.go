package main

import (
	"fmt"
	"gateway/pkg/api"
	xloader "github.com/mhthrh/common_pkg/pkg/loader"
	l "github.com/mhthrh/common_pkg/pkg/logger"
	cnfg "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/util/generic"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	configPath   = "src/gateway/config/file"
	appName      = "gateway"
	grpcPoolName = "user"
	url          = "https://vault.mhthrh.co.ca"
	secret       = "AnKoloft@~delNazok!12345"
	logName      = "x-app.gateway.service"
)

var (
	osInterrupt       chan os.Signal
	internalInterrupt chan error
)

func init() {
	osInterrupt = make(chan os.Signal)
	internalInterrupt = make(chan error)
	signal.Notify(osInterrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logger := l.NewLogger(logName)
	defer logger.LogSync()

	logger.Info(ctx, "Loading config...")

	config, err := xloader.New(url, configPath, "", "", secret, true)
	if err != nil {
		logger.Fatal(ctx, "config loader error", zap.Any("config loader failed", err))
	}
	err = config.Read()
	if err != nil {
		logger.Fatal(ctx, "config reader error", zap.Any("config loader failed", err))
	}
	logger.Info(ctx, "gateway config loaded successfully")

	srvConfig, err := config.GetServer()
	if err != nil {
		logger.Fatal(ctx, "config loader error", zap.Any("config loader failed", err))
	}
	logger.Info(ctx, "create gateway server")

	grpcs, err := config.GetGrpcs()
	if err != nil {
		logger.Fatal(ctx, "config loader error", zap.Any("config loader failed", err))
	}

	g := generic.Filter(grpcs, grpcPoolName, func(t cnfg.Grpc, s string) bool {
		if t.Srv == grpcPoolName {
			return true
		}
		return false
	})
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", srvConfig.Host, srvConfig.Port),
		Handler:      api.Run(false, logger, fmt.Sprintf("%s:%d", g.Host, g.Port), g.Count),
		ReadTimeout:  srvConfig.ReadTimeOut,
		WriteTimeout: srvConfig.WriteTimeOut,
		IdleTimeout:  srvConfig.IdleTimeOut,
	}
	logger.Info(ctx, "gateway server init successfully")

	go func() {
		defer log.Println("listener has been stopped")

		log.Printf("listener started and waiting for connection on %s:%d", srvConfig.Host, srvConfig.Port)
		if err := srv.ListenAndServe(); err != nil {
			internalInterrupt <- err
		}
		fmt.Println("kir")
	}()

	select {
	case <-osInterrupt:
		log.Println("OS interrupt received: shutting down server gracefully....")
		_ = srv.Shutdown(ctx)
	case err := <-internalInterrupt:
		log.Printf("Server listener encountered an error:%v shutting down....", err)
	}
	cancel()
}
