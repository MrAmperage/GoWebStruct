package ORMModule

import (
	"gorm.io/gorm"
)

type ORMInterface interface {
	GetName() (Name string)
	SetConnection(ConnectionLink *gorm.DB)
}

type ORMArray struct {
	ORMElements []ORMInterface
}
