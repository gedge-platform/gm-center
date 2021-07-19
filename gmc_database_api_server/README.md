# GM-Center Database API Server
A RESTful API for GM-Center Database with Go

## Installation & Run
```bash
# Download this project
go get github.com/gedge-platform/gm-center/main/gmc_database_api_server
```

Before running API server, you should set the .env file with yours or set the your .env with my values on [.env](github.com/gedge-platform/gm-center/blob/main/gmc_database_api_server/config/config.go)
```go
DB_DIALECT=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=gedge
DB_CHARSET=utf8

PORT=:8008
CORS=CORS_ORIGIN
```

```bash
# Build and Run
cd gmc_database_api_server
go build && ./main
# or
go run main.go

# API Endpoint : http://127.0.0.1:8000
```

## Structure
```
├── app
│   ├── app.go
│   ├── api          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── members.go  // APIs for Member model
│   │   ├── clusters.go  // APIs for clusters model
│   │   ├── projects.go  // APIs for clusters model
│   │   ├── workspaces.go  // APIs for clusters model
│   │   ├── apps.go  // APIs for clusters model
│   │   ├── custom.go // APIs for kubernetes model
│   └── db
│       └── db.go
│   └── model
│       └── apps.go
│       └── clusters.go
│       └── members.go
│       └── projects.go
│       └── workspaces.go
│       └── kubernetes.go
│   └── routes
│       └── routes.go
├── config
│   └── config.go        // Configuration
└── .env.sample
└── main.go
└── go.mod
└── README.md
```

## API

#### Lists
- members
- clusters
- projects
- workspaces
- apps
- kubernetes

<br />

### members, clusters, projects, workspaces, apps
---
#### /api/v1/{lists_name}
* `GET` : Get all {lists_name}
* `POST` : Create a new {lists_name}

#### /api/v1/{lists_name}/:{id or name}
* `GET` : Get a {lists_name}/:{id or name}
* `PUT` : Update a {lists_name}/:{id or name}
* `DELETE` : Delete a {lists_name}/:{id or name}


<br />

### kubernetes
---

#### /api/v2/:cluster_name/:namespace_name
* `GET` : Get a {lists_name}/:namespace_name
* `CREATE` : Create a {lists_name}/:namespace_name
* `PUT` : Update a {lists_name}/:namespace_name
* `PATCH` : Patch a {lists_name}/:namespace_name
* `DELETE` : Delete a {lists_name}/:namespace_name

#### /api/v2/:cluster_name/:namespace_name/:kind_name
* `GET` : Get a {lists_name}/:namespace_name/:kind_name
* `CREATE` : Create a {lists_name}/:namespace_name/:kind_name
* `PUT` : Update a {lists_name}/:namespace_name/:kind_name
* `PATCH` : Patch a {lists_name}/:namespace_name/:kind_name
* `DELETE` : Delete a {lists_name}/:namespace_name/:kind_name

#### /api/v2/:cluster_name/:namespace_name/:kind_name/*
* `GET` : Get a {lists_name}/:namespace_name/:kind_name/*
* `CREATE` : Create a {lists_name}/:namespace_name/:kind_name/*
* `PUT` : Update a {lists_name}/:namespace_name/:kind_name/*
* `PATCH` : Patch a {lists_name}/:namespace_name/:kind_name/*
* `DELETE` : Delete a {lists_name}/:namespace_name/:kind_name/*

---

### To do ✓
- [x] MEMBER_INFO
- [x] CLUSTER_INFO
- [X] PROJECT_INFO
- [x] WORKSPACE_INFO
- [x] APPSTORE_INFO
- [x] KUBERNETES
- [ ] APP_DETAIL


### In Progress
- [x] APP_DETAIL

### Done ✓
- [x] First Commit
- [x] MEMBER_INFO
  - [x] GetAllMembers(GET, "/members")
  - [x] CreateMember(POST, "/members")
```
{
    "memberId": "memberId",
    "memberName": "memberName",
    "memberEmail": "member@gedge.com",
    "memberPassword": "memberPassword"
}
```
  - [x] GetMember(GET, "/members/{id}")
  - [x] UpdateMember(PUT, "/members/{id}")
```
{
    "memberId": "memberId",
    "memberName": "memberName",
    "memberEmail": "member@gedge.com",
    "memberPassword": "memberPassword"
}
```
  - [x] DeleteMember(DELETE, "/members/{id}")

- [x] CLUSTER_INFO
  - [x] GetAllClusters(GET, "/clusters")
  - [x] CreateCluster(POST, "/clusters")
```
{
	"ipAddr": "127.0.0.1",
	"clusterName": "value",
	"clusterRole": "value",
	"clusterType": "value",
	"clusterEndpoint": "10.10.10.10",
	"clusterCreator": "value",
}
```
  - [x] GetCluster(GET, "/clusters/{name}")
  - [x] UpdateCluster(PUT, "/clusters/{name}")
```
{
	"ipAddr": "127.0.0.1",
	"clusterName": "value",
	"clusterRole": "value",
	"clusterType": "value",
	"clusterEndpoint": "10.10.10.10",
	"clusterCreator": "value",
}
```
  - [x] DeleteCluster(DELETE, "/clusters/{name}")


- [x] WORKSPACE_INFO
  - [x] GetAllWorkspaces(GET, "/workspaces")
  - [x] CreateWorkspace(POST, "/workspaces")
```
{
	"clusterName": "value",
	"workspaceName": "value",
	"workspaceDescription": "value",
	"selectCluster": "1,3",
	"workspaceOwner": "value",
	"workspaceCreator": "value"
}
```
  - [x] GetWorkspace(GET, "/workspaces/{name}")
  - [x] UpdateWorkspace(PUT, "/workspaces/{name}")
```
{
	"clusterName": "value",
	"workspaceName": "value",
	"workspaceDescription": "value",
	"selectCluster": "1,3",
	"workspaceOwner": "value",
	"workspaceCreator": "value"
}
```
  - [x] DeleteWorkspace(DELETE, "/workspaces/{name}")


- [x] PROJECT_INFO
  - [x] GetAllProjects(GET, "/projects")
  - [x] CreateProject(POST, "/projects")
```
{
	"projectName": "value",
	"projectPostfix": "value",
	"projectDescription": "value",
	"projectType": "value",
	"projectOwner": "value",
	"projectCreator": "value",
	"workspaceName": "value"
}
```
  - [x] GetProject(GET, "/projects/{name}")
  - [x] UpdateProject(PUT, "/projects/{name}")
```
{
	"projectName": "value",
	"projectPostfix": "value",
	"projectDescription": "value",
	"projectType": "value",
	"projectOwner": "value",
	"projectCreator": "value",
	"workspaceName": "value"
}
```
  - [x] DeleteProject(DELETE, "/projects/{name}")

- [x] APPSTORE_INFO
  - [x] GetAllApps(GET, "/apps")
  - [x] CreateApp(POST, "/apps")
```
{
	"appName": "value",
	"appDescription": "value",
	"appCategory": "value",
	"appInstalled": 0
}
```
  - [x] GetApp(GET, "/apps/{name}")
  - [x] UpdateApp(PUT, "/apps/{name}")
```
{
	"appName": "value",
	"appDescription": "value",
	"appCategory": "value",
	"appInstalled": 0
}
```
  - [x] DeleteApp(DELETE, "/apps/{name}")

