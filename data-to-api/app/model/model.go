package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
	// gorm.Model
	Num 				int 			`gorm:"column:memberNum" json:"num"`
	Id 					string 			`gorm:"column:memberId" json:"id"`
	Name 				string 			`gorm:"column:memberName" json:"name"`
	Email 				string 			`gorm:"column:memberEmail" json:"email"`
	Contact 			string 			`gorm:"column:memberContact" json:"contact"`
	Description 		string 			`gorm:"column:memberDescription" json:"description"`
	enabled 			bool 			`gorm:"column:enabled" json:"enabled"`
	Created_at 			*time.Time 		`gorm:"column:Created_at" json:"created_at"`
	Logined_at 			*time.Time 		`gorm:"column:Logined_at" json:"logined_at"`
	RoleName 			string 			`gorm:"column:roleName" json:"roleName"`
	
	// text    string `gorm:"unique" json:"text"`
	// text bool   `json:"text"`
	// text    []text `gorm:"ForeignKey:textID" json:"text"`
  }

// Set Member table name to be `MEMBER_INFO`
func (Member) TableName() string {
	return "MEMBER_INFO"
  }

func (m *Member) Enabled() {
	m.enabled = true
}

func (m *Member) Disabled() {
	m.enabled = false
}


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Member{})
	// db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
