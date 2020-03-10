package main

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/storage"
	"os"

	// "github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/config"
	//"github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/logger"
	// "github.com/a1ekaeyVorobyev/otus_go_hw/hw13/internal/storage"
)

func main() {

	config, err := config.ReadFromFile("config/config.yaml")
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	logger,f := logger.GetLogger(config)
	if f != nil{
		defer f.Close()
	}
	logger.Debug("Test")
	inMemory := storage.InFile{}
	inMemory.Init()
	defer func(){
		err := inMemory.SaveEvents()
		logger.Error(err)
	}()
	//_ = calendar.Calendar{Config: config, Storage: &inMemory, Logger: logger}
	fmt.Println("Hello, playground")
	fmt.Println(config)
}

