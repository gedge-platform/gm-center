package common

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// type Response struct {
// 	data   string
// 	status int
// }

func HttpRequest(method string, url string, _body io.ReadCloser) (map[string]string, error) {

	client := resty.New()
	client.SetAllowGetMethodPayload(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(1 * time.Minute)
	client.SetHeaders(map[string]string{
		"Access-Control-Allow-Origin": "*",
		"Content-Type":                "application/json; charset=utf-8",
		"Accept":                      "application/json; charset=utf-8",
	})

	var responseString string
	var responseStatus int
	body := ResponseBody(_body)

	switch method {
	case "GET":
		if resp, err := client.R().SetBody([]byte(body)).Get(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}

	case "POST":
		if resp, err := client.R().SetBody([]byte(body)).Post(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	case "PATCH":
		if resp, err := client.R().SetBody([]byte(body)).Patch(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	case "PUT":
		if resp, err := client.R().SetBody([]byte(body)).Put(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	case "DELETE":
		if resp, err := client.R().Delete(url); err != nil {
			panic(err)
		} else {
			responseString = string(resp.Body())
			responseStatus = resp.StatusCode()
		}
	}
	return map[string]string{
		"data":   responseString,
		"status": strconv.Itoa(responseStatus),
	}, nil
}

func ResponseBody(req io.ReadCloser) string {
	var bodyBytes []byte
	if req != nil {
		bodyBytes, _ = ioutil.ReadAll(req)
	}

	// Restore the io.ReadCloser to its original state
	req = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	buf := new(bytes.Buffer)
	buf.ReadFrom(req)
	newStr := buf.String()

	return newStr
}
