package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gojek-1st/pkg/driver/mocks"
	"gojek-1st/pkg/driver/usecase"
	"gojek-1st/pkg/entity"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDriverHandler(t *testing.T) {
	t.Run("TestUpdateDriver_validId_validLatLng", TestUpdateDriver_validId_validLatLng)
	t.Run("TestUpdateDriver_invalidId_validLatLng", TestUpdateDriver_invalidId_validLatLng)
	t.Run("TestUpdateDriver_validId_invalidLatLng", TestUpdateDriver_validId_invalidLatLng)
	t.Run("TestFindDrivers_validLatlng", TestFindDrivers_validLatlng)
	t.Run("TestFindDrivers_invalidLatLng", TestFindDrivers_invalidLatLng)
}

func CreateTestDriver1() entity.Driver {
	test_driver_1 := entity.Driver{
		Accuracy: 0.7,
		Lat:      51.506752,
		Long:     -0.132912,
		Id:       2,
	}
	return test_driver_1
}

func CreateTestDriver2() entity.Driver {
	return entity.Driver{
		Accuracy: 0.8,
		Lat:      51.508888,
		Long:     -0.125706,
		Id:       3,
	}
}

func TestUpdateDriver_validId_validLatLng(t *testing.T) {

	var dId int32 = 1
	ucase := new(mocks.Usecase)
	r := mux.NewRouter()
	ucase.On("UpdateLocation", dId, mock.Anything, mock.Anything, mock.Anything).Return(dId, nil)
	MakeDriverHandlers(r, ucase)
	path, err := r.GetRoute("updateDriver").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, updateDriverPath, path)
	ts := httptest.NewServer(r)
	body := fmt.Sprintf(`{
	"latitude": 12.971,
	"longitude": 23.1,
	"accuracy": 0.7
}`)

	req, err := http.NewRequest("PUT", ts.URL+"/drivers/"+strconv.Itoa(int(dId))+"/location", strings.NewReader(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	driver := entity.Driver{}
	json.NewDecoder(rr.Body).Decode(&driver)
	assert.Equal(t, dId, driver.Id)
	defer ts.Close()
}

func TestUpdateDriver_invalidId_validLatLng(t *testing.T) {

	var dId int32 = 0
	ucase := new(mocks.Usecase)
	r := mux.NewRouter()
	ucase.On("UpdateLocation", dId, mock.Anything, mock.Anything, mock.Anything).Return(dId, &usecase.IdErr{Msg: "ID out of bound"})
	MakeDriverHandlers(r, ucase)
	path, err := r.GetRoute("updateDriver").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, updateDriverPath, path)
	ts := httptest.NewServer(r)
	body := fmt.Sprintf(`{
	"latitude": 12.971,
	"longitude": 23.1,
	"accuracy": 0.7
}`)

	req, err := http.NewRequest("PUT", ts.URL+"/drivers/"+strconv.Itoa(int(dId))+"/location", strings.NewReader(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	defer ts.Close()

}

func TestUpdateDriver_validId_invalidLatLng(t *testing.T) {
	var dId int32 = 1
	ucase := new(mocks.Usecase)
	r := mux.NewRouter()
	ucase.On("UpdateLocation", dId, 92.971, 13.1, mock.Anything).Return(dId, &usecase.LatLngErr{"Latitude should be between +/- 90"})
	MakeDriverHandlers(r, ucase)
	path, err := r.GetRoute("updateDriver").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, updateDriverPath, path)
	ts := httptest.NewServer(r)
	body := fmt.Sprintf(`{
	"latitude": 92.971,
	"longitude": 13.1,
	"accuracy": 0.7
}`)

	req, err := http.NewRequest("PUT", ts.URL+"/drivers/"+strconv.Itoa(int(dId))+"/location", strings.NewReader(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	defer ts.Close()
}

func TestFindDrivers_validLatlng(t *testing.T) {

	user_lat_str := "51.507351"
	user_long_str := "-0.127758"
	user_lat, _ := strconv.ParseFloat(user_lat_str, 64)
	user_long, _ := strconv.ParseFloat(user_long_str, 64)
	driver1 := CreateTestDriver1()
	driver2 := CreateTestDriver2()
	expectedDrivers := []*entity.Driver{&driver1, &driver2}
	ucase := new(mocks.Usecase)
	ucase.On("FindDrivers", user_lat, user_long, mock.Anything, mock.Anything).Return(expectedDrivers, nil)

	r := mux.NewRouter()
	MakeDriverHandlers(r, ucase)
	path, err := r.GetRoute("findDrivers").GetPathTemplate()
	assert.NoError(t, err)
	assert.Equal(t, findDriversPath, path)
	ts := httptest.NewServer(r)
	defer ts.Close()

	url, err := r.Get("findDrivers").URL("latitude", user_lat_str, "longitude", user_long_str)
	res, err := http.Get(ts.URL + url.String())
	assert.NoError(t, err)
	var actualDrivers []*entity.Driver
	err = json.NewDecoder(res.Body).Decode(&actualDrivers)
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedDrivers, actualDrivers)
}

func TestFindDrivers_invalidLatLng(t *testing.T) {
	user_lat_str := "91.507351"
	user_long_str := "91.127758"
	user_lat, _ := strconv.ParseFloat(user_lat_str, 64)
	user_long, _ := strconv.ParseFloat(user_long_str, 64)
	ucase := new(mocks.Usecase)
	ucase.On("FindDrivers", user_lat, user_long, mock.Anything, mock.Anything).Return(nil, &usecase.LatLngErr{"LatLng err"})

	r := mux.NewRouter()
	MakeDriverHandlers(r, ucase)
	ts := httptest.NewServer(r)
	defer ts.Close()

	url, err := r.Get("findDrivers").URL("latitude", user_lat_str, "longitude", user_long_str)
	assert.NoError(t, err)
	res, err := http.Get(ts.URL + url.String())
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}
