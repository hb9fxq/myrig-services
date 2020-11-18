package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/hb9fxq/myrig-services/globals"
	"github.com/hb9fxq/myrig-services/polling"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var appContext globals.ApplicationContext
	globals.GlobalAppCtx = &appContext

	readApplicationConfig()
	openMqttClient()

	go polling.StartPollingRotor()
	go polling.StartPollingAnt()

	go func() {
		router := CreateRouter()
		log.Fatal(http.ListenAndServe(appContext.ListenString, router))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
}

func openMqttClient() {
	opts := mqtt.NewClientOptions().AddBroker(globals.GlobalAppCtx.MqttBroker).SetClientID(globals.GlobalAppCtx.MqttClientId)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetCleanSession(false)
	globals.GlobalAppCtx.MqttClient = mqtt.NewClient(opts)
	if token := globals.GlobalAppCtx.MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func readApplicationConfig() {
	viper.SetConfigName("ms")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/ms/")
	viper.AddConfigPath("$HOME/.ms")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Panicf("fatal error: %s", err)
	}
	viper.Unmarshal(&globals.GlobalAppCtx)
}
