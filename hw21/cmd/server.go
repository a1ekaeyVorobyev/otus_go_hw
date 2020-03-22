package main

import (
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/storage"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/grps"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/web"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	fmt.Println("Start")
	var cFile string
	flag.StringVar(&cFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	fmt.Println(cFile)
	if cFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "Not set config file")
		os.Exit(2)
	}

	conf, err := config.ReadFromFile(cFile)
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
	fmt.Println(conf)

	cal := calendar.Calendar{Config: conf, Storage: &inFile, Logger: &logger}
	grpcServer := grpc.Server{Config: conf, Logger: &logger, Calendar: &cal}

	go web.RunServer(conf, &logger)
	go grpcServer.Run()

	logger.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	defer grpcServer.Shutdown()

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

exit:
	for {
		select {
		case c:=<-sigs:
			logger.Infof("Got signal: %v. Exit.", c)
			break exit
		}
	}
}
