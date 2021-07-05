package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Member struct {
	Num 				int 			`gorm:"column:memberNum; primary_key" json:"memberNum"`
	Id 					string 			`gorm:"column:memberId; not null" json:"memberId"`
	Name 				string 			`gorm:"column:memberName; not null" json:"memberName"`
	Email 				*string 		`gorm:"column:memberEmail; not null" json:"memberEmail"`
	Contact 			string 			`gorm:"column:memberContact; not null" json:"memberContact,omitempty"`
	Description 		string 			`gorm:"column:memberDescription" json:"memberDescription,omitempty"`
	Enabled 			int				`gorm:"column:memberEnabled; DEFAULT:0" json:"memberEnabled,omitempty"`
	Created_at 			time.Time 		`gorm:"column:created_at; DEFAULT:''" json:"created_at"`
	Logined_at 			time.Time 		`gorm:"column:logined_at" json:"logined_at"`
	RoleName 			string 			`gorm:"column:roleName; DEFAULT:''" json:"roleName"`
  }


  type MemberWithPassword struct {
	Member
	Password				string 			`gorm:"column:memberPassword" json:"memberPassword"`
  }


// Set Member table name to be `MEMBER_INFO`
func (Member) TableName() string {
	return "MEMBER_INFO"
  }

// Set Member table name to be `MEMBER_INFO`
func (MemberWithPassword) TableName() string {
	return "MEMBER_INFO"
  }

func (m *Member) Enable() {
	m.Enabled = 1
}

func (m *Member) Disable() {
	m.Enabled = 0
}


  type Cluster struct {
	Num				int			`gorm:"column:clusterNum; primary_key" json:"clusterNum"`
	Ip				string			`gorm:"column:ipAddr; not null" json:"ipAddr"`
	extIp				string			`gorm:"column:extIpAddr" json:"extIpAddr"`
	Name 				string 			`gorm:"column:clusterName; not null" json:"clusterName"`
	Role 				*string 			`gorm:"column:clusterRole; not null" json:"clusterRole"`
	Type				string			`gorm:"column:clusterType; not null" json:"clusterType"`
	Endpoint				string			`gorm:"column:clusterEndpoint; not null" json:"clusterEndpoint"`
	Creator				string			`gorm:"column:clusterCreator; not null" json:"clusterCreator"`
	State				string			`gorm:"column:clusterState; not null; DEFAULT:'pending'" json:"clusterState"`
	Version				string			`gorm:"column:kubeVersion" json:"kubeVersion"`
	Created_at 			time.Time 		`gorm:"column:created_at" json:"created_at"`
  }

// Set Cluster table name to be `CLUSTER_INFO`
func (Cluster) TableName() string {
	return "CLUSTER_INFO"
  }


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Member{})
	return db
}
