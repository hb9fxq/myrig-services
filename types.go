package main

type ApplicationContext struct {
	ListenString   string
	RotorstatusUrl string
	RotorStatus    *RotorStatusType
}

type RotorStatusType struct {
	Deg int
}
