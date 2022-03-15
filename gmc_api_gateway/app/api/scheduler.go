package api

import (
	"fmt"
  "bytes"
  "mime/multipart"
  "net/http"
//   "github.com/labstack/echo/v4"
  "io/ioutil"
	// "fmt"

	// // "log"
	// "net/http"
	// // "github.com/tidwall/sjson"
)
func CreateScheduler()  {
	// fmt.Printf("## echo : %s" , c)
	
	url := "http://101.79.4.15:32527/yaml"
	method := "POST"
  
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("callbackUrl", "http://192.168.150.197:8101")
	_ = writer.WriteField("requestId", "request01")
	_ = writer.WriteField("yaml", "apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: nginx-deployment\n  labels:\n    app: nginx\nspec:\n  replicas: 3\n  selector:\n    matchLabels:\n      app: nginx\n  template:\n    metadata:\n      labels:\n        app: nginx\n    spec:\n      containers:\n      - name: nginx\n        image: nginx:1.14.2\n        ports:\n        - containerPort: 80")
	err := writer.Close()
	if err != nil {
	  fmt.Println(err)
	  return
	}
   
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return 
	}
	// req.Header.Add("Cookie", "JSESSIONID=568D2934F435E811000015C1963C4962")
  
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}

	fmt.Println(string(body))
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"data": string(body),
	// })
}