package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	_id           primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
	Name          string             `json:"projectName,omitempty" bson:"projectName" validate:"required"`
	Description   string             `json:"projectDescription,omitempty" bson:"projectDescription" validate:"required"`
	Type          string             `json:"projectType,omitempty" bson:"projectType" validate:"required"`
	Owner         primitive.ObjectID `json:"projectOwner,omitempty" bson:"projectOwner"`
	Creator       primitive.ObjectID `json:"projectCreator,omitempty" bson:"projectCreator"`
	MemberName    string             `json:"memberName,omitempty" bson:"memberName" validate:"required"`
	WorkspaceName string             `json:"workspaceName,omitempty" bson:"workspaceName" validate:"required"`
	Workspace     primitive.ObjectID `json:"workspace,omitempty" bson:"workspace"`
	Selectcluster []ProjectClusters  `json:"selectCluster,omitempty" bson:"selectCluster"`
	ClusterName   []string           `json:"clusterName,omitempty" bson:"selectCluster2"`
	Created_at    primitive.DateTime `json:"created_at,omitempty"`
}

type ProjectClusters struct {
	Cluster primitive.ObjectID `json:"cluster,omitempty" bson:"cluster"`
}

type RequestProject struct {
	_id           primitive.ObjectID   `json:"objectId,omitempty" bson:"_id"`
	Name          string               `json:"projectName,omitempty" bson:"projectName"`
	Description   string               `json:"projectDescription,omitempty" bson:"projectDescription"`
	Type          string               `json:"projectType,omitempty" bson:"projectType"`
	Owner         primitive.ObjectID   `json:"projectOwner,omitempty" bson:"projectOwner"`
	Creator       primitive.ObjectID   `json:"projectCreator,omitempty" bson:"projectCreator"`
	MemberName    string               `json:"memberName,omitempty" bson:"memberName"`
	WorkspaceName string               `json:"workspaceName,omitempty" bson:"workspaceName"`
	Workspace     primitive.ObjectID   `json:"workspace,omitempty" bson:"workspace"`
	Selectcluster []primitive.ObjectID `json:"selectCluster,omitempty" bson:"selectCluster"`
	ClusterName   []string             `json:"clusterName,omitempty" bson:"selectCluster2"`
	Created_at    primitive.DateTime   `json:"created_at,omitempty"`
}

type NewProject struct {
	_id           primitive.ObjectID   `json:"objectId,omitempty" bson:"_id"`
	Name          string               `json:"projectName,omitempty" bson:"projectName" validate:"required"`
	Description   string               `json:"projectDescription,omitempty" bson:"projectDescription"`
	Type          string               `json:"projectType,omitempty" bson:"projectType" validate:"required"`
	Owner         primitive.ObjectID   `json:"projectOwner,omitempty" bson:"projectOwner" validate:"required"`
	Creator       primitive.ObjectID   `json:"projectCreator,omitempty" bson:"projectCreator" validate:"required"`
	Workspace     primitive.ObjectID   `json:"workspace,omitempty" bson:"workspace" validate:"required"`
	Selectcluster []primitive.ObjectID `json:"selectCluster,omitempty" bson:"selectCluster" validate:"required"`
	Created_at    primitive.DateTime   `json:"created_at,omitempty"`
}
