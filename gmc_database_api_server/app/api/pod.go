package api

import (
	"github.com/labstack/echo/v4"
)

func GetPods(c echo.Context) (err error) {
	// pod_name := c.Param("name")
	// workspace_name := c.QueryParam("workspace")
	// project_name := c.QueryParam("project")
	// cluster_name := c.QueryParam("cluster")
	// fmt.Printf("cronjob_name is %s\n, workspace name is %s\n, project name is %s", pod_name, workspace_name, project_name, cluster_name)
	// var podDetail model.PODDETAIL
	// data, err := common.GetModel(c, "pods")
	// if err != nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, err)
	// 	return nil
	// }
	// fmt.Printf("[#55555]data is info %s", data)

	// // if err != nil {
	// // 	common.ErrorMsg(c, http.StatusNotFound, err)
	// // 	return nil
	// // }
	// // fmt.Printf("[#456789]eventdata is info %s", eventdata)
	// pod_uid := common.FilterStr(data, "metadata.uid")
	// fmt.Printf("[#666]pod_uid type is %s\n", pod_uid)
	// getCallEvent(c)
	// // GetClusterEP(C, getURL(c, "pods"))
	// // eventlist := getURL(c, "pods")
	// // eventList := getEvents(c, "events", pod_uid)
	// // fmt.Printf("[#7777]eventlist is %s\n", eventlist)
	// // eventdata, err := common.GetModel(c, "events")
	// // getCallEvent(c, "pods")
	// // fmt.Printf("[#1eventdata]data is info %s", eventdata)
	// // fmt.Printf("[#함수]data is info %s", getCallEvent())
	// name := common.Finding(data, "metadata", "name")
	// kind := common.Finding(data, "metadata", "kind")
	// fmt.Printf("kind is info %s", kind)

	// namespace := common.Finding(data, "metadata", "namespace")
	// startTime := common.FilterStr(data, "metadata.creationTimestamp")
	// status := common.Finding(data, "status", "phase")

	// timeformat, _ := time.Parse("2006-01-02 15:04:05\n", startTime)
	// fmt.Println(timeformat.Format(time.RFC3339))
	// log.Println("[#48888] data is\n", startTime)
	// log.Println("[#48888] data is\n", timeformat)

	// qosClass := common.Finding(data, "status", "qosClass")

	// podIP := common.Finding(data, "status", "podIP")
	// nodeName := common.Finding(data, "spec", "nodeName")

	// log.Println("[#4] data is", name)
	// // ownerReferencesInfo := []model.CONTAINERS{}
	// // ownerReference := model.CONTAINERS{}
	// ownerReferencesInfo := []model.OwnerReference{}

	// ownerReferences := common.FilterStr(data, "metadata.ownerReferences")
	// fmt.Printf("data1 type is %s\n", common.Typeof(ownerReferences))
	// log.Printf("[#523] t is %s\n", ownerReferences)
	// ownerReferences_err := json.Unmarshal([]byte(ownerReferences), &ownerReferencesInfo)
	// if err != nil {
	// 	panic(ownerReferences_err)
	// }

	// log.Printf("[#53] t is %s\n", ownerReferencesInfo)
	// containerspec := []model.PODCONTAINERS{}
	// containers := common.FilterStr(data, "spec.containers")
	// fmt.Printf("data1 type is %s\n", common.Typeof(containers))
	// log.Printf("[#523] t is %s\n", containers)

	// sepcContainer := json.Unmarshal([]byte(containers), &containerspec)
	// if err != nil {
	// 	panic(sepcContainer)
	// }
	// log.Printf("[#53] t is %s\n", containerspec)

	// data3 := common.Finding(data, "spec.containers.volumeMounts", "name")
	// fmt.Printf("readonly type is %s\n", data3)
	// log.Printf("[#789789798] data3 is %s\n", data3)
	// // StringToInt(data3)
	// // fmt.Printf("data3 type change type is %s\n", data3)
	// // if data3 == "true" {
	// // 	data3 = strng1
	// // } else {
	// // 	data3 = 0
	// // }
	// containerStatuses := []model.ContainerStatuses{}
	// ContainerStatuses_info := common.FilterStr(data, "status.containerStatuses")
	// fmt.Printf("###888 envs type is %s\n", ContainerStatuses_info)
	// fmt.Printf("###555 envs type is %s\n", common.Typeof(ContainerStatuses_info))
	// log.Printf("[#798] t is %s\n", ContainerStatuses_info)

	// ContainerStatuses_pod := json.Unmarshal([]byte(ContainerStatuses_info), &containerStatuses)
	// if err != nil {
	// 	panic(ContainerStatuses_pod)
	// }

	// label := make(map[string]string)
	// labels := common.FilterStr(data, "metadata.labels")
	// fmt.Printf("[labels] is %s\n", labels)
	// err_labels := json.Unmarshal([]byte(labels), &label)

	// if err_labels != nil {
	// 	fmt.Printf("Error : %s\n", err_labels)
	// }
	// annotation := make(map[string]string)
	// annotations := common.FilterStr(data, "metadata.annotations")
	// err_annotation := json.Unmarshal([]byte(annotations), &annotation)
	// if err_annotation != nil {
	// 	fmt.Printf("Error : %s\n", err_annotation)
	// }
	// // var ReadyCount string
	// // if ContainerStatuses_info == true {
	// // 	ReadyCount = "1"
	// // } else {
	// // 	ReadyCount = "2"
	// // }

	// podDetail.Name = pod_name
	// podDetail.Workspace = workspace_name
	// podDetail.Project = project_name
	// podDetail.Cluster = cluster_name
	// podDetail.Namespace = namespace
	// podDetail.Status = status
	// podDetail.CreatedAt = startTime
	// podDetail.NodeName = nodeName
	// podDetail.PodIP = podIP
	// podDetail.QosClass = qosClass
	// podDetail.Podcontainers = containerspec
	// podDetail.ContainerStatuses = containerStatuses
	// podDetail.OwnerReference = ownerReferencesInfo
	// podDetail.Lable = label
	// podDetail.Annotations = annotation

	// return c.JSON(http.StatusOK, echo.Map{
	// 	"items": podDetail,
	// 	//  "events": eventList
	// })
	return nil
}

