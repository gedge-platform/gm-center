package api

import (
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"

	"github.com/tidwall/gjson"
)

func getCallEvent(params model.PARAMS) []model.EVENT {
	staticKind := params.Kind
	var List []model.EVENT
	searchName := params.Name
	params.Name = ""
	params.Kind = "events"
	fmt.Printf("[#]staticKind : %s\n", staticKind)
	getData, _ := common.GetModel(params)
	fmt.Printf("[###]getData : %+v\n", getData)
	if staticKind == "nodes" || staticKind == "namespaces" {
		params.Cluster = params.Name
		events := common.FindingArray(common.Finding(getData, "items"))
		for e, _ := range events {
			event := model.EVENT{
				Kind:      (gjson.Get(events[e].String(), "regarding.name")).String(),
				Name:      (gjson.Get(events[e].String(), "metadata.name")).String(),
				Namespace: (gjson.Get(events[e].String(), "metadata.namespace")).String(),
				Message:   (gjson.Get(events[e].String(), "note")).String(),
				Reason:    (gjson.Get(events[e].String(), "reason")).String(),
				Type:      (gjson.Get(events[e].String(), "type")).String(),
				EventTime: (gjson.Get(events[e].String(), "metadata.creationTimestamp")).Time(),
			}
			List = append(List, event)
		}
		return List
	} else {
		fmt.Printf("[#]searchName : %s\n", searchName)
		events, err := common.FindDataArrStr2(getData, "items", "name", searchName)
		if err != nil {
			return nil
		}
		fmt.Printf("[#]events : %s\n", events)

		for e, _ := range events {
			event := model.EVENT{
				Kind:      common.InterfaceToString(common.FindData(events[e], "regarding", "kind")),
				Name:      common.InterfaceToString(common.FindData(events[e], "metadata", "name")),
				Namespace: common.InterfaceToString(common.FindData(events[e], "metadata", "namespace")),
				Message:   common.InterfaceToString(common.FindData(events[e], "note", "")),
				Reason:    common.InterfaceToString(common.FindData(events[e], "reason", "")),
				Type:      common.InterfaceToString(common.FindData(events[e], "type", "")),
				EventTime: common.InterfaceToTime(common.FindData(events[e], "metadata", "creationTimestamp")),
			}
			List = append(List, event)
		}
		return List
	}
}

// func getCallEvent(c echo.Context) []model.Event {
// // var Only model.Event
// poddata, err := common.GetModel(c, "pods")
// if err != nil {
// 	common.ErrorMsg(c, http.StatusNotFound, err)
// 	// return nil
// }
// eventdata, err := common.GetModel(c, "events")
// if err != nil {
// 	common.ErrorMsg(c, http.StatusNotFound, err)
// 	return nil
// }
// // if eventdata := getData(c, url, false); eventdata != "nf" {
// n := gjson.Parse(eventdata)
// k := n.Get("items").Array()
// log.Printf("[#44] %s", n)
// log.Printf("[#45] %s", k)
// // log.Printf("[#46] %s", Only)
// // for t, _ := range k {
// // 	if k[t].Get(`regarding.uid`).String() == uniq {
// // 		log.Printf("same uid %s is %s", k[t].Get(`regarding.uid`).String(), uniq)
// // 		log.Printf("[#45] %s", k[t])
// // 		err := json.Unmarshal([]byte(k[t].String()), &Only)
// // 		if err != nil {
// // 			panic(err)
// // 		}
// // 		List = append(List, Only)
// // 	}
// // }
// // }
// log.Println("[#77777 eventdata] data is", eventdata)
// podUid := common.Finding(poddata, "metadata", "uid")
// log.Println("[#88888getcallevent] data is", podUid)

// return List

// List := []model.EVENT{}
// Only := model.EVENT{}
// url := getURL(c, kind)

// log.Printf("[#43] %s", url)

// if data := getData(c, url, false); data != "nf" {
// 	n := gjson.Parse(data)
// 	k := n.Get("items").Array()
// 	log.Printf("[#44] %s", n)
// 	log.Printf("[#45] %s", k)
// 	log.Printf("[#46] %s", Only)
// 	// for t, _ := range k {
// 	// 	if k[t].Get(`regarding.uid`).String() == uniq {
// 	// 		log.Printf("same uid %s is %s", k[t].Get(`regarding.uid`).String(), uniq)
// 	// 		log.Printf("[#45] %s", k[t])
// 	// 		err := json.Unmarshal([]byte(k[t].String()), &Only)
// 	// 		if err != nil {
// 	// 			panic(err)
// 	// 		}
// 	// 		List = append(List, Only)
// 	// 	}
// 	// }
// }
// return List
// }
