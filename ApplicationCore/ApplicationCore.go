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
		Error = http.ListenAndServe(fmt.Sprintf(":%d", ApplicationCore.ApplicationPort), ApplicationCore.WebCore.Router)
		if Error != nil {
			return Error
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
			ApplicationCore.WebCore.RabbitMQConnections = append(ApplicationCore.WebCore.RabbitMQConnections, NewRabbitMQSetting)

		case "FileServer":
			if ApplicationCore.WebCore.Router == nil {
				ApplicationCore.WebCore.Router = mux.NewRouter()
				fileServer := http.FileServer(http.Dir("./Static"))
				ApplicationCore.WebCore.Router.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))
			}

		case "Database":
		}

	}
	ApplicationPort, HasApplicationPort := ApplicationCore.ApplicationSettings["ApplicationPort"]
	if HasApplicationPort {
		if ApplicationCore.WebCore.Router == nil {

			ApplicationCore.WebCore.Router = mux.NewRouter()
		}
		ApplicationCore.ApplicationPort = ApplicationPort.(int64)
	} else {

		return errors.New("В файле Settings.json не указан порт приложения")
	}
	return Error
}
