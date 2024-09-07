package main

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func processInfluxRequest(writeAPI api.WriteAPI, spaces []storageSpace, appConfig appConfiguration, currentTime time.Time) {
	var i int
	//var err error

	fmt.Println("Writing data to InfluxDB")
	for i = 0; i < len(spaces); i++ {
		fmt.Printf("\rProcessing Row %d/%d \n", i, len(spaces))

		p := influxdb2.NewPointWithMeasurement("filesystem").
			AddTag("host", appConfig.TrueNasHostName).
			AddTag("os", appConfig.TrueNasOS).
			AddTag("type", spaces[i].Filesystem).
			//			AddField("name", spaces[i].Filesystem).
			AddField("mountedOn", spaces[i].MountedOn).
			AddField("size", spaces[i].Size).
			SetTime(currentTime)
		writeAPI.WritePoint(p)

		p = influxdb2.NewPointWithMeasurement("filesystem").
			AddTag("host", appConfig.TrueNasHostName).
			AddTag("os", appConfig.TrueNasOS).
			AddTag("type", spaces[i].Filesystem).
			//			AddField("name", spaces[i].Filesystem).
			AddField("mountedOn", spaces[i].MountedOn).
			AddField("used", spaces[i].Used).
			SetTime(currentTime)
		writeAPI.WritePoint(p)

		p = influxdb2.NewPointWithMeasurement("filesystem").
			AddTag("host", appConfig.TrueNasHostName).
			AddTag("os", appConfig.TrueNasOS).
			AddTag("type", spaces[i].Filesystem).
			//			AddField("name", spaces[i].Filesystem).
			AddField("mountedOn", spaces[i].MountedOn).
			AddField("avail", spaces[i].Avail).
			SetTime(currentTime)
		writeAPI.WritePoint(p)

		p = influxdb2.NewPointWithMeasurement("filesystem").
			AddTag("host", appConfig.TrueNasHostName).
			AddTag("os", appConfig.TrueNasOS).
			AddTag("type", spaces[i].Filesystem).
			//			AddField("name", spaces[i].Filesystem).
			AddField("mountedOn", spaces[i].MountedOn).
			AddField("use", spaces[i].UsePerc).
			SetTime(currentTime)
		writeAPI.WritePoint(p)
	}
}
