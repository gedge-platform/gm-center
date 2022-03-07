package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	_id           primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name          string             `json:"projectName,omitempty" bson:"projectName"`
	Description   string             `json:"projectDescription,omitempty" bson:"projectDescription"`
	Type          string             `json:"projectType,omitempty" bson:"projectType"`
	Owner         primitive.ObjectID `json:"projectOwner,omitempty" bson:"projectOwner"`
	Creator       primitive.ObjectID `json:"projectCreator,omitempty" bson:"projectCreator"`
	Workspace     primitive.ObjectID `json:"workspace,omitempty" bson:"workspace"`
	Selectcluster []ProjectClusters  `json:"selectCluster,omitempty" bson:"selectCluster"`
	Created_at    string             `json:"created_at,omitempty"`
}

type ProjectClusters struct {
	Cluster primitive.ObjectID `json:"cluster,omitempty" bson:"cluster"`
}
