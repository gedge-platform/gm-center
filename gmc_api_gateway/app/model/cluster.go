package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cluster struct {
	_id      primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Endpoint string             `json:"clusterEndpoint,omitempty" bson:"clusterEndpoint" validate:"required"`
	Type     string             `json:"clusterType,omitempty" bson:"clusterType" validate:"required"`
	Name     string             `json:"clusterName,omitempty" bson:"clusterName" validate:"required"`
	Token    string             `json:"token,omitempty" bson:"token" validate:"required"`
}

type RequestCluster struct {
	_id      primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Endpoint string             `json:"clusterEndpoint,omitempty" bson:"clusterEndpoint"`
	Type     string             `json:"clusterType,omitempty" bson:"clusterType"`
	Name     string             `json:"clusterName,omitempty" bson:"clusterName"`
	Token    string             `json:"token,omitempty" bson:"token"`
}

