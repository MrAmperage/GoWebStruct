package ORMModule

import (
	"errors"

	"gorm.io/gorm"
)

func (ORMArray *ORMArray) Add(ORM ORMInterface) {

	ORMArray.ORMElements = append(ORMArray.ORMElements, ORM)
}
func (ORMArray *ORMArray) SetDatabaseConnection(DatabaseConnectionLink *gorm.DB) {

	for Index := range ORMArray.ORMElements {

		ORMArray.ORMElements[Index].SetConnection(DatabaseConnectionLink)
	}
}
func (ORMArray *ORMArray) FindByName(Name string) (ORM ORMInterface, Error error) {
	for Index, Element := range ORMArray.ORMElements {
		if Element.GetName() == Name {
			return ORMArray.ORMElements[Index], nil
		}

	}
	return ORM, errors.New("ORM с таким именем не найдена")

}
