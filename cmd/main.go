package cmd

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
	"strconv"
	"time"
)

func main() {
	mongoClient, err := repository.CreateMongoClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	mongoRepo := repository.CreateMongoRepository(mongoClient, config.MONGODB_DB_NAME)
	driverUCase := usecase.NewDriverUsecase(mongoRepo)
	r := mux.NewRouter()
	handler.MakeDriverHandlers(r, driverUCase)
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.REST_API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
