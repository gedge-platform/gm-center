package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gmc_api_gateway/config"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

var (
	// common
	ErrNoData            = errors.New("No Data")
	ErrNotFound          = errors.New("Not Found")
	ErrClusterNotFound   = errors.New("Cluster Not Found")
	ErrWorkspaceNotFound = errors.New("Workspace Not Found")
	ErrWorkspaceInvalid  = errors.New("Workspace Empty")
	ErrProjectInvalid    = errors.New("Project Empty")
	ErrDetailNameInvalid = errors.New("Detail Name Empty")
	ErrClusterInvalid    = errors.New("Required Cluster Name")
	// Account
	ErrIdInvalid = errors.New("id is empty")
	ErrBodyEmpty = errors.New("Body is empty")
)

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
}

// Return Error Message
func ErrorMsg(c echo.Context, status int, err error) {
	errMsg := messageFormat{
		StatusCode: status,
		Message:    err.Error(),
	}
	c.JSON(status, echo.Map{"error": errMsg})
}
func SearchNestedValue(obj interface{}, key string, uniq string) (interface{}, bool) {
	// log.Println("obj is ", obj)

	switch t := obj.(type) {
	case map[string]interface{}:
		if v, ok := t[key]; ok {
			if strings.Compare(InterfaceToString(v), uniq) == 0 {
				log.Printf("[#1] sfsdfs v is %s, uniq is %s, ok is %t", v, uniq, ok)
				return v, ok
			}
		}
		for _, v := range t {
			if result, ok := SearchNestedValue(v, key, uniq); ok {
				return result, ok
			}
		}
	case []interface{}:
		for _, v := range t {
			if result, ok := SearchNestedValue(v, key, uniq); ok {
				return result, ok
			}
		}
	}

	// key not found
	return nil, false
}

// Environment Value ("LISTEN_PORT")
func GetListenPort(config *config.Config) string {
	port := os.Getenv("LISTEN_PORT")

	if len(port) == 0 {
		port = config.COMMON.Port
	}
	intPort, err := strconv.Atoi(port)
	if err != nil || intPort < 1 || 65535 < intPort {
		port = config.COMMON.Port
	}

	return ":" + port
}

// String to Uint
func ConvertStringToUint(data string) (uintData uint, err error) {
	u64, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return
	}

	return uint(u64), nil
}

// String to Int
func ConvertStringToInt(data string) (uintData int, err error) {
	u64, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return
	}

	return int(u64), nil
}

// String to uuid.UUID
func ConvertStringToUuid(data string) (uuidData uuid.UUID, err error) {
	uuid, err := uuid.FromString(data)
	if err != nil {
		return uuid, errors.New("failed to parse UUID")
	}

	return uuid, nil
}

func StringToMapInterface(i string) map[string]interface{} {
	x := make(map[string]interface{})
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
}

func InterfaceToString(i interface{}) string {
	str := fmt.Sprintf("%v", i)
	return str
}

func StringToInterface(i string) interface{} {
	var x interface{}
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
}

// Find key in interface (recursively) and return value as interface
func Find(obj interface{}, key string) (interface{}, bool) {

	//if the argument is not a map, ignore it
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}

	for k, v := range mobj {
		// key match, return value
		if k == key {
			return v, true
		}

		// if the value is a map, search recursively
		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := Find(m, key); ok {
				return res, true
			}
		}
		// if the value is an array, search recursively
		// from each element
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := Find(a, key); ok {
					return res, true
				}
			}
		}
	}

	// element not found
	return nil, false
}

func InterfaceToArray(i interface{}) []string {
	str := reflect.ValueOf(i)
	var x []string
	for j := 0; j < str.Len(); j++ {
		data := InterfaceToString(str.Index(j))
		x = append(x, data)
	}
	return x
}

func FindDataArr(i interface{}, p, f, u string) (interface{}, error) {
	log.Println("[In #FindDataArr]")
	log.Println("[#1] Data is ", i)
	log.Println("[#2] find path string is ", p)
	log.Println("[#2] find key string is ", f)
	log.Println("[#3] uniq string is ", u)

	// var itemCheck bool
	var parse, data gjson.Result
	var arr []gjson.Result
	var result interface{}
	var results []interface{}
	ia := InterfaceToString(i)

	parse = gjson.Parse(ia)
	log.Println("[#4] Parse is ", parse)

	pathCheck := strings.Compare(p, "") != 0
	// itemCheck = len(parse.Get("items").Array()) > 0
	// log.Println("[#4] itemCheck is ", itemCheck)

	if pathCheck {
		data = parse.Get(p)
		log.Println("[#5] filter data is ", data)
	} else {
		data = parse
		log.Println("[#5] filter data is ", data)
	}

	len := len(data.Array())
	log.Println("[#6] len(data) is ", len)

	if len > 0 {
		// list
		arr = data.Array()
		log.Println("[#7-1] len > 0, list")
		for t, _ := range arr {

			dataInterface := StringToMapInterface(arr[t].String())

			if v, ok := SearchNestedValue(dataInterface, f, u); ok {
				fmt.Printf("Arr[%d] Found it ! \n", t)
				fmt.Printf("Unique is : %+v\n", v)
				fmt.Printf("data is %s\n", arr[t])
				err := json.Unmarshal([]byte(arr[t].String()), &result)
				if err != nil {
					fmt.Println("[!53] error")
				}
				results = append(results, result)
				fmt.Printf("[%d] result Data is %s", t, results)
			} else {
				fmt.Printf("Arr[%d] Key not found\n", t)
			}
		}

		if len == 1 {
			log.Println("[#7-2] len == 1, list")
			dataInterface := StringToInterface(arr[0].String())
			if v, ok := SearchNestedValue(dataInterface, f, u); ok {
				fmt.Println("Found it !")
				fmt.Printf("Unique is : %+v\n", v)
				return StringToInterface(arr[0].String()), nil
			} else {
				return nil, nil
			}
		}

		// list 출력
		return results, nil

	} else {
		return StringToInterface(data.String()), nil
	}
}
