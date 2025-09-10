package mqttclient

import (
    "fmt"
    u "sbam/src/utils"
    mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Client mqtt.Client

func InitMQTT(MQTT_USERNAME, MQTT_PASSWORD string) {
    opts := mqtt.NewClientOptions().
        AddBroker("tcp://core-mosquitto:1883").
        SetClientID("sbam-addon").
        SetUsername(MQTT_USERNAME).
        SetPassword(MQTT_PASSWORD)

    Client = mqtt.NewClient(opts)
    if token := Client.Connect(); token.Wait() && token.Error() != nil {
        u.Log.Errorf("MQTT connection failed: %s\n", token.Error())
    } else {
        u.Log.Info("Connected to MQTT broker")
    }
}

func Publish(topic string, payload string) {
    if Client != nil && Client.IsConnected() {
        token := Client.Publish(topic, 0, false, payload)
        if token.Wait() && token.Error() != nil {
            u.Log.Errorf("MQTT publish failed: %s", token.Error())
        }
    } else {
        u.Log.Errorf("MQTT client not connected")
    }
}

func PublishNumber(topic string, value int) {
    Publish(topic, fmt.Sprintf("%d", value))
}
