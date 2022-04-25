package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	_id             primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name            string             `json:"name,omitempty" bson:"name" validate:"required"`
	Type            string             `json:"type,omitempty" bson:"type" validate:"required"`
	Domain          string             `json:"domain,omitempty" bson:"domain"`
	Region          string             `json:"region,omitempty" bson:"region"`
	Url             string             `json:"url,omitempty" bson:"url"`
	Tenant          string             `json:"tenant,omitempty" bson:"tenant"`
	Access_id       string             `json:"access_id,omitempty" bson:"access_id"`
	Access_token    string             `json:"access_token,omitempty" bson:"access_token"`
	Project         string             `json:"project,omitempty" bson:"project"`
	Subscription_id string             `json:"subscription_id,omitempty" bson:"subscription_id"`
	//KeyValueInfoList map[string]interface{} `json:"KeyValueInfoList,omitempty`
	Created_at primitive.DateTime `json:"created_at,omitempty"`
}

type KeyValueInfoList struct {
	_id              primitive.ObjectID `json:"objectId,omitempty"`
	DomainName       string             `json:"DomainName,omitempty"`
	IdentityEndPoint string             `json:"IdentityEndPoint,omitempty"`
	Password         string             `json:"Password,omitempty"`
	ProjectID        string             `json:"ProjectID,omitempty"`
	Username         string             `json:"Username,omitempty"`
	// Project         string             `json:"project,omitempty" bson:"project"`
	// Subscription_id string             `json:"subscription_id,omitempty" bson:"subscription_id"`
}
