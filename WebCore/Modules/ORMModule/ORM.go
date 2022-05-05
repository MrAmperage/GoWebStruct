package ORMModule

import "gorm.io/gorm"

func (ORM *ORM) SetConnection(ConnectionLink *gorm.DB) {
	ORM.ConnectionLink = ConnectionLink
}
func (ORM *ORM) GetName() string {
	return ORM.Name
}

func (ORM *ORM) SetName(Name string) {

	ORM.Name = Name
}
