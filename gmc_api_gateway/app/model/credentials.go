package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	_id          primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name         string             `json:"name,omitempty" bson:"credentialsName"`
	Type         string             `json:"type,omitempty" bson:"credentialsType"`
	Domain       string             `json:"domain,omitempty" bson:"credentialsDomain"`
	Region       string             `json:"region,omitempty" bson:"credentialsRegion"`
	Url          string             `json:"url,omitempty" bson:"url"`
	Tenant       string             `json:"userTenant,omitempty" bson:"tenant"`
	Access_id    string             `json:"access_id,omitempty" bson:"access_id"`
	Access_token string             `json:"access_token,omitempty" bson:"access_token"`
	Project      primitive.ObjectID `json:"project,omitempty" bson:"project"`
	//Created_at   primitive.DateTime `json:"created_at,omitempty"`
}

/* type CredentialsProject struct {
	Project primitive.ObjectID `json:"project,omitempty" bson:"project"`
} */

type RequestCredentials struct {
	_id          primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name         string             `json:"name,omitempty" bson:"credentialsName"`
	Type         string             `json:"type,omitempty" bson:"credentialsType"`
	Domain       string             `json:"domain,omitempty" bson:"credentialsDomain"`
	Region       string             `json:"region,omitempty" bson:"credentialsRegion"`
	Url          string             `json:"url,omitempty" bson:"url"`
	Tenant       string             `json:"userTenant,omitempty" bson:"tenant"`
	Access_id    string             `json:"access_id,omitempty" bson:"access_id"`
	Access_token string             `json:"access_token,omitempty" bson:"access_token"`
	Project      primitive.ObjectID `json:"project,omitempty" bson:"project"`
	//Created_at   primitive.DateTime `json:"created_at,omitempty"`
}

type NewCredentials struct {
	_id          primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name         string             `json:"name,omitempty" bson:"name" validate:"required"`
	Type         string             `json:"type,omitempty" bson:"type" validate:"required"`
	Region       string             `json:"region,omitempty" bson:"region" validate:"required"`
	Domain       string             `json:"domain,omitempty" bson:"domain" validate:"required"`
	Url          string             `json:"url,omitempty" bson:"url" validate:"required"`
	Tenant       string             `json:"userTenant,omitempty" bson:"tenant" validate:"required"`
	Access_id    string             `json:"access_id,omitempty" bson:"access_id" validate:"required"`
	Access_token string             `json:"access_token,omitempty" bson:"access_token" validate:"required"`
	Project      primitive.ObjectID `json:"project,omitempty" bson:"project" validate:"required"`
	//Created_at   primitive.DateTime `json:"created_at,omitempty"`
}
