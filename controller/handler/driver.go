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
			return
		}
		id, err := strconv.ParseInt(vars["id"], 10, 32)
		d.Id = int32(id)
		if err != nil {
			writeError(&w, err, errorMessage, http.StatusInternalServerError)
			return
		}

		d.Id, err = driverUsecase.UpdateLocation(d.Id, d.Lat, d.Long, d.Accuracy)

		if err != nil {
			if _, ok := err.(*usecase.IdErr); ok {
				writeError(&w, err, errorMessage, http.StatusNotFound)
				return
			}
			if _, ok := err.(*usecase.LatLngErr); ok {
				writeError(&w, err, errorMessage, http.StatusUnprocessableEntity)
				return
			}

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
}

func FindDrivers(driverUsecase driver.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorTag := "Error finding drivers"
		queryParams := r.URL.Query()

		latitude, err := strconv.ParseFloat(queryParams.Get("latitude"), 64)
		longitude, err := strconv.ParseFloat(queryParams.Get("longitude"), 64)
		limit_str := queryParams.Get("limit")
		radius_str := queryParams.Get("radius")

		var limit int8 = 10
		var radius float64 = 500

		if len(limit_str) > 0 {
			limit_int64, err := strconv.ParseInt(limit_str, 10, 8)
			if err != nil {
				writeError(&w, err, errorTag, http.StatusInternalServerError)
				return
			}
			limit = int8(int(limit_int64))
		}
		if len(radius_str) > 0 {
			radius, err = strconv.ParseFloat(radius_str, 64)
		}

		log.Println("About to find drivers from service & DB")
		drivers, err := driverUsecase.FindDrivers(latitude, longitude, radius, limit)
		log.Println("DB find drivers call finish")

		if err != nil {
			if _, ok := err.(*usecase.LatLngErr); ok {
				writeError(&w, err, errorTag, http.StatusUnprocessableEntity)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(drivers); err != nil {
			writeError(&w, err, errorTag, http.StatusInternalServerError)
			return
		}

	})
}

func MakeDriverHandlers(r *mux.Router, service driver.Usecase) {
	r.Handle(updateDriverPath, UpdateDriver(service)).Methods("PUT").Name("updateDriver")
	r.Handle(findDriversPath, FindDrivers(service)).Methods("GET").Queries("latitude", "{latitude}", "longitude", "{longitude}").Name("findDrivers")

}
