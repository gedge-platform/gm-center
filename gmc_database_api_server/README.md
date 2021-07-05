# GM-Center Database API Server
A RESTful API for GM-Center Database with Go

## Installation & Run
```bash
# Download this project
go get github.com/gedge-platform/gm-center/main/gmc_database_api_server
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](github.com/gedge-platform/gm-center/main/gmc_database_api_server/blob/main/config/config.go)
```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "username",
			Password: "userpassword",
			Name:     "gedge",
			Charset:  "utf8",
		},
	}
}
```

```bash
# Build and Run
cd gmc_database_api_server
go build
./gmc_database_api_server

# API Endpoint : http://127.0.0.1:8000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── members.go  // APIs for Member model
│   │   ├── clusters.go  // APIs for clusters model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
└── go.mod
└── README.md
```

## API

#### /members
* `GET` : Get all members
* `POST` : Create a new member

#### /members/:id
* `GET` : Get a member
* `PUT` : Update a member
* `DELETE` : Delete a member

#### /members/:id/enabled
* `PUT` : Enabled a member
* `DELETE` : Disabled a member 

#### /clusters
* `GET` : Get all clusters
* `POST` : Create a new clusters

#### /clusters/:name
* `GET` : Get a clusters
* `PUT` : Update a clusters
* `DELETE` : Delete a clusters

---

### To do ✓
- [x] MEMBER_INFO
- [x] CLUSTER_INFO
- [ ] APPSTORE_INFO
- [ ] PERMISSION_INFO
- [ ] PROJECT_INFO
- [ ] ROLE_INFO
- [ ] WORKSPACE_INFO
- [ ] APP_DETAIL


### In Progress
- [x] WORKSPACE_INFO

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

