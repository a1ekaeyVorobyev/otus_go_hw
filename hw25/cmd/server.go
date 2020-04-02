package main

import (
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/calendar"
	grpcserver "github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/grpc"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/pkg"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/scheduler"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/web"
	"os/signal"
	"syscall"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/logger"
	"os"
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

	logger, f := logger.GetLogger(conf.Log)
	if f != nil {
		defer f.Close()
	}

	//osSignals := make(chan os.Signal, 1)
	//signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	var st pkg.Storage
	var sh pkg.Scheduler
	switch conf.DB.Type {
	case "Postgres":
		post := storage.Postgres{}
		post.Config = conf.DB
		st = &post
		sh = &post
	default:
		inFile := storage.InFile{}
		st = &inFile
	}
	st.Init()

	done := make(chan bool)
	if (sh!=nil) {
		fmt.Println("run sh")
		p := storage.Postgres{}
		p.Init()
		s := scheduler.Scheduler{&p, &logger, conf.Sheduler, conf.Rmq, done}
		go s.Run()
		defer s.ShutDown()
	}

	cal := calendar.Calendar{Config: conf.DB, Storage: st, Logger: &logger}
	//grpcServer := 	grps.Server{conf,&logger,&cal} get error too few values ?
	grpcServer := grpcserver.Server{}
	grpcServer.Calendar = &cal
	grpcServer.Config = conf.Grps
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
