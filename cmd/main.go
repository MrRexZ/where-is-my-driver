package main

import (
	"encoding/json"
	"github.com/icrowley/fake"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"where-is-my-driver/cmd/app"
	"where-is-my-driver/config"
	"where-is-my-driver/pkg/entity"
)

func main() {
	go app.StartServer()
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
	serverCfg := config.GetConfig().ServerCfg
	driverBody := map[string]float64{"latitude": driver.Lat, "longitude": driver.Long, "accuracy": driver.Accuracy}
	driverJsonMarshalled, _ := json.Marshal(driverBody)
	req, err := http.NewRequest("PUT", serverCfg.HostName+":"+serverCfg.Port+"/drivers/"+strconv.Itoa(int(dId))+"/location", strings.NewReader(string(driverJsonMarshalled)))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var decodedDriver entity.Driver
	json.NewDecoder(resp.Body).Decode(&decodedDriver)
	log.Printf("Driver ID %d updated. ", decodedDriver.Id)
	defer resp.Body.Close()

}
