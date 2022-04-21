package WebCore

import (
	"encoding/json"
	"net/http"
)

type ControllerFunction func(ResponseWriter http.ResponseWriter, Request *http.Request, WebCoreLink *WebCore) (Data interface{}, Error error)

func (Middleware *Middleware) HandlerMiddleware(ControllerFunction ControllerFunction, WebCore *WebCore) http.HandlerFunc {
	type ResponseData struct {
		Data  interface{} `json:"Data,omitempty"`
		Info  string      `json:"Info,omitempty"`
		Error string      `json:"Error,omitempty"`
	}
	return func(ResponseWriter http.ResponseWriter, Request *http.Request) {
		Response := &ResponseData{}
		Data, Error := ControllerFunction(ResponseWriter, Request, WebCore)
		if Error != nil {
			Response.Error = Error.Error()
		} else {
			switch Data.(type) {
			case string:
				Response.Info = Data.(string)
			default:
				Response.Data = Data
			}
		}
		ResponseByte, Error := json.Marshal(Response)
		if Error != nil {

			ResponseWriter.Write([]byte(`{"Error": "` + Error.Error() + `"}`))
		} else {
			ResponseWriter.Write(ResponseByte)
		}
	}
}
