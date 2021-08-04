package api

import (
	"context"
	"fmt"
	"gmc_database_api_server/app/common"
	"gmc_database_api_server/app/db"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"

	// "github.com/prometheus/common/config"
	"github.com/labstack/echo/v4"
)

var Mtemplates = map[string]string{
	"query_apiPrefix":      "/api/v1/query?query=",
	"queryrange_apiPrefix": "/api/v1/query_range?query=",

	"clusters":                   "sum by (cluster) (kube_node_info)",
	"cluster_cpu_util":           "round(100-(avg(irate(node_cpu_seconds_total{mode='idle', $1}[5m]))by(cluster)*100),0.1)",
	"cluster_cpu_usage":          "round(sum(rate(container_cpu_usage_seconds_total{id='/', $1}[2m]))by(cluster),0.01)",
	"cluster_cpu_core_total":     "sum(machine_cpu_cores{$1})by(cluster)",
	"cluster_memory_util":        "round(sum(node_memory_MemTotal_bytes{$1}-node_memory_MemFree_bytes-node_memory_Buffers_bytes-node_memory_Cached_bytes-node_memory_SReclaimable_bytes)by(cluster)/sum(node_memory_MemTotal_bytes)by(cluster)*100,0.1)",
	"cluster_memory_usage":       "round(sum(node_memory_MemTotal_bytes{$1}-node_memory_MemFree_bytes-node_memory_Buffers_bytes-node_memory_Cached_bytes-node_memory_SReclaimable_bytes)by(cluster)/1024/1024/1024,0.01)",
	"cluster_memory_total":       "round(sum(node_memory_MemTotal_bytes{$1})by(cluster)/1024/1024/1024,0.01)",
	"cluster_disk_util":          "round(((sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)-sum(node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster))/(sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)))*100,0.1)",
	"cluster_disk_usage":         "round((sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)-sum(node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster))/1000/1000/1000,0.01)",
	"cluster_disk_capacity":      "round(sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)/1000/1000/1000,0.01)",
	"cluster_pod_running":        "count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(cluster,pod))by(cluster)",
	"cluster_pod_quota":          "sum(max(kube_node_status_capacity{resource='pods', $1})by(node,cluster)unless on(node,cluster)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0))by(cluster)",
	"cluster_pod_util":           "round((count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(cluster,pod))by(cluster))/(sum(max(kube_node_status_capacity{resource='pods', $1})by(node,cluster)unless on(node,cluster)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0))by(cluster))*100,0.1)",
	"pod_cpu":                    "round(sum(irate(container_cpu_usage_seconds_total{job='kubelet',pod!='',image!='', $1}[5m]))by(namespace,pod,cluster)*1000,0.001)",
	"pod_memory":                 "sum(container_memory_usage_bytes{job='kubelet',pod!='',image!='', $1})by(namespace,pod,cluster)",
	"pod_net_bytes_transmitted":  "round(sum(irate(container_network_transmit_bytes_total{pod!='',interface!~'^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)',job='kubelet', $1}[5m]))by(namespace,pod,cluster)/125,0.01)",
	"pod_net_bytes_received":     "round(sum(irate(container_network_receive_bytes_total{pod!='',interface!~'^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)',job='kubelet', $1}[5m]))by(namespace,pod,cluster)/125,0.01)",
	"node_cpu_utilisation":       "100-(avg(irate(node_cpu_seconds_total{mode='idle', $1}[5m]))by(instance)*100)",
	"node_cpu_usage":             "sum(rate(container_cpu_usage_seconds_total{id='/'}[5m]))by(node)",
	"node_cpu_total":             "sum(machine_cpu_cores)by(node)",
	"node_memory_utilisation":    "(node_memory_MemTotal_bytes-node_memory_MemAvailable_bytes)/node_memory_MemTotal_bytes",
	"node_memory_usage":          "node_memory_MemTotal_bytes-node_memory_MemFree_bytes-node_memory_Buffers_bytes-node_memory_Cached_bytes-node_memory_SReclaimable_bytes",
	"node_memory_total":          "sum(node_memory_MemTotal_bytes)by(instance)",
	"node_disk_size_utilisation": "100-((node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs'} * 100)/node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs'})",
	"node_disk_size_usage":       "(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs'}-(node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs'}))",
	"node_disk_size_capacity":    "/api/v1/query_range?query(node_filesystem_size_bytes{mountpoint='/'',fstype!='rootfs'})",
	// node_pod_utilisation/{cluster_name} "sum(kubelet_running_pods)by(node)/(max(kube_node_status_capacity%7Bcluster='{cluster_name}',resource='pods'%7D)by(node)unless%20on(node)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0))*100"
	"node_pod_running_count":            "sum(kubelet_running_pods)by(node)",
	"node_pod_quota":                    "max(kube_node_status_capacity{resource='pods'})by(node)unless on(node)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0)",
	"node_disk_inode_utilisation":       "100-(node_filesystem_files_free{mountpoint='/'}/node_filesystem_files{mountpoint='/'}*100)",
	"node_disk_inode_total":             "node_filesystem_files{mountpoint='/'}",
	"node_disk_inode_usage":             "node_filesystem_files{mountpoint='/'}-node_filesystem_files_free{mountpoint='/'}",
	"node_disk_read_iops":               "rate(node_disk_reads_completed_total[5m])",
	"node_disk_write_iops":              "rate(node_disk_writes_completed_total[5m])",
	"node_disk_read_throughput":         "irate(node_disk_read_bytes_total[5m])",
	"node_disk_write_throughput":        "irate(node_disk_written_bytes_total[5m])",
	"node_net_bytes_transmitted":        "irate(node_network_transmit_bytes_total{device='ens3'}[5m])",
	"node_net_bytes_received":           "irate(node_network_receive_bytes_total{device='ens3'}[5m])",
	"apiserver_request_rate":            "round(sum(irate(apiserver_request_total{$1}[5m]))by(cluster),0.001)",
	"scheduler_schedule_attempts_total": "scheduler_pod_scheduling_attempts_count{$1}",
	"scheduler_schedule_fail":           "sum(rate(scheduler_pending_pods{$1}[5m]))by(cluster)",
	"scheduler_schedule_fail_total":     "sum(scheduler_pending_pods{$1})by(cluster)",
	"namespace_cpu":                     "round(sum(irate(container_cpu_usage_seconds_total{job='kubelet',pod!='',image!='', $1}[5m]))by(namespace,pod,cluster),0.001)",
	"namespace_memory":                  "sum(sum(container_memory_usage_bytes{job='kubelet',pod!='',image!='', $1})by(namespace,pod,cluster))by(namespace,cluster)",
	"namespace_pod_count":               "count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(pod,cluster,namespace))by(cluster,namespace)",
}

