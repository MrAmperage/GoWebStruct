package WebCore

import (
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/PostgreSQLModule"
	Modules "github.com/MrAmperage/GoWebStruct/WebCore/Modules/RabbitMQModule"
	"github.com/gorilla/mux"
)

type WebCore struct {
	Router     *mux.Router
	RabbitMQ   Modules.RabbitMQ
	PostgreSQL PostgreSQLModule.PostgreSQLArray
	Middleware Middleware
	FileServer FileServerSetting
}

type Middleware struct{}

type FileServerSetting struct {
	StaticDirectory string
}
type ControllerFunction func(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreLink *WebCore) (Data interface{}, Error error)
