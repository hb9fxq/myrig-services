package globals

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/hb9fxq/flexlib-go/obj"
	"net/http"
	"time"
)

var GlobalAppCtx *ApplicationContext

type ApiKeys struct {
	User string
	Key  string
}

type ApplicationContext struct {
	ListenString       string
	RotorstatusUrl     string
	LoopRotorstatusUrl string
	RotorStatus        *RotorStatusType
	AntennaStatus      *AntennaStatusType
	LoopRotorStatus    *RotorStatusType
	ApiKeys            []*ApiKeys
	MqttBroker         string
	MqttClientId       string
	MqttClient         mqtt.Client
	SmartSdrPanadapter map[string]obj.Panadapter
	SmartSdrSlice      map[string]obj.Slice
}

type RotorStatusType struct {
	Deg int
}

type AntennaStatusType struct {
	Ant        string
	Pattern    string
	LastUpdate time.Time `json:"LastUpdate"`
}

type Route struct {
	Name     string
	Method   string
	Pattern  string
	IsPublic bool
	Handler  AppHandlerFunc
	AppCtx   *ApplicationContext
}

type AppHandlerFunc func(*Route) http.HandlerFunc
