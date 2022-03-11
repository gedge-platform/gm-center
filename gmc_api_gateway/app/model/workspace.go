package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	_id           primitive.ObjectID  `json:"objectId,omitempty" bson:"_id"`
	Name          string              `json:"workspaceName,omitempty" bson:"workspaceName" validate:"required"`
	Description   string              `json:"workspaceDescription,omitempty" bson:"workspaceDescription" validate:"required"`
	Owner         primitive.ObjectID  `json:"workspaceOwner,omitempty" bson:"workspaceOwner" validate:"required"`
	Creator       primitive.ObjectID  `json:"workspaceCreator,omitempty" bson:"workspaceCreator" validate:"required"`
	Selectcluster []WorkspaceClusters `json:"selectCluster,omitempty" bson:"selectCluster"`
}

type WorkspaceClusters struct {
	Cluster primitive.ObjectID `json:"cluster,omitempty" bson:"cluster"`
}
