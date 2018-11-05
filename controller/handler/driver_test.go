package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gojek-1st/pkg/driver/mocks"
	"gojek-1st/pkg/entity"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateDriver(t *testing.T) {
	ucase := new(mocks.Usecase)
	r := mux.NewRouter()
	MakeDriverHandlers(r, ucase)
	path, err := r.GetRoute("updateDriver").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, updateDriverPath, path)
	handler := UpdateDriver(ucase)
	ts := httptest.NewServer(handler)
	var dId int32 = 1
	body := fmt.Sprintf(`{
	"latitude": 12.971,
	"longitude": 23.1
	"accuracy": 0.7
}`)

	req, err := http.NewRequest("PUT", ts.URL+"/drivers/"+string(dId)+"/location", strings.NewReader(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	driver := entity.Driver{}
	json.NewDecoder(rr.Body).Decode(&driver)
	assert.Equal(t, dId, driver.Id)
	defer ts.Close()
}

func TestFindDrivers(t *testing.T) {

}
