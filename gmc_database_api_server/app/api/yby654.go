package api

import (
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/model"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func GetModelList(params model.PARAMS) []string {
	// t := time.Now()
	// start_time := time.Unix(t.Unix(), 0)

	// fmt.Printf("#################params List : %+v\n", params)
	staticKind := params.Kind
	var DataList []string

	if params.Workspace == "" && params.Cluster == "" && params.Project == "" {
		fmt.Println("#################ALL List")
		Clusters := GetAllDBClusters(params)
		for c, _ := range Clusters {
			params.Cluster = Clusters[c].Name
			params.Workspace = Clusters[c].Name
			params.Project = Clusters[c].Name
			params.Kind = "namespaces"
			getData, _ := common.GetModel(params)
			getData0 := common.FindingArray(common.Finding(getData, "items"))
			for p, _ := range getData0 {
				params.Kind = staticKind
				params.Project = (gjson.Get(getData0[p].String(), "metadata.name")).String()
				getItems, _ := common.GetModel(params)
				getList := common.FindingArray(common.Finding(getItems, "items"))
				for k, _ := range getList {
					str := getList[k].String()
					strVal, _ := sjson.Set(str, "clusterName", Clusters[c].Name)
					DataList = append(DataList, strVal)
				}
				// 호출값없을때 예외처리 필요 함 ㅋㅋ

			}
		}
		// t2 := time.Now()
		// end_time := time.Unix(t2.Unix(), 0)
		// fmt.Printf("Time : %s,%s", end_time, start_time)
		return DataList

	} else if params.Workspace == "" && params.Cluster != "" && params.Project == "" {
		fmt.Println("#################Cluster List")
		params.Workspace = params.Cluster
		params.Project = params.Cluster
		params.Kind = "namespaces"
		getData, _ := common.GetModel(params)
		getData0 := common.FindingArray(common.Finding(getData, "items"))
		for p, _ := range getData0 {
			params.Kind = staticKind
			params.Project = (gjson.Get(getData0[p].String(), "metadata.name")).String()
			getItems, _ := common.GetModel(params)
			getList := common.FindingArray(common.Finding(getItems, "items"))
			// fmt.Printf("#################getList : %+v", getList)
			for k, _ := range getList {
				str := getList[k].String()
				strVal, _ := sjson.Set(str, "clusterName", params.Cluster)
				DataList = append(DataList, strVal)
			}
		}
		return DataList
	} else if params.Workspace != "" && params.Cluster == "" && params.Project == "" {
		fmt.Println("#################Workspace List")
		workspace := GetDBWorkspace(params)
		selectCluster := workspace.SelectCluster
		slice := strings.Split(selectCluster, ",")
		for w, _ := range slice {
			params.Cluster = slice[w]
			params.Project = slice[w]
			// fmt.Printf("#################clusterName:%s\n", params.Cluster)
			params.Kind = "namespaces"
			getData, _ := common.GetModel(params)
			getData0 := common.FindingArray(common.Finding(getData, "items"))
			for p, _ := range getData0 {
				params.Kind = staticKind
				params.Project = (gjson.Get(getData0[p].String(), "metadata.name")).String()
				// fmt.Printf("#################params:%s\n", params.Project)
				getItems, _ := common.GetModel(params)
				getList := common.FindingArray(common.Finding(getItems, "items"))
				// fmt.Printf("#################getList : %+v", getList)
				for k, _ := range getList {
					str := getList[k].String()
					strVal, _ := sjson.Set(str, "clusterName", params.Cluster)
					DataList = append(DataList, strVal)
				}
				// if err == nil {
				// }
			}
		}
		return DataList
	} else if params.Project != "" && params.Workspace != "" {
		fmt.Println("#################Project List")
		params.Name = params.Project
		project := GetDBProject(params)
		if project.Type == "user" {
			fmt.Println("#################user project")
			if project.WorkspaceName != params.Workspace {
				msg := common.ErrorMsg2(http.StatusNotFound, common.ErrNotFound)
				DataList = append(DataList, common.InterfaceToString(msg))
				return DataList
			}
			selectCluster := project.SelectCluster
			slice := strings.Split(selectCluster, ",")
			for w, _ := range slice {
				// fmt.Printf("#################Project List : %s", slice[w])
				params.Cluster = slice[w]
				params.Kind = staticKind
				params.Name = ""
				getItems, _ := common.GetModel(params)
				getList := common.FindingArray(common.Finding(getItems, "items"))
				for k, _ := range getList {
					str := getList[k].String()
					strVal, _ := sjson.Set(str, "clusterName", params.Cluster)
					DataList = append(DataList, strVal)
				}

			}
			return DataList
		} else {
			fmt.Println("#################system project")
			workspace := GetDBWorkspace(params)
			selectCluster := workspace.SelectCluster
			slice := strings.Split(selectCluster, ",")
			for w, _ := range slice {
				params.Cluster = slice[w]
				params.Name = ""
				params.Kind = staticKind
				getData, _ := common.GetModel(params)
				getData0 := common.FindingArray(common.Finding(getData, "items"))
				for k, _ := range getData0 {
					str := getData0[k].String()
					strVal, _ := sjson.Set(str, "clusterName", params.Cluster)
					DataList = append(DataList, strVal)
				}
				// DataList = append(DataList, getData0)
			}
		}
		return DataList
	}
	return nil
}
