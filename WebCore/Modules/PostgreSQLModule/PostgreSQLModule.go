package PostgreSQLModule

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (PostgreSQL *PostgreSQL) InitPostgreSQL() (Error error) {
	PostgreSQL.ConnectionPool, Error = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", PostgreSQL.Adress, PostgreSQL.Login, PostgreSQL.Password, PostgreSQL.DatabaseName, PostgreSQL.Port)))
	if Error != nil {
		return
	}
	return
}
