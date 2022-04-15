package WebCore

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules"
	"github.com/gorilla/mux"
)

type WebCore struct {
	Router     *mux.Router
	RabbitMQ   Modules.RabbitMQ
	Middleware Middleware
}

type Middleware struct{}
