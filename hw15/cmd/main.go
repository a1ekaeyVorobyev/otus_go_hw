package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/storage"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/web"
	_ "github.com/a1ekaeyVorobyev/otus_go_hw/hw15/web"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan bool, 1)
	sigs := make(chan os.Signal, 1)

	config, err := config.ReadFromFile("config/config.yaml")
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	logger, f := logger.GetLogger(config)
	if f != nil {
		defer f.Close()
	}
	inMemory := storage.InFile{}
	inMemory.Init()
	defer func() {
		err := inMemory.SaveEvents()
		logger.Error(err)
	}()

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	_ = calendar.Calendar{Config: config, Storage: &inMemory, Logger: &logger}

	go web.RunServer(config, &logger)

exit:
	for {
		select {
		case <-done:
			logger.Info("Exit service.")
			break exit
		}
	}

	logger.Exit(0)
}
