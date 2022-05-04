package ORMModule

import "github.com/MrAmperage/GoWebStruct/WebCore/Modules/PostgreSQLModule"

func (ORM *ORM) InitORM(PostgreSQLArray *PostgreSQLModule.PostgreSQLArray) {
	ORM.PostgreSQLArray = PostgreSQLArray
}
