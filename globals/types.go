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
	ApiKeys        []*ApiKeys
}

type RotorStatusType struct {
	Deg int
}

type Route struct {
	Name     string
	Method   string
	Pattern  string
	IsPublic bool
	Handler  AppHandlerFunc
	AppCtx   *ApplicationContext
}

type AppHandlerFunc func(Route) http.HandlerFunc
