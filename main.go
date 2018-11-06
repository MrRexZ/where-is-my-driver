package main

import (
	"encoding/json"
	"github.com/icrowley/fake"
	"gojek-1st/cmd"
	"gojek-1st/config"
	"gojek-1st/pkg/entity"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	go cmd.StartServer()
	generate50000Drivers()
	c := time.Tick(60 * time.Second)
	for range c {
		generate50000Drivers()
	}

}

func generate50000Drivers() {
	client := http.DefaultClient
	for i := 1; i <= 50000; i++ {
		driver := entity.Driver{Id: int32(i), Lat: float64(fake.Latitude()), Long: float64(fake.Longitude()), Accuracy: 0.8}
		updateDriver(client, driver)
	}
}

func updateDriver(client *http.Client, driver entity.Driver) {
	dId := driver.Id
	driverBody := map[string]float64{"latitude": driver.Lat, "longitude": driver.Long, "accuracy": driver.Accuracy}
	driverJsonMarshalled, _ := json.Marshal(driverBody)
	req, err := http.NewRequest("PUT", config.HOST+":"+strconv.Itoa(config.REST_API_PORT)+"/drivers/"+strconv.Itoa(int(dId))+"/location", strings.NewReader(string(driverJsonMarshalled)))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var decodedDriver entity.Driver
	json.NewDecoder(resp.Body).Decode(&decodedDriver)
	defer resp.Body.Close()

}
