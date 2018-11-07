package app

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gojek-1st/config"
	"gojek-1st/controller/handler"
	"gojek-1st/pkg/driver/repository"
	"gojek-1st/pkg/driver/usecase"
	"log"
	"net/http"
	"os"
	"time"
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
