package schema

type Course struct { // table name: Course
	BaseUUID
	Name      string `gorm:"not null"`
	CompanyID string `gorm:"null"`
	Files     []File
}

func (Course) TableName() string {
	return "course"
}

func (Course) Pk() string {
	return "id"
}

func (f Course) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f Course) AddForeignKeys() {
}

func (f Course) InsertDefaults() {
}
