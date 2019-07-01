package models

// TODO: temporarily use relational database, use nosql in future
type FileRef struct {
	ID       int `gorm:"primary_key"`
	FileName string
	FileKey  string
}

func (FileRef) TableName() string {
	return "file_ref"
}