func Metrics(c echo.Context) (err error) {

	kind := c.Param("kind")
	name := c.Param("name")
	cluster_name := c.QueryParam("cluster")
	clusterAll := false

	fmt.Printf("%s", Mtemplates)
	if kind != "" {
		log.Printf("kind exist : %s", kind)
		if strings.Compare(cluster_name, "all") == 0 {
			log.Println("cluster 'all'")
			clusterAll = true
		}

		switch clusterAll {
		case true:
			getMonitoring("all", name, c)
		case false:
			getMonitoring(cluster_name, name, c)
		}

	}

	return nil
}

func Query(endpointAddr string) model.Value {

	client, err := api.NewClient(api.Config{
		Address: endpointAddr,
	})
	if err != nil {
		log.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, "up", time.Now())
	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	// log.Printf("Result:\n%v\n", result)
	return result
}

func QueryRange(endpointAddr string, query string, c echo.Context) model.Value {
	log.Println("queryrange in")

	var start_time time.Time
	var end_time time.Time
	var step time.Duration

	tm, _ := strconv.ParseInt(c.QueryParam("start"), 10, 64)
	start_time = time.Unix(tm, 0)
	log.Println(start_time)

	tm2, _ := strconv.ParseInt(c.QueryParam("end"), 10, 64)
	end_time = time.Unix(tm2, 0)
	log.Println(end_time)

	tm3, _ := time.ParseDuration(c.QueryParam("step"))
	step = tm3
	log.Println(step)

	client, err := api.NewClient(api.Config{
		Address: endpointAddr,
	})
	if err != nil {
		log.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}
	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: start_time,
		End:   end_time,
		Step:  step,
	}
	result, warnings, err := v1api.QueryRange(ctx, query, r)
	if err != nil {
		log.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	// log.Printf("Result:\n%v\n", result)
	return result
}

