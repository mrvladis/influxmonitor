package main

import (
	"time"
)

// Define a custom type
type httpSchema string

// Define constants of the custom type
const (
	Value1 httpSchema = "http"
	Value2 httpSchema = "https"
)

func (m httpSchema) ToString() string {
	return string(m)
}

type appConfiguration struct {
	InfluxHost       string     `json:"Influx_Host,omitempty"`
	InfluxHttpSchema httpSchema `json:"Influx_HttpSchema,omitempty"`
	InfluxPort       int        `json:"Influx_Port,omitempty"`
	InfluxToken      string     `json:"Influx_Token,omitempty"`
	InfluxBucket     string     `json:"Influx_Bucket,omitempty"`
	InfluxOrg        string     `json:"Influx_Org,omitempty"`
	TrueNasHostName  string     `json:"TrueNas_HostName,omitempty"`
	TrueNasOS        string     `json:"TrueNas_OS,omitempty"`
	TrueNasCategory  string     `json:"TrueNas_Category,omitempty"`
}

type storageSpace struct {
	TimeStamp  time.Time
	Filesystem string
	Size       uint64
	Used       uint64
	Avail      uint64
	UsePerc    uint64
	MountedOn  string
}
