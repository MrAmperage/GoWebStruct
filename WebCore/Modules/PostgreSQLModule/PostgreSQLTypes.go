package PostgreSQLModule

import "gorm.io/gorm"

type PostgreSQL struct {
	ConnectionPool *gorm.DB
	ConnectionName string
	Adress         string
	Port           int64
	Login          string
	Password       string
	DatabaseName   string
}
