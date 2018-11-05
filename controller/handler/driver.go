package handler

import (
	"github.com/gorilla/mux"
	"gojek-1st/pkg/driver"
	"net/http"
)

const (
	updateDriverPath = "/drivers/{id}/location"
	getDriversPath   = "/drivers"
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
	r.Handle(updateDriverPath, UpdateDriver(service)).Methods("PUT").Name("updateDriver")
	r.Handle(getDriversPath, GetDrivers(service)).Methods("GET").Name("getDrivers")

}
