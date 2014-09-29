package main

// Reference : http://tdoc.info/blog/2014/09/25/mqtt_golang.html

import (
    "fmt"
    "log"
    "time"

    MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func onMessageReceived(client *MQTT.MqttClient, message MQTT.Message) {
    fmt.Printf("Received message on topic: %s\n", message.Topic())
    fmt.Printf("Message: %s\n", message.Payload())
}

func Subscribe(client *MQTT.MqttClient) error {
    topic := "mqtt/a/b/c"
    qos := 0

    topicFilter, err := MQTT.NewTopicFilter(topic, byte(qos))
    if err != nil {
        return err
    }

    _, err = client.StartSubscription(onMessageReceived, topicFilter)
    if err != nil {
        return err
    }

    for {
        time.Sleep(1 * time.Second)
    }
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

    err = Subscribe(client)
    if err != nil {
        log.Fatal(err)
    }
}
