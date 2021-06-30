package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
	Num 				int 			`gorm:"column:memberNum; primary_key" json:"memberNum"`
	Id 					string 			`gorm:"column:memberId; DEFAULT:''; not null" json:"memberId"`
	Name 				string 			`gorm:"column:memberName; DEFAULT:''; not null" json:"memberName"`
	Email 				*string 		`gorm:"column:memberEmail; DEFAULT:''; not null" json:"memberEmail"`
	Contact 			string 			`gorm:"column:memberContact" json:"memberContact,omitempty"`
	Description 		string 			`gorm:"column:memberDescription" json:"memberDescription,omitempty"`
	enabled 			int				`gorm:"column:memberEnabled; DEFAULT:0" json:"memberEnabled,omitempty"`
	Created_at 			time.Time 		`gorm:"column:created_at; DEFAULT:''" json:"created_at"`
	Logined_at 			time.Time 		`gorm:"column:logined_at" json:"logined_at"`
	RoleName 			string 			`gorm:"column:roleName; DEFAULT:''" json:"roleName"`
  }

  type MemberWithPassword struct {
	Member
	Password			string 			`gorm:"column:memberPassword; DEFAULT:''; not null" json:"memberPassword"`
  }


// Set Member table name to be `MEMBER_INFO`
func (Member) TableName() string {
	return "MEMBER_INFO"
  }

// Set Member table name to be `MEMBER_INFO`
func (MemberWithPassword) TableName() string {
	return "MEMBER_INFO"
  }


func (m *Member) Enabled() {
	m.enabled = 1
}

func (m *Member) Disabled() {
	m.enabled = 0
}


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Member{})
	// db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
