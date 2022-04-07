package ApplicationCore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules"
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
	for _, RabbitMQSetting := range ApplicationCore.WebCore.RabbitMQSettings {
		RabbitMQSetting.RabbitMQChanel.Chanel, Error = RabbitMQSetting.Connection.Channel()
		if Error != nil {
			return Error
		}
		for _, QueueUP := range RabbitMQSetting.RabbitMQChanel.QueuesUP {

			QueueUP.Queue, Error = RabbitMQSetting.RabbitMQChanel.Chanel.QueueDeclare(QueueUP.Name, false,
				false,
				false,
				false,
				nil)
			if Error != nil {
				return Error
			}

		}
		for _, RabbitMQSubscribe := range RabbitMQSetting.RabbitMQChanel.Subscribes {
			RabbitMQSubscribe.Messages, Error = RabbitMQSetting.RabbitMQChanel.Chanel.Consume(RabbitMQSubscribe.Name, "", true, false, false, false, nil)
			if Error != nil {
				return Error
			}

		}

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
			var NewRabbitMQSetting Modules.RabbitMQSetting
			mapstructure.Decode(Setting, &NewRabbitMQSetting)
			NewRabbitMQSetting.Connection, Error = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", NewRabbitMQSetting.Login, NewRabbitMQSetting.Password, NewRabbitMQSetting.Adress, NewRabbitMQSetting.Port))
			if Error != nil {
				return Error
			}
			ApplicationCore.WebCore.RabbitMQSettings = append(ApplicationCore.WebCore.RabbitMQSettings, NewRabbitMQSetting)

		case "FileServer":
			if ApplicationCore.WebCore.Router == nil {
				ApplicationCore.WebCore.Router = mux.NewRouter()
				FileServer := http.FileServer(http.Dir("./Static"))
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
