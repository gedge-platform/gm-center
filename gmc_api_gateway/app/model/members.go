package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	_id        primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Id         string             `json:"memberId,omitempty" bson:"memberId" validate:"required"`
	Name       string             `json:"memberName,omitempty" bson:"memberName" validate:"required"`
	Password   string             `json:"password,omitempty" bson:"password" validate:"required,gte=0,lte=10"`
	Email      string             `json:"email,omitempty" bson:"email" validate:"required,email"`
	Contact    string             `json:"contact,omitempty" bson:"contact"`
	Enabled    int                `json:"enabled,omitempty" bson:"enabled" validate:"gte=0,lte=1"`
	RoleName   string             `json:"memberRole,omitempty" bson:"memberRole"`
	Created_at primitive.DateTime `json:"created_at,omitempty"`
	Logined_at primitive.DateTime `json:"logined_at,omitempty"`
}

type RequestMember struct {
	_id        primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Id         string             `json:"memberId,omitempty" bson:"memberId"`
	Name       string             `json:"memberName,omitempty" bson:"memberName"`
	Password   string             `json:"password,omitempty" bson:"password" validate:"gte=0,lte=10"`
	Email      string             `json:"email,omitempty" bson:"email" validate:"email"`
	Contact    string             `json:"contact,omitempty" bson:"contact"`
	Enabled    int                `json:"enabled,omitempty" bson:"enabled" validate:"gte=0,lte=1"`
	RoleName   string             `json:"memberRole,omitempty" bson:"memberRole"`
	Created_at primitive.DateTime `json:"created_at,omitempty"`
	Logined_at primitive.DateTime `json:"logined_at,omitempty"`
}

func (m *Member) Enable() {
	m.Enabled = 1
}

func (m *Member) Disable() {
	m.Enabled = 0
}
