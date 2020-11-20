package polling

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/hb9fxq/flexlib-go/obj"
	"github.com/hb9fxq/myrig-services/globals"
	"os"
)

var updatePanadapters mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	globals.GlobalAppCtx.SmartSdrPanadapter = make(map[string]obj.Panadapter)
	err := json.Unmarshal(msg.Payload(), &globals.GlobalAppCtx.SmartSdrPanadapter)

	if err != nil {
		fmt.Println(err)
	}
}

var updateSlices mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	globals.GlobalAppCtx.SmartSdrSlice = make(map[string]obj.Slice)
	err := json.Unmarshal(msg.Payload(), &globals.GlobalAppCtx.SmartSdrSlice)

	if err != nil {
		fmt.Println(err)
	}

}

func StartPollingSmartSdr() {

	if token := globals.GlobalAppCtx.MqttClient.Subscribe("smartsdr/slices", 0, updateSlices); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	if token := globals.GlobalAppCtx.MqttClient.Subscribe("smartsdr/panadapters", 0, updatePanadapters); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

}
