package configuration

import (
	"net/http"
)

type Database struct {
	Uri      string
	Username string
	Password string
}

type Api struct {
	Ip   string
	Port uint16
}

type IContainer interface {
	InitDependency(interface{}) error
	InitVariable() error
	DefineDatabase(databaseWrappers ...any) error
	DefineRoute(router interface{}) error
	DefineGrpc() error
	SetRouter(router any)

	GetHttpHandler() http.Handler
}
