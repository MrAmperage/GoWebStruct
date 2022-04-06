package ApplicationCore

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/mitchellh/mapstructure"
)

func (ApplicationCore *ApplicationCore) Start() (Error error) {
	Error = ApplicationCore.ReadSettings()
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
			var NewRabbitMQSetting WebCore.RabbitMQSetting
			mapstructure.Decode(Setting, &NewRabbitMQSetting)
			ApplicationCore.WebCore.RabbitMQConnections = append(ApplicationCore.WebCore.RabbitMQConnections, NewRabbitMQSetting)

		case "FileServer":

		case "Database":
		}

	}

	return Error
}
