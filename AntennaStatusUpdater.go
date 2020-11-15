package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/krippendorf/myrig-services/globals"
	"os"
	"strings"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	mRaw := fmt.Sprintf("%s", msg.Payload())
	tokens := strings.Split(mRaw, " ")

	appCtx.AntennaStatus.Ant = strings.TrimSuffix(tokens[0], "\r\n")
	appCtx.AntennaStatus.Pattern = strings.TrimSuffix(strings.Split(tokens[1], ">")[1], "\r\n")
}

func startPollingAnt() {
	var antennaStatus globals.AntennaStatusType
	appCtx.AntennaStatus = &antennaStatus
	appCtx.AntennaStatus.Ant = "U"
	appCtx.AntennaStatus.Pattern = "STATE>????????????"

	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.92.7:1883").SetClientID("myrig-services") //TODO: cfg
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("ant/silent", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		token := c.Publish("ant/cmd", 0, false, "RS")
		token.Wait()
		<-time.After(2 * time.Second)
	}
}
