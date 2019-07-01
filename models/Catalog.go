package models

type Catalog struct {
	ID       int `gorm:"primary_key"`
	Name     string
	ParentID int // which catalog it belongs toj
}

func (Catalog) TableName() string {
	return "catalog"
}
