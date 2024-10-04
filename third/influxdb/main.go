package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	bucket = "aaaa"
	org    = "root"
	token  = "34xdpYgz-xoSCqhQ5-Sy9SWrlZRBG-E8C_rmXP9f8YsayzFCvibKMGDrxVYE24-N2LizCrlttiwN6-HoyhK76w=="
	// Store the URL of your InfluxDB instance
	url = "http://172.16.56.129:8086"
)

// connect
func connInflux() influxdb2.Client {
	client := influxdb2.NewClient(url, token)
	defer client.Close()
	return client
}

// query
/*func queryDB(cli influxdb2.Client, cmd string) (res []influxdb2.)  {
	result, err := cli.QueryAPI(org).QueryRaw(context.Background(), cmd, influxdb2.DefaultDialect())
	if err != nil {
		fmt.Printf("query db fail, err:%v\n", err)
		return
	}
}*/

func main() {
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, bucket)

	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("状态", "正常").
		AddField("max", 35.0)

	writeAPI.WritePoint(p)
	writeAPI.Flush()

	defer client.Close()
}
