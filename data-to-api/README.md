# Gedge Database to Rest API
A RESTful API for GEdge Database with Go

## Installation & Run
```bash
# Download this project
go get github.com/gedge-platform/gm-center/develop/data-to-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](github.com/gedge-platform/gm-center/develop/data-to-api/blob/master/config/config.go)
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
cd data-to-api
go build
./data-to-api

# API Endpoint : http://127.0.0.1:8000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── members.go  // APIs for Member model
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

---

### To do ✓
- [x] MEMBER_INFO
- [ ] CLUSTER_INFO
- [ ] APPSTORE_INFO
- [ ] PERMISSION_INFO
- [ ] PROJECT_INFO
- [ ] ROLE_INFO
- [ ] WORKSPACE_INFO
- [ ] APP_DETAIL


### In Progress
- [x] MEMBER_INFO
  - [x] GetAllMembers(GET, "/members")
  - [x] CreateMember(POST, "/members")
  - [x] GetMember(GET, "/members/{id}")
  - [ ] UpdateMember(PUT, "/members/{id}")
  - [x] DeleteMember(DELETE, "/members/{id}")
  - [ ] EnabledMember(PUT, "/members/{id}")
  - [ ] DisabledMember(DELETE, "/members/{id}")


### Done ✓
- [x] First Commit

