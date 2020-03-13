package web

import (
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunServer(config config.Config, logger logrus.Logger) {
	adress := config.Host + ":" + config.Port
	logger.Info("Start web server: ", adress)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/about.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/index.html")
	})
	err := http.ListenAndServe(adress, nil)
	if err != nil {
		logger.Error("Got ", err.Error())
	}
}
