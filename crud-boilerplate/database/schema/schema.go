package schema

import (
	"time"

	"github.com/jinzhu/gorm"
	satoriUUID "github.com/satori/go.uuid"
)

var (
	Database *gorm.DB // Scope is only within schema package
)

type TableInterface interface {
	Pk() string
	Ref() string
	AddForeignKeys()
	InsertDefaults()
}

type BaseUUID struct {
	ID        string     `gorm:"primary_key;type:varchar(255);"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedBy string     `gorm:"default:null"`
	UpdatedBy string     `gorm:"default:null"`
	DeletedAt *time.Time `sql:"index"`
}

type Base struct {
	ID        string     `gorm:"primary_key;type:int;"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	CreatedBy string     `gorm:"default:null"`
	UpdatedBy string     `gorm:"default:null"`
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseUUID) BeforeCreate(scope *gorm.Scope) error {
	uuid := satoriUUID.NewV4()
	return scope.SetColumn("ID", uuid.String())
}

func AutoMigrate(database *gorm.DB) {

	Database = database

	// Drop all the tables
	//database.DropTableIfExists()

	// Auto migrate
	database.AutoMigrate(
		&Statuses{},
		&Users{},
		&Course{},
		&File{},
		//&Roles{},
	)

	// Relationship

	// Insert defaults
	//Roles{}.InsertDefaults()
	Statuses{}.InsertDefaults()

	//Migration
	// database.Exec(`
	// 	ALTER TABLE roles MODIFY id INTEGER;
	// `)
}
