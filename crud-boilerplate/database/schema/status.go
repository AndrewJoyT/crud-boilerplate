package schema

type Statuses struct { // table name: statuses
	Base
	Name string `gorm:"not null"`
}

func (Statuses) TableName() string {
	return "statuses"
}

func (Statuses) Pk() string {
	return "id"
}

func (f Statuses) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f Statuses) AddForeignKeys() {
}

func (f Statuses) InsertDefaults() {
	// Database.Exec(`INSERT INTO statuses (id, name)
	// SELECT * FROM (SELECT 1, 'Registered') AS tmp
	// 	WHERE NOT EXISTS (
	// 		SELECT name FROM statuses WHERE name = 'Registered'
	// ) LIMIT 1;`)
	// Database.Exec(`INSERT INTO statuses (id, name)
	// SELECT * FROM (SELECT 2, 'Confirmed') AS tmp
	// 	WHERE NOT EXISTS (
	// 		SELECT name FROM statuses WHERE name = 'Confirmed'
	// ) LIMIT 1;`)
	// Database.Exec(`INSERT INTO statuses (id, name)
	// SELECT * FROM (SELECT 3, 'Active') AS tmp
	// 	WHERE NOT EXISTS (
	// 		SELECT name FROM statuses WHERE name = 'Active'
	// ) LIMIT 1;`)
	// Database.Exec(`INSERT INTO statuses (id, name)
	// SELECT * FROM (SELECT 4, 'On Hold') AS tmp
	// 	WHERE NOT EXISTS (
	// 		SELECT name FROM statuses WHERE name = 'On Hold'
	// ) LIMIT 1;`)
}
