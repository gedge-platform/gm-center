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
		// findPath 입력이 있을 경우
		data = Filter(data, findPath)
	} else {
		// findPath 가 "" 일 경우

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

func StringToInterface(i string) interface{} {
	var x interface{}
	if err := json.Unmarshal([]byte(i), &x); err != nil {
		fmt.Printf("Error : %s\n", err)
	}
	return x
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
			log.Println("endPoints 뿌려주기 : ", PodData)

			splits := strings.SplitN(InterfaceToString(FindData(data, "subsets.#.addresses.0", "targetRef.name")), "-", 3)
			log.Printf("Endpoints [%s] Data is %s \n", params.Name, data)
			log.Printf("splits %s \n", splits)

			// TODO: splits 가 여러개 일 수 있으니, 수정 필요(맨 마지막 - 이후로 제거)
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
		log.Println("[#5] data is ", data)

		// params.Kind = "deployments"
		// params.Name = InterfaceToString(FindData(data, "metadata", "uid"))

		// if data, err := GetModel(params); err != nil {
		// 	return nil, err
		// } else {
		// 	log.Println("data is ", data)
		// }

	case "jobs":
		params.Kind = "pods"
		params.Name = ""
		jobName := InterfaceToString(FindData(data, "metadata", "name"))
		uid := InterfaceToString(FindData(data, "metadata", "uid"))
		log.Println("PARAMS.NAEMData12 : ", params.Name)
		log.Println("jobName Data12 : ", jobName)
		log.Println("uid Data12 : ", uid)
		podData, err := GetModel(params)
		if err != nil {
			return nil, err
		}
		podRefer, err := FindDataArr(podData, "items", "job-name", jobName)
		if err != nil {
			return nil, err
		}
		log.Println("jobRefer Data12 : ", podRefer)

		params.Kind = "events"
		params.Name = ""

		eventsData, err := GetModel(params)
		if err != nil {
			return nil, err
		}

		eventRefer, err := FindDataArr(eventsData, "items", "uid", uid)
		if err != nil {
			return nil, err
		}
		var EventInfo []model.EVENT
		EventData := eventRefer
		// log.Printf("# 확인 ", EventData)
		Transcode(EventData, &EventInfo)

		var PodsInfo []model.ReferPodList
		PodsData := podRefer
		// log.Printf("# pod list확인 ", podRefer)
		Transcode(PodsData, &PodsInfo)
		// log.Printf("# job list확인22 ", PodsInfo)

		ReferDataJob := model.ReferDataJob{
			ReferPodList: PodsInfo,
			Event:        EventInfo,
		}
		return ReferDataJob, nil

	case "cronjobs":
		params.Kind = "jobs"
		params.Name = ""
		cronjobName := InterfaceToString(FindData(data, "metadata", "name"))
		uid := InterfaceToString(FindData(data, "metadata", "uid"))

		log.Println("PARAMS.NAEMData12 : ", params.Name)
		log.Println("jobName Data12 : ", cronjobName)
		jobData, err := GetModel(params)
		if err != nil {
			return nil, err
		}
		jobRefer, err := FindDataArr(jobData, "items", "name", cronjobName)
		log.Printf("[###123]jobData", jobData)
		if err != nil {
			return nil, err
		}
		log.Println("jobRefer Data12 : ", jobRefer)

		params.Kind = "events"
		params.Name = ""
		test := InterfaceToString(FindData(jobData, "metadata", "name"))
		log.Println("[#11test Data12 : ", test)
		eventsData, err := GetModel(params)
		if err != nil {
			return nil, err
		}

		eventRefer, err := FindDataArr(eventsData, "items", "uid", uid)
		log.Println("[#11eventRefer Data12 : ", eventRefer)
		if err != nil {
			return nil, err
		}
		var EventInfo []model.EVENT
		EventData := eventRefer
		log.Printf("# 확인 ", EventData)
		Transcode(EventData, &EventInfo)

		var JobsInfo []model.JOBList
		JobData := jobRefer
		log.Printf("# job list확인 ", jobRefer)
		Transcode(JobData, &JobsInfo)
		log.Printf("# job list확인22 ", JobsInfo)

		ReferData := model.ReferCronJob{
			JOBList: JobsInfo,
			Event:   EventInfo,
		}

		return &ReferData, nil
	case "pods":
		params.Kind = "deployments"
		params.Name = ""
		labelApp := InterfaceToString(FindData(data, "metadata.labels", "app"))
		uid := InterfaceToString(FindData(data, "metadata", "uid"))

		log.Println("PARAMS.NAEMData12 : ", params.Name)
		log.Println("serviceAccount Data12 : ", labelApp)

		deployData, err := GetModel(params)
		if err != nil {
			return nil, err
		}
		deployRefer, err := FindDataArr(deployData, "items", "app", labelApp)
		log.Printf("[###123]jobData", deployData)
		if err != nil {
			return nil, err
		}
		log.Println("DeployRefer Data12 : ", deployRefer)

		params.Kind = "events"
		eventsData, err := GetModel(params)
		if err != nil {
			return nil, err
		}
		eventRefer, err := FindDataArr(eventsData, "items", "uid", uid)
		log.Println("[#11eventRefer Data12 : ", eventRefer)
		if err != nil {
			return nil, err
		}

		params.Kind = "endpoints"
		serviceData, err := GetModel(params)
		log.Printf("[#222]", serviceData)
		if err != nil {
			return nil, err
		}
		// serviceRefer, err := FindDataArr(serviceData, "items", "uid", uid)
		// log.Println("[#11eventRefer Data12 : ", eventRefer)
		// if err != nil {
		// 	return nil, err
		// }
		var EventInfo []model.EVENT
		EventData := eventRefer
		log.Printf("# 확인 ", EventData)
		Transcode(EventData, &EventInfo)

		var DeployInfo []model.DeployInfo
		DeployData := deployRefer
		Transcode(DeployData, &DeployInfo)

		ReferData := model.ReferDataDeploy{
			DeployInfo: DeployInfo,
			Event:      EventInfo,
		}
		return ReferData, nil
	}

	return nil, errors.New("")

}
