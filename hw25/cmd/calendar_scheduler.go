package main

import (
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/pkg"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/scheduler"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage"
	"github.com/sirupsen/logrus"
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

	logger, f := logger.GetLogger(conf.Log)
	if f != nil {
		defer f.Close()
	}
	var sh pkg.Scheduler
	switch conf.DB.Type {
	case "Postgres":
		var post *storage.Postgres
		post, err = storage.NewPG(conf.DB, &logger)
		sh = post
	default:
	}
	if err != nil {
		logrus.Error("Error with create storage:", err.Error())
		os.Exit(2)
	}
	if sh != nil {
		s, err := scheduler.NewScheduler(sh, &logger, conf.Sheduler, conf.Rmq)
		if err != nil {
			logrus.Error("Error with create scheduler:", err.Error())
			os.Exit(2)
		}
		go s.Run()
		defer s.ShutDown()
	}
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
