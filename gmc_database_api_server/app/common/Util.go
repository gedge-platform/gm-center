package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

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

func Filter(i string, path string) map[string]string {
	x := make(map[string]string)
	parse := gjson.Parse(i)
	Dat := parse.Get(path)
	Arr := parse.Get(path).Array()
	len := len(parse.Get(path).Array())

	if len > 0 {
		// list
		// log.Printf("[#36] Arr is %+v\n", Arr)
		err3 := json.Unmarshal([]byte(Arr[0].String()), &x)
		if err3 != nil {
			fmt.Printf("Error : %s\n", err3)
		}
		// fmt.Println("[#35] is ", x)
	} else {
		// not list
		log.Printf("[#36] Arr is %+v\n", Dat)
		err3 := json.Unmarshal([]byte(Dat.String()), &x)
		if err3 != nil {
			fmt.Printf("Error : %s\n", err3)
		}
		// fmt.Println("[#35] is ", x)

	}

	return x
}

func FilterStr(i string, path string) (string, error) {
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
		return Dat.String(), nil
	} else {
		return Dat.String(), nil
	}
}

func Finding(i string, path string, find string) (string, error) {
	parse := gjson.Parse(i)
	Dat := parse.Get(path)
	Arr := parse.Get(path).Array()
	len := len(parse.Get(path).Array())
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

	return ReturnVal, nil
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
