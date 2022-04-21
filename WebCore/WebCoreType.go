package WebCore

import (
	"net/http"

	Modules "github.com/MrAmperage/GoWebStruct/WebCore/Modules/RabbitMQModule"
	"github.com/gorilla/mux"
)

type WebCore struct {
	Router     *mux.Router
	RabbitMQ   Modules.RabbitMQ
	Middleware Middleware
	FileServer FileServerSetting
}

type Middleware struct{}

type FileServerSetting struct {
	StaticDirectory string
}
type ControllerFunction func(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreLink *WebCore) (Data interface{}, Error error)
