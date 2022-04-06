package WebCore

import "github.com/MrAmperage/GoWebStruct/DatabaseCore"

type WebCore struct {
	DatabaseCore        DatabaseCore.DatabaseCore
	RabbitMQConnections []RabbitMQSetting
}

type RabbitMQSetting struct {
	ModuleType string
	Login      string
	Password   string
	Port       int64
	Adress     string
}
