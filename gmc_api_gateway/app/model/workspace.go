package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	_id           primitive.ObjectID  `json:"objectId,omitempty" bson:"_id"`
	Name          string              `json:"workspaceName,omitempty" bson:"workspaceName"`
	Description   string              `json:"workspaceDescription,omitempty" bson:"workspaceDescription"`
	Owner         primitive.ObjectID  `json:"workspaceOwner,omitempty" bson:"workspaceOwner"`
	Creator       primitive.ObjectID  `json:"workspaceCreator,omitempty" bson:"workspaceCreator"`
	Selectcluster []WorkspaceClusters `json:"selectCluster,omitempty" bson:"selectCluster"`
}

type WorkspaceClusters struct {
	Cluster primitive.ObjectID `json:"cluster,omitempty" bson:"cluster"`
}