func GetDuration(c echo.Context) int64 {
	t, _ := time.ParseDuration(c.QueryParam("step"))
	log.Printf("#4d - %s", t)
	returnVal := int64(t / time.Second)
	log.Printf("#5d - %t", returnVal)
	return returnVal
}

func clusterExpr(val string, clusterName string) string {
	var returnVal string
	if clusterName != "" {
		returnVal = fmt.Sprintf(`cluster="%s"`, clusterName)
	} else {
		returnVal = fmt.Sprintf(`cluster!=""`)
	}
	return strings.Replace(val, "$1", returnVal, -1)
}

func checkQueryParams(c echo.Context) error {
	if c.QueryParam("cluster") == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "cluster is empty", "how-to": "cluster={cluster_name}", "value": "all, {cluster_name}"})
	}
	if c.QueryParam("start") == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "start is empty", "how-to": "start={start_time}"})
	}
	if c.QueryParam("end") == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "end is empty", "how-to": "end={end_time}"})
	}
	if c.QueryParam("step") == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error": "step is empty", "how-to": "step={step_time}"})
	}
	return nil
}

func getMonitoring(cluster string, name string, c echo.Context) error {
	// cluster - 'all' or '{cluster_name}'
	// name - /cpu, /memory, /disk 등
	db := db.DbManager()
	addr := "http://192.168.150.115:31298"

	start := c.QueryParam("start")
	end := c.QueryParam("end")
	step := c.QueryParam("step")

	log.Printf("[getMonitoring in]\ncluster : %s\nname : %s", cluster, name)
	log.Printf("cluster %s\n start %s\n end %s\n step %s", cluster, start, end, step)

	if err := checkQueryParams(c); err != nil {
		log.Println("error")
	}

	log.Printf("cluster name is : %s", cluster)
	models := FindClusterDB(db, "Name", cluster)
	if models == nil {
		common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
	} else {
		log.Println("models find it !")
	}

	log.Println("Select Cluster")
	log.Printf("cluster : %s, name : %s", cluster, name)
	
	switch name {
	case "cpu":
		cpu_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_cpu_util"], cluster), c)
		cpu_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_cpu_usage"], cluster), c)
		cpu_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_cpu_core_total"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"cpu": echo.Map{
				"cpu_util":  cpu_util,
				"cpu_usage": cpu_usage,
				"cpu_total": cpu_total,
			},
		})
	case "memory":
		memory_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_memory_util"], cluster), c)
		memory_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_memory_usage"], cluster), c)
		memory_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_memory_total"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"memory": echo.Map{
				"memory_util":  memory_util,
				"memory_usage": memory_usage,
				"memory_total": memory_total,
			},
		})
	case "disk":
		disk_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_disk_util"], cluster), c)
		disk_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_disk_usage"], cluster), c)
		disk_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_disk_capacity"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"disk": echo.Map{
				"disk_util":  disk_util,
				"disk_usage": disk_usage,
				"disk_total": disk_total,
			},
		})
	case "pod":
		pod_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_pod_util"], cluster), c)
		pod_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_pod_running"], cluster), c)
		pod_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_pod_quota"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"pod": echo.Map{
				"pod_util":  pod_util,
				"pod_usage": pod_usage,
				"pod_total": pod_total,
			},
		})
	case "api":
		apiserver_request := QueryRange(addr, clusterExpr(Mtemplates["apiserver_request_rate"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"api": echo.Map{
				"apiserver_request": apiserver_request,
				"apiserver_latency": "Not Found",
			},
		})
	case "schedule":
		tmpStr := "round(rate(scheduler_pod_scheduling_attempts_count{$1}[" + c.QueryParam("step") + "])*" + strconv.FormatInt(GetDuration(c), 10) + ")"
		log.Println("tmpStr : ", tmpStr)
		scheduler_schedule_attempts := QueryRange(addr, clusterExpr(tmpStr, cluster), c)
		schedule_attempts_total := QueryRange(addr, clusterExpr(Mtemplates["scheduler_schedule_attempts_total"], cluster), c)
		schedule_fail := QueryRange(addr, clusterExpr(Mtemplates["scheduler_schedule_fail"], cluster), c)
		schedule_fail_total := QueryRange(addr, clusterExpr(Mtemplates["scheduler_schedule_fail_total"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"scheduler": echo.Map{
				"scheduler_schedule_attempts": scheduler_schedule_attempts,
				"schedule_attempts_total":     schedule_attempts_total,
				"schedule_fail":               schedule_fail,
				"schedule_fail_total":         schedule_fail_total,
			},
		})
	case "":
		// 모든 클러스터 출력

		cpu_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_cpu_util"], cluster), c)
		cpu_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_cpu_usage"], cluster), c)
		cpu_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_cpu_core_total"], cluster), c)
		memory_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_memory_util"], cluster), c)
		memory_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_memory_usage"], cluster), c)
		memory_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_memory_total"], cluster), c)
		disk_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_disk_util"], cluster), c)
		disk_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_disk_usage"], cluster), c)
		disk_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_disk_capacity"], cluster), c)
		pod_util := QueryRange(addr, clusterExpr(Mtemplates["cluster_pod_util"], cluster), c)
		pod_usage := QueryRange(addr, clusterExpr(Mtemplates["cluster_pod_running"], cluster), c)
		pod_total := QueryRange(addr, clusterExpr(Mtemplates["cluster_pod_quota"], cluster), c)
		apiserver_request := QueryRange(addr, clusterExpr(Mtemplates["apiserver_request_rate"], cluster), c)
		tmpStr := "round(rate(scheduler_pod_scheduling_attempts_count{$1}[" + c.QueryParam("step") + "])*" + strconv.FormatInt(GetDuration(c), 10) + ")"
		log.Println("tmpStr : ", tmpStr)
		scheduler_schedule_attempts := QueryRange(addr, clusterExpr(tmpStr, cluster), c)
		schedule_attempts_total := QueryRange(addr, clusterExpr(Mtemplates["scheduler_schedule_attempts_total"], cluster), c)
		schedule_fail := QueryRange(addr, clusterExpr(Mtemplates["scheduler_schedule_fail"], cluster), c)
		schedule_fail_total := QueryRange(addr, clusterExpr(Mtemplates["scheduler_schedule_fail_total"], cluster), c)

		return c.JSON(http.StatusOK, echo.Map{
			"cpu": echo.Map{
				"cpu_util":  cpu_util,
				"cpu_usage": cpu_usage,
				"cpu_total": cpu_total,
			},
			"memory": echo.Map{
				"memory_util":  memory_util,
				"memory_usage": memory_usage,
				"memory_total": memory_total,
			},
			"disk": echo.Map{
				"disk_util":  disk_util,
				"disk_usage": disk_usage,
				"disk_total": disk_total,
			},
			"pod": echo.Map{
				"pod_util":  pod_util,
				"pod_usage": pod_usage,
				"pod_total": pod_total,
			},
			"apiserver": echo.Map{
				"apiserver_request": apiserver_request,
				"apiserver_latency": "Not Found",
			},
			"scheduler": echo.Map{
				"scheduler_schedule_attempts": scheduler_schedule_attempts,
				"schedule_attempts_total":     schedule_attempts_total,
				"schedule_fail":               schedule_fail,
				"schedule_fail_total":         schedule_fail_total,
			},
		})

	default:

		return c.JSON(http.StatusNotFound, echo.Map{
			"errors": echo.Map{
				"status_code": http.StatusNotFound,
				"message":     "Not Found",
				"command":     "cpu, memory, disk..",
			},
		})
	}
	return nil
}
