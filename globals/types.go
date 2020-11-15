package globals

import "net/http"

type ApiKeys struct {
	User string
	Key  string
}

type ApplicationContext struct {
	ListenString   string
	RotorstatusUrl string
	RotorStatus    *RotorStatusType
	AntennaStatus  *AntennaStatusType
	ApiKeys        []*ApiKeys
}

type RotorStatusType struct {
	Deg int
}

type AntennaStatusType struct {
	Ant     string
	Pattern string
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
