package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	_id         primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Id 			string             `json:"requestId,omitempty" bson:"requestId" validate:"required"`
	Status      string             `json:"status,omitempty" bson:"status"`
	Message     string             `json:"message,omitempty" bson:"message"`
	Workspace   primitive.ObjectID             `json:"workspace,omitempty" bson:"workspace" validate:"required"`
	Project     primitive.ObjectID             `json:"project,omitempty" bson:"project" validate:"required"`
	Date        primitive.DateTime `json:"date,omitempty" bson:"date"`
	Cluster     primitive.ObjectID`json:"cluster,omitempty" bson:"cluster" validate:"required"`
	Name        string             `json:"name,omitempty" bson:"name" validate:"required"`
	Reason      string             `json:"reason,omitempty" bson:"reason" validate:"required"`
	// Type        string             `json:"type,omitempty" bson:"type" validate:"required"`
}

type RequestUpdate struct {
	_id         primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Id 			string             `json:"requestId,omitempty" bson:"requestId"`
	Status      string             `json:"status,omitempty" bson:"status" validate:"required"`
	Message     string             `json:"message,omitempty" bson:"message"`
	Workspace   primitive.ObjectID             `json:"workspace,omitempty" bson:"workspace"`
	Project     primitive.ObjectID             `json:"project,omitempty" bson:"project"`
	Date        primitive.DateTime `json:"date,omitempty" bson:"date" validate:"required"`
	Cluster     primitive.ObjectID`json:"cluster,omitempty" bson:"cluster"`
	Name        string             `json:"name,omitempty" bson:"name"`
	Reason      string             `json:"reason,omitempty" bson:"reason"`
	// Type        string             `json:"type,omitempty" bson:"type"`
}