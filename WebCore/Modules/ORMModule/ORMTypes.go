package ORMModule

import (
	"gorm.io/gorm"
)

type ORMInterface interface {
	GetName() (Name string)
	SetConnection(ConnectionLink *gorm.DB)
	SetName(Name string)
}

type ORMArray struct {
	ORMElements []ORMInterface
}

type ORM struct {
	Name           string
	ConnectionLink *gorm.DB
}
