package main

import (
	"github.com/krippendorf/myrig-services/globals"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func startPollingRotor() {

	var rotorStatus globals.RotorStatusType
	appCtx.RotorStatus = &rotorStatus
	appCtx.RotorStatus.Deg = 1000

	for {
		updateRotorStatus()
		<-time.After(10 * time.Second)
		updateRotorStatus()
	}
}

func updateRotorStatus() {
	getResult := getHttpString(appCtx.RotorstatusUrl)
	appCtx.RotorStatus.Deg = 1000
	if len(getResult) == 0 {
		log.Printf("/rotatorcontrol/get operation failed\n")
	}

	tokens := strings.Split(getResult, "|")
	stateInt, err := strconv.Atoi(tokens[3])

	if err != nil {
		log.Printf("HTTP GET ERR: %s\n", err)
	} else {
		appCtx.RotorStatus.Deg = stateInt
	}
}

func getHttpString(url string) (responseString string) {

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("HTTP GET ERR: %s\n", err)
	}

	if resp.StatusCode == 200 {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)

		if err2 != nil {
			log.Printf("HTTP GET ERR: %s\n", err2)
		}

		responseString = string(bodyBytes)
	}
	return
}
