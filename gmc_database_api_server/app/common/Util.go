package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"gmc_database_api_server/app/model"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func SearchNested(obj interface{}, key string) (interface{}, bool) {
	log.Println("obj is ", obj)

	switch t := obj.(type) {
	case map[string]interface{}:
		if v, ok := t[key]; ok {
			return v, ok
		}
		for _, v := range t {
			if result, ok := SearchNested(v, key); ok {
				return result, ok
			}
		}
	case []interface{}:
		for _, v := range t {
			if result, ok := SearchNested(v, key); ok {
				return result, ok
			}
		}
	}

	// key not found
	return nil, false
}

func CreateKeyValuePairs(m map[string]string) []string {
	var r []string
	for key, value := range m {
		val := key + "=" + "\"" + value + "\""
		r = append(r, val)
	}
	return r
}

func ArrStringtoBytes(i []string) []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(i)
	bs := buf.Bytes()
	return bs
}

func FindData(data string, findPath, findValue string) interface{} {

	log.Println("FindPath is ", findPath)
	findPathCheck := strings.Compare(findPath, "") != 0
	if findPathCheck {
		// findkey 입력이 있을 경우
		data = Filter(data, findPath)
	} else {
		// findkey 가 "" 일 경우

	}

	log.Println("findValue is ", findValue)
	findValueCheck := strings.Compare(findValue, "") != 0
	if findValueCheck {
		// findValue 입력이 있을 경우
		data = Finding(data, findValue)

	} else {
		// findValue 가 "" 일 경우
	}

	log.Println("최종 data is : ", data)
	fmt.Println("type:", reflect.ValueOf(data).Type())

	var x interface{}
	if err := json.Unmarshal([]byte(data), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
		x = data
		return x
	}

	return x
}

func FindDataStr(data string, findPath, findValue string) string {

	log.Println("FindPath is ", findPath)
	findPathCheck := strings.Compare(findPath, "") != 0
	if findPathCheck {
		// findkey 입력이 있을 경우
		data = Filter(data, findPath)
	} else {
		// findkey 가 "" 일 경우

	}

	log.Println("findValue is ", findValue)
	findValueCheck := strings.Compare(findValue, "") != 0
	if findValueCheck {
		// findValue 입력이 있을 경우
		data = Finding(data, findValue)

	} else {
		// findValue 가 "" 일 경우
	}

	log.Println("최종 data is : ", data)
	fmt.Println("type:", reflect.ValueOf(data).Type())

	return data
}

// func FindDataArr(data string, findPath, uniq string) string {

// 	log.Println("FindPath is ", findPath)
// 	log.Println("findValue is ", compareValue)

// 	findValueCheck := strings.Compare(compareValue, "") != 0
// 	if findValueCheck {
// 		// findValue 입력이 있을 경우
// 		data = Finding(data, compareValue)

// 	} else {
// 		// findValue 가 "" 일 경우
// 	}

// 	log.Println("최종 data is : ", data)
// 	fmt.Println("type:", reflect.ValueOf(data).Type())

// 	return data
// }

// func FilterUniq(i string, uniq string) string {
// 	parse := gjson.Parse(i)
// 	Dat := parse.Get(path)
// 	// Arr := parse.Get(path).Array()
// 	len := len(parse.Get(path).Array())

// 	if len > 0 {
// 		// list
// 		// log.Printf("[#36] Arr is %+v\n", Arr)
// 		// err3 := json.Unmarshal([]byte(Arr[0].String()), &x)
// 		// if err3 != nil {
// 		// 	fmt.Printf("Error : %s\n", err3)
// 		// }
// 		// fmt.Println("[#35] is ", x)
// 		return Dat.String()
// 	} else {
// 		return Dat.String()
// 	}
// }

func Filter(i string, path string) string {
	parse := gjson.Parse(i)
	Dat := parse.Get(path)
	// Arr := parse.Get(path).Array()
	len := len(parse.Get(path).Array())

	if len > 0 {
		// list
		// log.Printf("[#36] Arr is %+v\n", Arr)
		// err3 := json.Unmarshal([]byte(Arr[0].String()), &x)
		// if err3 != nil {
		// 	fmt.Printf("Error : %s\n", err3)
		// }
		// fmt.Println("[#35] is ", x)
		return Dat.String()
	} else {
		return Dat.String()
	}
}

func Finding(i string, find string) string {
	parse := gjson.Parse(i)
	Dat := parse
	Arr := parse.Array()
	len := len(parse.Array())
	ReturnVal := ""

	if len > 0 {
		// list
		for n, _ := range Arr {
			ReturnVal = Arr[n].Get(find).String()
		}
	} else {
		// not list
		ReturnVal = Dat.Get(find).String()
	}

	return ReturnVal
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func Transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}

