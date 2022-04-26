package ApplicationCore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/PostgreSQLModule"
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

func (ApplicationCore *ApplicationCore) StartWebServer() (Error error) {

	if ApplicationCore.WebCore.Router == nil {
		return errors.New("вы не инициализировани настройки приложения")
	} else {
		Error = http.ListenAndServe(fmt.Sprintf(":%d", ApplicationCore.ApplicationPort), ApplicationCore.WebCore.Router)
		if Error != nil {
			return Error
		}
	}

	return Error
}
func (ApplicationCore *ApplicationCore) StartDatabaseConnections() (Error error) {
	for Index, _ := range ApplicationCore.WebCore.PostgreSQL {
		Error = ApplicationCore.WebCore.PostgreSQL[Index].InitPostgreSQL()
		if Error != nil {
			return
		}

	}

	return
}
func (ApplicationCore *ApplicationCore) StartRabbitMQ() (Error error) {

	ApplicationCore.WebCore.RabbitMQ.RabbitMQChanel.Chanel, Error = ApplicationCore.WebCore.RabbitMQ.Connection.Channel()
	if Error != nil {
		return Error
	}
	Error = ApplicationCore.WebCore.RabbitMQ.ExchangeRiseAndBind()
	if Error != nil {

		return Error
	}
	Error = ApplicationCore.WebCore.RabbitMQ.QueuesRiseAndBind()
	if Error != nil {
		return Error
	}
	Error = ApplicationCore.WebCore.RabbitMQ.QueuesSubscribe()
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
			var NewFileServerSettings WebCore.FileServerSetting
			if ApplicationCore.WebCore.Router == nil {
				ApplicationCore.WebCore.Router = mux.NewRouter()
				mapstructure.Decode(Setting, &NewFileServerSettings)
				ApplicationCore.WebCore.FileServer.StaticDirectory = NewFileServerSettings.StaticDirectory
				FileServer := http.FileServer(http.Dir(fmt.Sprintf("./%s/", NewFileServerSettings.StaticDirectory)))
				ApplicationCore.WebCore.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static", FileServer))
			}

		case "PostgreSQLDatabase":
			var NewPostgreSQLSetting PostgreSQLModule.PostgreSQL
			mapstructure.Decode(Setting, &NewPostgreSQLSetting)
			ApplicationCore.WebCore.PostgreSQL = append(ApplicationCore.WebCore.PostgreSQL, NewPostgreSQLSetting)

		}

	}
	ApplicationPort, HasApplicationPort := ApplicationCore.ApplicationSettings["ApplicationPort"]
	if HasApplicationPort {
		if ApplicationCore.WebCore.Router == nil {

			ApplicationCore.WebCore.Router = mux.NewRouter()
		}
		ApplicationCore.ApplicationPort = int64(ApplicationPort.(float64))
	} else {

		return errors.New("в файле Settings.json не указан порт приложения")
	}
	return Error
}
