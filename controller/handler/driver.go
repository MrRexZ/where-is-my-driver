package handler

import (
	"github.com/gorilla/mux"
	"gojek-1st/pkg/driver"
	"net/http"
)

func UpdateDriver(driverUsecase driver.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GetDrivers(driverUsecase driver.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MakeDriverHandlers(r *mux.Router, service driver.Usecase) {
	r.Handle("/drivers/{id}/location", UpdateDriver(service)).Methods("PUT").Name("updateDriver")
	r.Handle("/drivers", GetDrivers(service)).Methods("GET").Name("getDrivers")

}
