package polling

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/hb9fxq/myrig-services/globals"
	"os"
	"strings"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	mRaw := fmt.Sprintf("%s", msg.Payload())
	tokens := strings.Split(mRaw, " ")

	globals.GlobalAppCtx.AntennaStatus.Ant = strings.TrimSuffix(tokens[0], "\r\n")
	globals.GlobalAppCtx.AntennaStatus.Pattern = strings.TrimSuffix(strings.Split(tokens[1], ">")[1], "\r\n")
	globals.GlobalAppCtx.AntennaStatus.LastUpdate = time.Now()
	fmt.Print(time.Now())
	fmt.Println("    >>>>" + mRaw)
}

func StartPollingAnt() {
	var antennaStatus globals.AntennaStatusType
	globals.GlobalAppCtx.AntennaStatus = &antennaStatus
	globals.GlobalAppCtx.AntennaStatus.Ant = "U"
	globals.GlobalAppCtx.AntennaStatus.Pattern = "STATE>????????????"

	if token := globals.GlobalAppCtx.MqttClient.Subscribe("ant/silent", 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		token := globals.GlobalAppCtx.MqttClient.Publish("ant/cmd", 0, false, "RS")
		token.Wait()
		<-time.After(1 * time.Second)
	}
}