// data, err := common.GetModel(c, "pods")
// if err != nil {
// 	common.ErrorMsg(c, http.StatusNotFound, err)
// 	return nil
// }

// fmt.Printf("data type is %s\n", common.Typeof(data))
// // var podModel model.PODALL
// List := []model.PODALL{}
// podall := model.PODALL{}

// params := model.PARAMS{
// 	Kind:      "pods",
// 	Name:      "",
// 	Cluster:   c.QueryParam("cluster"),
// 	Workspace: c.QueryParam("workspace"),
// 	Project:   c.QueryParam("project"),
// 	Method:    c.Request().Method,
// 	Body:      c.Request().Body,
// }

// getData, err := common.GetModel(params)
// if err != nil {
// 	common.ErrorMsg(c, http.StatusNotFound, err)
// 	return nil
// }
// // cluster_name := c.QueryParam("cluster")
// // name := common.Finding(data, "metadata", "name")
// // data2 := common.Filter(data, "metadata.ownerReferences")
// // fmt.Printf("data2 type is %s\n", common.Typeof(data2))
// // podModel.Name = name
// // podModel.Cluster = cluster_name
// n := gjson.Parse(data)
// k := gjson.Get(data, "items").Array()
// log.Printf("[#44] %s", n)
// // itemsmap := common.FilterStr(data, "items")
// fmt.Printf("itemsmap data is %s\n", k)
// for t, _ := range k {
// 	test := json.Unmarshal([]byte(k[t].String()), &podall)
// 	fmt.Printf("#11111teststet", test)
// 	// data2 := gjson.Get(data, "metadata.name").String()
// 	// fmt.Printf("#11111teststet", data2)
// 	List = append(List, podall)
// }

// 	return c.JSON(http.StatusOK, echo.Map{
// 		// "testData": podall,
// 		// "items": List,
// 		"tset": getData,
// 	})
// }
