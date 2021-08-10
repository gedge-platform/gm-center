package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

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

func FilterbyUID(i string, kind string, findValue string) []string {
	var multiStr []string

	parse := gjson.Parse(i)
	k := parse.Get("items").Array()
	for num, _ := range k {
		arr := k[num].Get("metadata.ownerReferences").Array()
		if len(arr) > 0 {
			for num2, _ := range arr {
				if strings.Contains(arr[num2].Get("kind").String(), kind) == true && strings.Contains(arr[num2].Get("uid").String(), findValue) == true {

					log.Println("[#4] data : ", k[num].String())
					multiStr = append(multiStr, k[num].String())
					// err := json.Unmarshal([]byte(k[num].String()), &jobOnly)
					// if err != nil {
					// 	panic(err)
					// }
					// jobList = append(jobList, jobOnly)
				}
			}
		}
	}

	return multiStr
}

func FilterbyName(i string, kind string, findValue string) []string {
	var multiStr []string

	iData := FindData(i, "", "")

	log.Println("iData is ", iData)
	log.Println("findValue is ", findValue)

	if v, ok := SearchNested(iData, findValue); ok {
		// fmt.Printf("[#4] %+v\n", v)
		fmt.Println("hello [34] ", strings.Contains(InterfaceToString(v), findValue))
		fmt.Println("Key found")
	} else {
		fmt.Println("Key not found")
	}

	parse := gjson.Parse(i)
	k := parse.Get("items").Array()
	for num, _ := range k {
		arr := k[num].Get("metadata.ownerReferences").Array()
		if len(arr) > 0 {
			for num2, _ := range arr {
				if strings.Contains(arr[num2].Get("kind").String(), kind) == true && strings.Contains(arr[num2].Get("uid").String(), findValue) == true {

					log.Println("[#4] data : ", k[num].String())
					multiStr = append(multiStr, k[num].String())
					// err := json.Unmarshal([]byte(k[num].String()), &jobOnly)
					// if err != nil {
					// 	panic(err)
					// }
					// jobList = append(jobList, jobOnly)
				}
			}
		}
	}

	return multiStr
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
