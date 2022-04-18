package ApplicationCore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	Modules "github.com/MrAmperage/GoWebStruct/WebCore/Modules/RabbitMQModule"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
)

func (ApplicationCore *ApplicationCore) Init() (Error error) {
	Error = ApplicationCore.ReadSettings()
	if Error != nil {
		return Error
	}

	return Error
}

func (ApplicationCore *ApplicationCore) Start() (Error error) {

	if ApplicationCore.WebCore.Router == nil {
		return errors.New("Вы не инициализировани настройки приложения")
	} else {
		Error := ApplicationCore.InitRabbitMQ()
		if Error != nil {
			return Error
		}
		Error = http.ListenAndServe(fmt.Sprintf(":%d", ApplicationCore.ApplicationPort), ApplicationCore.WebCore.Router)
		if Error != nil {
			return Error
		}
	}

	return Error
}
func (ApplicationCore *ApplicationCore) InitRabbitMQ() (Error error) {

	ApplicationCore.WebCore.RabbitMQ.RabbitMQChanel.Chanel, Error = ApplicationCore.WebCore.RabbitMQ.Connection.Channel()
	if Error != nil {
		return Error
	}
	Error = ApplicationCore.WebCore.RabbitMQ.QueuesRise()
	if Error != nil {
		return Error
	}
	Error = ApplicationCore.WebCore.RabbitMQ.QueuesSubscribe()
	if Error != nil {
		return Error
	}

	Error = ApplicationCore.WebCore.RabbitMQ.ExchangeRiseAndBind()
	if Error != nil {

		return Error
	}

	return Error
}

func (ApplicationCore *ApplicationCore) ReadSettings() (Error error) {

	_, Error = os.Stat("Settings.json")
	if Error != nil {

		return Error
	}

	ByteSettings, Error := ioutil.ReadFile("Settings.json")
	if Error != nil {
		return Error

	}
	Error = json.Unmarshal(ByteSettings, &ApplicationCore.ApplicationSettings)
	if Error != nil {
		return Error
	}
	ApplicationModulesSettings := ApplicationCore.ApplicationSettings["ApplicationModules"].([]interface{})
	for _, Setting := range ApplicationModulesSettings {
		ModuleType := Setting.(map[string]interface{})["ModuleType"]
		switch ModuleType {

		case "RabbitMQ":
			var NewRabbitMQSetting Modules.RabbitMQ
			mapstructure.Decode(Setting, &NewRabbitMQSetting)
			NewRabbitMQSetting.Connection, Error = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", NewRabbitMQSetting.Login, NewRabbitMQSetting.Password, NewRabbitMQSetting.Adress, NewRabbitMQSetting.Port))
			if Error != nil {
				return Error
			}
			ApplicationCore.WebCore.RabbitMQ = NewRabbitMQSetting

		case "FileServer":
			if ApplicationCore.WebCore.Router == nil {
				ApplicationCore.WebCore.Router = mux.NewRouter()
				FileServer := http.FileServer(http.Dir("./Static/"))
				ApplicationCore.WebCore.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static", FileServer))
			}

		case "Database":
		}

	}
	ApplicationPort, HasApplicationPort := ApplicationCore.ApplicationSettings["ApplicationPort"]
	if HasApplicationPort {
		if ApplicationCore.WebCore.Router == nil {

			ApplicationCore.WebCore.Router = mux.NewRouter()
		}
		ApplicationCore.ApplicationPort = int64(ApplicationPort.(float64))
	} else {

		return errors.New("В файле Settings.json не указан порт приложения")
	}
	return Error
}
