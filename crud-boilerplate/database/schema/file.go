package schema

type File struct { // table name: File
	BaseUUID
	Name     string `gorm:"not null"`
	FileUrl  string `gorm:"null"`
	CourseId string `gorm:"null"`
}

func (File) TableName() string {
	return "file"
}

func (File) Pk() string {
	return "id"
}

func (f File) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f File) AddForeignKeys() {
}

func (f File) InsertDefaults() {
}
