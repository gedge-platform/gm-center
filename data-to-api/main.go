package main

import (
  "fmt"
  
	"github.com/gedge-platform/gm-center/develop/data-to-api/app"
	"github.com/gedge-platform/gm-center/develop/data-to-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	fmt.Println("Server started on: http://localhost:8000")
	app.Run(":8000")
}
