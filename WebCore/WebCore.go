package WebCore

import (
	"encoding/json"
	"net/http"
)

func (Middleware *Middleware) HandlerMiddleware(ControllerFunction ControllerFunction, WebCore *WebCore) http.HandlerFunc {

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
