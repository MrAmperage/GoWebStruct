package ORMModule

import "github.com/MrAmperage/GoWebStruct/WebCore/Modules/PostgreSQLModule"

type ORM struct {
	PostgreSQLArray *PostgreSQLModule.PostgreSQLArray
}

type ORMInterface interface {
	InitORM(PostgreSQLArray *PostgreSQLModule.PostgreSQLArray)
}
