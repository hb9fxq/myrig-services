package polling

import (
	"github.com/hb9fxq/myrig-services/globals"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func StartPollingRotor() {

	globals.GlobalAppCtx.RotorStatus = &globals.RotorStatusType{}
	globals.GlobalAppCtx.LoopRotorStatus = &globals.RotorStatusType{}

	for {
		updateRotorStatus(globals.GlobalAppCtx.RotorstatusUrl, globals.GlobalAppCtx.RotorStatus)
		updateRotorStatus(globals.GlobalAppCtx.LoopRotorstatusUrl, globals.GlobalAppCtx.LoopRotorStatus)
		<-time.After(10 * time.Second)
	}
}

func updateRotorStatus(url string, statusType *globals.RotorStatusType) {
	getResult := getHttpString(url)
	statusType.Deg = 1000

	if len(getResult) == 0 {
		log.Printf("/rotatorcontrol/get operation failed\n")
		return
	}

	tokens := strings.Split(getResult, "|")
	stateInt, err := strconv.Atoi(tokens[3])

	if err != nil {
		log.Printf("HTTP GET ERR: %s\n", err)
	} else {
		statusType.Deg = stateInt
	}

}

func getHttpString(url string) (responseString string) {

	responseString = ""
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("HTTP GET ERR: %s\n", err)
		return
	}

	if resp != nil && resp.StatusCode == 200 {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)

		if err2 != nil {
			log.Printf("HTTP GET ERR: %s\n", err2)
		}

		responseString = string(bodyBytes)
	}
	return
}
