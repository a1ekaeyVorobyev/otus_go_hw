package web

import (
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw15/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunServer(config config.Config, logger *logrus.Logger) {
	adress := config.Host + ":" + config.Port
	logger.Info("Start web server: ", adress)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/about.html")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/hello.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "website/index.html")
	})
	err := http.ListenAndServe(adress, logRequest(http.DefaultServeMux, logger))
	if err != nil {
		logger.Error("Got ", err.Error())
	}
}

func logRequest(handler http.Handler, logger *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}
