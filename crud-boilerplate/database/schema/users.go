package schema

type Users struct { // table name: users
	BaseUUID
	Name              string `gorm:"not null"`
	FirstName         string `gorm:"null"`
	LastName          string `gorm:"null"`
	Email             string `gorm:"not null;unique"`
	Password          string `gorm:"not null"`
	ConfirmationToken string `gorm:"null"`
	AuthToken         string `gorm:"null"`
	Role              string `gorm:"null"`
	Image             string `gorm:"null"`
	Nip               string `gorm:"null"`
	StatusID          int32  `gorm:"default:1"`
}

func (Users) TableName() string {
	return "users"
}

func (Users) Pk() string {
	return "id"
}

func (f Users) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f Users) AddForeignKeys() {
}

func (f Users) InsertDefaults() {
}
