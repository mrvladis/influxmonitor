package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	var appConfig appConfiguration
	currentTime := time.Now()
	fmt.Println("Storage Metrics to  InfluxDB injection Tool")
	fmt.Println("Execution stated at [", currentTime, "]")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	configFilePath := dir + "/config/config.json"
	configFile, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println("Cant't open configuration file", configFilePath, ". Failed with the following error:", err)
		panic(err.Error())
	}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&appConfig)
	if err != nil {
		fmt.Println("Cant't decode Application parameters from the configuration file ", configFilePath, ". Failed with the following error:", err)
		panic(err.Error())
	}

	cmd := exec.Command("df", "-h")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	var spaces []storageSpace
	headerSkipped := false

	for scanner.Scan() {
		line := scanner.Text()
		if !headerSkipped {
			headerSkipped = true
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 6 {
			size, _ := parseSize(fields[1])
			used, _ := parseSize(fields[2])
			avail, _ := parseSize(fields[3])
			usePercent, _ := parsePercent(fields[4])

			space := storageSpace{
				Filesystem: fields[0],
				Size:       size,
				Used:       used,
				Avail:      avail,
				UsePerc:    usePercent,
				MountedOn:  fields[5],
			}
			spaces = append(spaces, space)
		}
	}

	// // You can now access the data in the `spaces` slice
	// for _, space := range spaces {
	// 	// Do something with the space data
	// 	println(space.Filesystem, space.Size, space.Used, space.Avail, space.UsePerc, space.MountedOn)
	// }

	//INFLUX DB connection
	// Create a new client using an InfluxDB server base URL and an authentication token
	fmt.Println("Trying to connect to InfluxDB host ", appConfig.InfluxHost, " and port ", appConfig.InfluxPort)
	influxdb2dsn := appConfig.InfluxHttpSchema.ToString() + "://" + appConfig.InfluxHost + ":" + strconv.Itoa(appConfig.InfluxPort)
	client := influxdb2.NewClient(influxdb2dsn, appConfig.InfluxToken)
	// Ensures background processes finishes
	defer client.Close()

	// Use  write client for writes to desired bucket
	writeAPI := client.WriteAPI(appConfig.InfluxOrg, appConfig.InfluxBucket)
	errorsCh := writeAPI.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()
	defer fmt.Println("InfluxDB Updated Successfully")
	defer writeAPI.Flush()
	defer fmt.Println("Flushing any of the data to InfluxDB")
	defer fmt.Println("")

	fmt.Println("InfluxDB connected was successfully")
	processInfluxRequest(writeAPI, spaces, appConfig, currentTime)

}

func parseSize(sizeStr string) (uint64, error) {
	sizeStr = strings.TrimSpace(sizeStr)
	if len(sizeStr) == 0 {
		return 0, nil
	}

	size, err := strconv.ParseUint(sizeStr[:len(sizeStr)-1], 10, 64)
	if err != nil {
		return 0, err
	}

	switch sizeStr[len(sizeStr)-1] {
	case 'G':
		size *= 1024 * 1024 * 1024
	case 'M':
		size *= 1024 * 1024
	case 'K':
		size *= 1024
	}

	return size, nil
}

func parsePercent(sizeStr string) (uint64, error) {
	sizeStr = strings.TrimSpace(sizeStr)
	if len(sizeStr) == 0 {
		return 0, nil
	}

	size, err := strconv.ParseUint(sizeStr[:len(sizeStr)-1], 10, 64)
	if err != nil {
		return 0, err
	}

	return size, nil
}
