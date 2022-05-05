package WebCore

import (
	"net/http"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/PostgreSQLModule"
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/RabbitMQModule"
	"github.com/gorilla/mux"
)

type WebCore struct {
	Router     *mux.Router
	RabbitMQ   RabbitMQModule.RabbitMQ
	PostgreSQL PostgreSQLModule.PostgreSQLArray
	Middleware Middleware
	FileServer FileServerSetting
}

type Middleware struct{}

type FileServerSetting struct {
	StaticDirectory string
}
type ControllerFunction func(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreLink *WebCore) (Data interface{}, Error error)
type ResponseData struct {
	Data  interface{} `json:"Data,omitempty"`
	Info  string      `json:"Info,omitempty"`
	Error string      `json:"Error,omitempty"`
}
type AuthenticationResponse struct {
	AuthenticationToken string
}

type AuthenticationRequest struct {
	Username string
	Password string
}
