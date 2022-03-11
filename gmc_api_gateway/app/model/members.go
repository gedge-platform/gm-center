package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	_id        primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Id         string             `json:"id,omitempty" bson:"memberId" validate:"required"`
	Name       string             `json:"name,omitempty" bson:"memberName" validate:"required"`
	Password   string             `json:"password,omitempty" bson:"password" validate:"required,gte=0,lte=10"`
	Email      string             `json:"email,omitempty" bson:"email" validate:"required,email"`
	Contact    string             `json:"contact,omitempty" bson:"contact"`
	Enabled    int                `json:"enabled,omitempty" bson:"enabled" validate:"gte=0,lte=1"`
	RoleName   string             `json:"role,omitempty" bson:"memberRole"`
	Created_at string             `json:"created_at,omitempty"`
	Logined_at string             `json:"logined_at,omitempty"`
}

// type MemberWithPassword struct {
// 	Member
// 	Password string `json:"password" bson:"password, omitempty"`
// }

func (m *Member) Enable() {
	m.Enabled = 1
}

func (m *Member) Disable() {
	m.Enabled = 0
}
