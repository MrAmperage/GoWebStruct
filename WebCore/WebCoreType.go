package WebCore

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules"
	"github.com/gorilla/mux"
)

type WebCore struct {
	Router           *mux.Router
	RabbitMQSettings []Modules.RabbitMQSetting
}