func InterfaceToString(i interface{}) string {
	str := fmt.Sprintf("%v", i)
	return str
}

func StringToMapInterface(i string) map[string]interface{} {
	x := make(map[string]interface{})
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
}

func InterfaceToTime(i interface{}) time.Time {
	createTime := InterfaceToString(i)
	timer, _ := time.Parse(time.RFC3339, createTime)
	return timer
}
func StringToInterface(i string) interface{} {
	var x interface{}
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
}
func GetModelRelatedList(params model.PARAMS) (interface{}, error) {

	data, err := GetModel(params)
	if err != nil {
		return "", err
	}

	switch strings.ToLower(params.Kind) {
	case "services":
		params.Kind = "endpoints"
		params.Name = InterfaceToString(FindData(data, "metadata", "name"))

		if data, err := GetModel(params); err != nil {
			return nil, err
		} else {
			PodData := FindData(data, "subsets.#.addresses", "")
			splits := strings.SplitN(InterfaceToString(FindData(data, "subsets.#.addresses.0", "targetRef.name")), "-", 3)
			log.Printf("Endpoints [%s] Data is %s \n", params.Name, data)
			log.Printf("splits %s \n", splits)
			podName := splits[0] + "-" + splits[1]

			params.Kind = "replicasets"
			params.Name = podName

			if data, err := GetModel(params); err != nil {
				return nil, err
			} else {
				log.Printf("replicasets [%s] Data is %s \n", params.Name, data)
				params.Kind = "deployments"
				params.Name = FindDataStr(data, "metadata.ownerReferences.0", "name")

				if data, err := GetModel(params); err != nil {
					return nil, err
				} else {
					var deployModel model.Deployment
					Transcode(data, &deployModel)

					services := model.SERVICELISTS{
						Pods: PodData,
						Deployments: model.SERVICEDEPLOYMENT{
							Name:     InterfaceToString(FindData(data, "metadata", "name")),
							UpdateAt: InterfaceToTime(FindData(data, "status.conditions", "lastUpdateTime")),
						},
					}
					return services, nil
				}
			}
		}

	case "deployments":
	// params2.Kind = "endpoints"
	// params2.Name = InterfaceToString(FindData(data, "metadata", "name"))

	case "jobs":
		params.Kind = "pods"
		uniq := InterfaceToString(FindData(data, "metadata", "name"))

		log.Println("PARAMS.NAEMData12 : ", params.Name)
		log.Println("uniq Data12 : ", uniq)
		if data, err := GetModel(params); err != nil {
			return nil, err
		} else {
			PodDataA := FindData(data, "metadata", "items")
			params.Name = InterfaceToString(FindData(data, "metadata", "name"))
			log.Println("Pod List Data12 : ", data)
			log.Println("Pod List 데이터뿌려주기  : ", PodDataA)
			log.Println("Pod params Data125 : ", params.Name)
			if data, err := GetModel(params); err != nil {
				return nil, err
			} else {
				var Only interface{}
				var List []interface{}
				n := gjson.Parse(data)
				// log.Println("all gjson parse List Data124 : ", n)
				k := n.Get("items").Array()
				log.Println("onlygjson parse List Data124 : ", k)
				result := n.Get("items.#.metadata").Array()
				result2 := n.Get("items.#.spec").Array()
				fmt.Println(result2)
				for num, _ := range k {

					// r2esult := n.Get(`items.#.metadata.labels.job-name`).Array()
					log.Printf("####135555confirm ", result[num].Get(`labels.job-name`))
					arr := k[num].Get(`metadata`).Get(`labels.job-name`).String()
					log.Printf("###arrr ", arr)

					if strings.Contains(arr, uniq) == true {
						fmt.Println(arr)
						// if k[num].Get(`metadata`).Get(`labels.job-name`).String() == uniq {
						// if k[num].Get(`metadata.labels.job-name`).String() == uniq {
						log.Printf("metadata ", k[num].Get(`metadata`).Get(`labels.job-name`).String())
						log.Printf("DATA CONFIRM")

						// err := json.Unmarshal([]byte(result[num].String()), &Only)
						// fmt.Println(err)
						test := k[num].String()
						// err := json.Unmarshal([]byte(test))
						fmt.Println("eeeeee", test)
						// fmt.Println("eeeeeeFFFF", err)
						// err2 := StringToInterface(test), &Only

						err := json.Unmarshal([]byte(test), &Only)
						if err != nil {
							fmt.Println("[!53] error")
						}
						List = append(List, Only)

						// fmt.Println(err2)
						// log.Printf("err confirm ", err)
						// log.Printf("List ", Only)
						// if err != nil {
						// 	panic(err)
						// }
						// List = append(List, Only)
						return List, nil
					}
				}
			}
		}
	}
	return nil, errors.New("")
}
