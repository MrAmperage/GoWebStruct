package WebCore

import (
	Modules "github.com/MrAmperage/GoWebStruct/WebCore/Modules/RabbitMQModule"
	"github.com/gorilla/mux"
)

type WebCore struct {
	Router     *mux.Router
	RabbitMQ   Modules.RabbitMQ
	Middleware Middleware
}

type Middleware struct{}
