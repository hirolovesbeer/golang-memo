package main

// Reference : http://tdoc.info/blog/2014/09/25/mqtt_golang.html

import (
    "fmt"
    "log"

    MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func Publish(client *MQTT.MqttClient, cnt int) error {
    topic := "mqtt/a/b/c"
    qos := 0
    message := fmt.Sprintf("MQTT from golang: %d", cnt)

    // https://godoc.org/git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git#MqttClient.Publish
    // return the Receipt chan
    result := client.Publish(MQTT.QoS(qos), topic, message)
    <- result

    return nil
}

func main() {
    host := "targethostname"
    port := 1883

    // read http://www.rabbitmq.com/mqtt.html : Authentication
    user := "/mqtt:mqtt"
    password := "hogehoge"

    opts := MQTT.NewClientOptions()
    opts.SetUsername(user)
    opts.SetPassword(password)

    brokerUri := fmt.Sprintf("tcp://%s:%d", host, port)
    opts.AddBroker(brokerUri)

    client := MQTT.NewClient(opts)

    _, err := client.Start()
    if err != nil {
        log.Fatal(err)
    }

    for i := 0; i < 1000000; i++ {
        err = Publish(client, i)
        if err != nil {
            log.Fatal(err)
        }
    }
}
