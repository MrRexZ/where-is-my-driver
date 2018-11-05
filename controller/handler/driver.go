package handler

import (
	"github.com/gorilla/mux"
	"gojek-1st/pkg/driver/usecase"
	"net/http"
)

func UpdateDriver(driverUsecase usecase.DriverUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetDrivers(driverUsecase usecase.DriverUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MakeDriverHandlers(r *mux.Router, service usecase.DriverUsecase) {
	r.Handle("/drivers/{id}/location", UpdateDriver(service)).Methods("PUT").Name("updateDriver")
	r.Handle("/drivers", GetDrivers(service)).Methods("GET").Name("getDrivers")

}
