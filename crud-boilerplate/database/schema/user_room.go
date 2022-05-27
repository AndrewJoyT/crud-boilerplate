package schema

type UserRoom struct { // table name: UserRoom
	BaseUUID
	UserUUID string `gorm:"not null"`
	RoomID   string `gorm:"not null"`
	User     Users
	Courses  []Course
}

func (UserRoom) TableName() string {
	return "user_room"
}

func (UserRoom) Pk() string {
	return "id"
}

func (f UserRoom) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f UserRoom) AddForeignKeys() {
}

func (f UserRoom) InsertDefaults() {
}
