package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	_id           primitive.ObjectID  `json:"objectId,omitempty" bson:"_id"`
	Name          string              `json:"workspaceName,omitempty" bson:"workspaceName" validate:"required"`
	Description   string              `json:"workspaceDescription,omitempty" bson:"workspaceDescription" validate:"required"`
	MemberName    string			  `json:"memberName,omitempty" bson:"memberName" validate:"required"`
	ClusterName   []string  		  `json:"clusterName,omitempty" bson:"selectCluster2"`
	Selectcluster []WorkspaceClusters `json:"selectCluster,omitempty" bson:"selectCluster"`
}

type WorkspaceClusters struct {
	Cluster primitive.ObjectID `json:"cluster,omitempty" bson:"cluster"`
}

type RequestWorkspace struct {
	_id           primitive.ObjectID   `json:"objectId,omitempty" bson:"_id"`
	Name          string               `json:"workspaceName,omitempty" bson:"workspaceName"`
	Description   string               `json:"workspaceDescription,omitempty" bson:"workspaceDescription"`
	Owner         primitive.ObjectID   `json:"workspaceOwner,omitempty" bson:"workspaceOwner"`
	Creator       primitive.ObjectID   `json:"workspaceCreator,omitempty" bson:"workspaceCreator"`
	MemberName    string			   `json:"memberName,omitempty" bson:"memberName"`
	ClusterName   []string  	   	   `json:"clusterName,omitempty" bson:"selectCluster2"`
	Selectcluster []primitive.ObjectID `json:"selectCluster,omitempty" bson:"selectCluster"`
}

type NewWorkspace struct {
	_id           primitive.ObjectID   `json:"objectId,omitempty" bson:"_id"`
	Name          string               `json:"workspaceName,omitempty" bson:"workspaceName" validate:"required"`
	Description   string               `json:"workspaceDescription,omitempty" bson:"workspaceDescription" validate:"required"`
	Owner         primitive.ObjectID   `json:"workspaceOwner,omitempty" bson:"workspaceOwner" validate:"required"`
	Creator       primitive.ObjectID   `json:"workspaceCreator,omitempty" bson:"workspaceCreator" validate:"required"`
	Selectcluster []primitive.ObjectID `json:"selectCluster,omitempty" bson:"selectCluster"`
}