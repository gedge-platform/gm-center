package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	_id          primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name         string             `json:"name,omitempty" bson:"name" validate:"required"`
	Type         string             `json:"type,omitempty" bson:"type" validate:"required"`
	Domain       string             `json:"domain,omitempty" bson:"domain" validate:"required"`
	Region       string             `json:"region,omitempty" bson:"region" validate:"required"`
	Url          string             `json:"url,omitempty" bson:"url" validate:"required"`
	Tenant       string             `json:"tenant,omitempty" bson:"tenant" validate:"required"`
	Access_id    string             `json:"access_id,omitempty" bson:"access_id" validate:"required"`
	Access_token string             `json:"access_token,omitempty" bson:"access_token" validate:"required"`
	Project      string             `json:"project,omitempty" bson:"project" validate:"required"`
	//Created_at   primitive.DateTime `json:"created_at,omitempty"`
}

type RequestCredentials struct {
	_id          primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name         string             `json:"name,omitempty" bson:"name"`
	Type         string             `json:"type,omitempty" bson:"type"`
	Domain       string             `json:"domain,omitempty" bson:"domain"`
	Region       string             `json:"region,omitempty" bson:"region"`
	Url          string             `json:"url,omitempty" bson:"url"`
	Tenant       string             `json:"tenant,omitempty" bson:"tenant"`
	Access_id    string             `json:"access_id,omitempty" bson:"access_id"`
	Access_token string             `json:"access_token,omitempty" bson:"access_token"`
	Project      string             `json:"project,omitempty" bson:"project"`
	//Created_at   primitive.DateTime `json:"created_at,omitempty"`
}

type NewCredentials struct {
	_id          primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name         string             `json:"name,omitempty" bson:"name" validate:"required"`
	Type         string             `json:"type,omitempty" bson:"type" validate:"required"`
	Region       string             `json:"region,omitempty" bson:"region" validate:"required"`
	Domain       string             `json:"domain,omitempty" bson:"domain" validate:"required"`
	Url          string             `json:"url,omitempty" bson:"url" validate:"required"`
	Tenant       string             `json:"tenant,omitempty" bson:"tenant" validate:"required"`
	Access_id    string             `json:"access_id,omitempty" bson:"access_id" validate:"required"`
	Access_token string             `json:"access_token,omitempty" bson:"access_token" validate:"required"`
	Project      string             `json:"project,omitempty" bson:"project" validate:"required"`
	//Created_at   primitive.DateTime `json:"created_at,omitempty"`
}
