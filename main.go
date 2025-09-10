package main

import (
	"os"
	"sbam/pkg/cmd"
	"sbam/src/utils/mqttclient"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "today"
)

func main() {
	mqttUsername := os.Getenv("MQTT_USERNAME")
	mqttPassword := os.Getenv("MQTT_PASSWORD")
	
	if mqttUsername != "" && mqttPassword != "" {
		mqttclient.InitMQTT(mqttUsername, mqttPassword)
	}
	
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
