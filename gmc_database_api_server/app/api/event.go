package api

import (
	"gmc_database_api_server/app/model"

	"github.com/labstack/echo/v4"
)

func getCallEvent(c echo.Context) []model.Event {
	// var List []model.Event
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
	return nil
}

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
