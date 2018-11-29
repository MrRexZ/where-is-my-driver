package app

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	"where-is-my-driver/config"
	"where-is-my-driver/controller/handler"
	"where-is-my-driver/pkg/driver/repository"
	"where-is-my-driver/pkg/driver/usecase"
)

func StartServer() {
	mongoCfg := config.GetConfig().MongoCfg
	readClient, err := repository.CreateMongoClient(mongoCfg.HostName)

	if err != nil {
		log.Fatal(err.Error())
	}

	writeClient, err := repository.CreateMongoClient(mongoCfg.HostName)

	if err != nil {
		log.Fatal(err.Error())
	}
	mongoRepo := repository.CreateMongoRepository(readClient, writeClient, mongoCfg.DbName)
	driverUCase := usecase.NewDriverUsecase(mongoRepo)
	r := mux.NewRouter()
	handler.MakeDriverHandlers(r, driverUCase)
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + config.GetConfig().ServerCfg.Port,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
