package main

import (
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/config"
	grpcserver "github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/grpc"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/storage"
	_ "github.com/a1ekaeyVorobyev/otus_go_hw/hw21/pkg/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	var configFile string
	flag.StringVar(&configFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if configFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "don't config file")
		os.Exit(2)
	}

	conf, err := config.ReadFromFile(configFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	logger, f := logger.GetLogger(conf)
	if f != nil {
		defer f.Close()
	}

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	inFile := storage.InFile{}
	inFile.Init()
	cal := calendar.Calendar{Config: conf, Storage: &inFile, Logger: &logger}
	//grpcServer := 	grps.Server{conf,&logger,&cal} get error too few values ?
	grpcServer := grpcserver.Server{}
	grpcServer.Calendar = &cal
	grpcServer.Config = conf
	grpcServer.Logger = &logger
	go web.RunServer(conf, &logger)
	go grpcServer.Run()

	defer grpcServer.Shutdown()

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

exit:
	for {
		select {
		case c := <-sigs:
			logger.Infof("Got signal: %v. Exit.", c)
			break exit
		}
	}
}
