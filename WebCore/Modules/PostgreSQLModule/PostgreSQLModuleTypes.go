package PostgreSQLModule

import (
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	ConnectionPool *gorm.DB
	ConnectionName string
	Adress         string
	Port           int64
	Login          string
	Password       string
	DatabaseName   string
	ORMs           ORMModule.ORMArray
}
type PostgreSQLArray struct {
	Elements []PostgreSQL
}
