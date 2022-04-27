package PostgreSQLModule

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (PostgreSQLArray *PostgreSQLArray) Add(PostgreSQL PostgreSQL) {
	PostgreSQLArray.Elements = append(PostgreSQLArray.Elements, PostgreSQL)

}
func (PostgreSQLArray *PostgreSQLArray) FindByName(Name string) (PostgreSQL PostgreSQL, Error error) {
	for Index, Postgre := range PostgreSQLArray.Elements {
		if Postgre.ConnectionName == Name {
			return PostgreSQLArray.Elements[Index], Error

		}
	}
	return PostgreSQL, errors.New("postgre подключение не найдено")
}

func (PostgreSQLArray *PostgreSQLArray) StartDatabaseConnections() (Error error) {
	for Index, PostgreSQL := range PostgreSQLArray.Elements {
		PostgreSQLArray.Elements[Index].ConnectionPool, Error = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", PostgreSQL.Adress, PostgreSQL.Login, PostgreSQL.Password, PostgreSQL.DatabaseName, PostgreSQL.Port)))
		if Error != nil {
			return
		}

	}

	return
}
