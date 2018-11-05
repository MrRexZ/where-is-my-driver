package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gojek-1st/pkg/driver"
	"gojek-1st/pkg/driver/usecase"
	"gojek-1st/pkg/entity"
	"log"
	"net/http"
	"strconv"
)

const (
	updateDriverPath = "/drivers/{id}/location"
	findDriversPath  = "/drivers"
)

func UpdateDriver(driverUsecase driver.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating driver"
		vars := mux.Vars(r)
		var d *entity.Driver

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			writeError(&w, err, errorMessage, http.StatusInternalServerError)
		}
		id, err := strconv.ParseInt(vars["id"], 10, 32)
		d.Id = int32(id)
		if err != nil {
			writeError(&w, err, errorMessage, http.StatusInternalServerError)
		}

		d.Id, err = driverUsecase.UpdateLocation(d.Id, d.Lat, d.Long, d.Accuracy)

		if err != nil {
			if _, ok := err.(*usecase.IdErr); ok {
				writeError(&w, err, errorMessage, http.StatusNotFound)
			}
			if _, ok := err.(*usecase.LatLngErr); ok {
				writeError(&w, err, errorMessage, http.StatusUnprocessableEntity)
			}
			return

		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(d); err != nil {
			writeError(&w, err, errorMessage, http.StatusInternalServerError)
			return
		}

	})
}

func writeError(w *http.ResponseWriter, error error, errorTag string, status int) {
	writer := *w
	log.Println(error.Error())
	writer.WriteHeader(status)
	writer.Write([]byte(errorTag))
	return
}

func FindDrivers(driverUsecase driver.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func MakeDriverHandlers(r *mux.Router, service driver.Usecase) {
	r.Handle(updateDriverPath, UpdateDriver(service)).Methods("PUT").Name("updateDriver")
	r.Handle(findDriversPath, FindDrivers(service)).Methods("GET").Queries("latitude", "{latitude}", "longitude", "{longitude}").Name("findDrivers")

}
