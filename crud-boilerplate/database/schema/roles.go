package schema

// type Roles struct { // table name: roles
// 	Base
// 	Name string `gorm:"not null"`
// }

// func (Roles) TableName() string {
// 	return "roles"
// }

// func (Roles) Pk() string {
// 	return "id"
// }

// func (f Roles) Ref() string {
// 	return f.TableName() + "(" + f.Pk() + ")"
// }

// func (f Roles) AddForeignKeys() {
// }

// func (f Roles) InsertDefaults() {
// 	Database.Exec(`INSERT INTO roles (id, name)
// 	SELECT * FROM (SELECT 1, 'Admin') AS tmp
// 		WHERE NOT EXISTS (
//     		SELECT name FROM roles WHERE name = 'Admin'
// 	) LIMIT 1;`)
// 	Database.Exec(`INSERT INTO roles (id, name)
// 	SELECT * FROM (SELECT 2, 'Admin Sekolah') AS tmp
// 		WHERE NOT EXISTS (
//     		SELECT name FROM roles WHERE name = 'Admin Sekolah'
// 	) LIMIT 1;`)
// 	Database.Exec(`INSERT INTO roles (id, name)
// 	SELECT * FROM (SELECT 3, 'Guru') AS tmp
// 		WHERE NOT EXISTS (
//     		SELECT name FROM roles WHERE name = 'Guru'
// 	) LIMIT 1;`)
// 	Database.Exec(`INSERT INTO roles (id, name)
// 	SELECT * FROM (SELECT 4, 'Murid') AS tmp
// 		WHERE NOT EXISTS (
//     		SELECT name FROM roles WHERE name = 'Murid'
// 	) LIMIT 1;`)
// }
