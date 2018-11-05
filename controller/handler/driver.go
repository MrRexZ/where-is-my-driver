package handler

import (
	"github.com/gorilla/mux"
	"gojek-1st/pkg/driver/usecase"
	"net/http"
)

func updateDrivers(driverUsecase usecase.DriverUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getDrivers(driverUsecase usecase.DriverUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MakeDriverHandlers(r *mux.Router, service usecase.DriverUsecase) {
	r.Handle("/drivers/{id}/location", updateDrivers(service)).Methods("PUT").Name("updateDriver")
	r.Handle("/drivers", getDrivers(service)).Methods("GET").Name("getDrivers")

}
