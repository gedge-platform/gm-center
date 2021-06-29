package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
	gorm.Model
	Num 				int 			`gorm:"primaryKey" json:"num"`
	Id 					string 			`json:"id"`
	Name 				string 			`json:"name"`
	Email 				string 			`json:"email"`
	Contact 			string 			`json:"contact"`
	Description 		string 			`json:"description"`
	enabled 			bool 			`json:"enabled"`
	Created_at 			*time.Time 		`json:"created_at"`
	Logined_at 			*time.Time 		`json:"logined_at"`
	RoleName 			string 			`json:"roleName"`
	
	// text    string `gorm:"unique" json:"text"`
	// text bool   `json:"text"`
	// text    []text `gorm:"ForeignKey:textID" json:"text"`
  }

func (m *Member) Enabled() {
	m.enabled = true
}

func (m *Member) Disabled() {
	m.enabled = false
}


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
// func DBMigrate(db *gorm.DB) *gorm.DB {
// 	db.AutoMigrate(&Project{}, &Task{})
// 	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
// 	return db
// }
