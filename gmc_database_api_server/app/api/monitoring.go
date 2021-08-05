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

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

var clusterMetric = map[string]string{
	"cpu_util":     "round(100-(avg(irate(node_cpu_seconds_total{mode='idle', $1}[5m]))by(cluster)*100),0.1)",
	"cpu_usage":    "round(sum(rate(container_cpu_usage_seconds_total{id='/', $1}[2m]))by(cluster),0.01)",
	"cpu_total":    "sum(machine_cpu_cores{$1})by(cluster)",
	"memory_util":  "round(sum(node_memory_MemTotal_bytes{$1}-node_memory_MemFree_bytes-node_memory_Buffers_bytes-node_memory_Cached_bytes-node_memory_SReclaimable_bytes)by(cluster)/sum(node_memory_MemTotal_bytes)by(cluster)*100,0.1)",
	"memory_usage": "round(sum(node_memory_MemTotal_bytes{$1}-node_memory_MemFree_bytes-node_memory_Buffers_bytes-node_memory_Cached_bytes-node_memory_SReclaimable_bytes)by(cluster)/1024/1024/1024,0.01)",
	"memory_total": "round(sum(node_memory_MemTotal_bytes{$1})by(cluster)/1024/1024/1024,0.01)",
	"disk_util":    "round(((sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)-sum(node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster))/(sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)))*100,0.1)",
	"disk_usage":   "round((sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)-sum(node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster))/1000/1000/1000,0.01)",
	"disk_total":   "round(sum(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs', $1})by(cluster)/1000/1000/1000,0.01)",
	"pod_running":  "count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(cluster,pod))by(cluster)",
	"pod_quota":    "sum(max(kube_node_status_capacity{resource='pods', $1})by(node,cluster)unless on(node,cluster)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0))by(cluster)",
	"pod_util":     "round((count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(cluster,pod))by(cluster))/(sum(max(kube_node_status_capacity{resource='pods', $1})by(node,cluster)unless on(node,cluster)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0))by(cluster))*100,0.1)",

	"apiserver_request_rate": "round(sum(irate(apiserver_request_total{$1}[5m]))by(cluster),0.001)",
	//
	"scheduler_schedule_attempts_total": "scheduler_pod_scheduling_attempts_count{$1}",
	"scheduler_schedule_fail":           "sum(rate(scheduler_pending_pods{$1}[5m]))by(cluster)",
	"scheduler_schedule_fail_total":     "sum(scheduler_pending_pods{$1})by(cluster)",
}

var namespaceMetric = map[string]string{ //쿼리문 확인 필요
	"namespace_cpu":       "round(sum(sum(irate(container_cpu_usage_seconds_total{job='kubelet',pod!='',image!='', $1}[5m]))by(namespace,pod,cluster))by(namespace,cluster),0.001)",
	"namespace_memory":    "sum(sum(container_memory_usage_bytes{job='kubelet',pod!='',image!='', $1})by(namespace,pod,cluster))by(namespace,cluster)",
	"namespace_pod_count": "count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(pod,cluster,namespace))by(cluster,namespace)",
}

var podMetric = map[string]string{
	"pod_cpu":                   "round(sum(irate(container_cpu_usage_seconds_total{job='kubelet',pod!='',image!='', $1}[5m]))by(namespace,pod,cluster)*1000,0.001)",
	"pod_memory":                "sum(container_memory_usage_bytes{job='kubelet',pod!='',image!='', $1})by(namespace,pod,cluster)",
	"pod_net_bytes_transmitted": "round(sum(irate(container_network_transmit_bytes_total{pod!='',interface!~'^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)',job='kubelet', $1}[5m]))by(namespace,pod,cluster)/125,0.01)",
	"pod_net_bytes_received":    "round(sum(irate(container_network_receive_bytes_total{pod!='',interface!~'^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)',job='kubelet', $1}[5m]))by(namespace,pod,cluster)/125,0.01)",
}

var nodeMetric = map[string]string{ //쿼리 수정 필요
	"node_cpu_util":           "100-(avg(irate(node_cpu_seconds_total{mode='idle', $1}[5m]))by(instance)*100)",
	"node_cpu_usage":          "sum(rate(container_cpu_usage_seconds_total{id='/'}[5m]))by(node)",
	"node_cpu_total":          "sum(machine_cpu_cores)by(node)",
	"node_memory_util":        "(node_memory_MemTotal_bytes-node_memory_MemAvailable_bytes)/node_memory_MemTotal_bytes",
	"node_memory_usage":       "node_memory_MemTotal_bytes-node_memory_MemFree_bytes-node_memory_Buffers_bytes-node_memory_Cached_bytes-node_memory_SReclaimable_bytes",
	"node_memory_total":       "sum(node_memory_MemTotal_bytes)by(instance)",
	"node_disk_size_util":     "100-((node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs'} * 100)/node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs'})",
	"node_disk_size_usage":    "(node_filesystem_size_bytes{mountpoint='/',fstype!='rootfs'}-(node_filesystem_avail_bytes{mountpoint='/',fstype!='rootfs'}))",
	"node_disk_size_capacity": "/api/v1/query_range?query(node_filesystem_size_bytes{mountpoint='/'',fstype!='rootfs'})",
	// node_pod_utilisation/{cluster_name} "sum(kubelet_running_pods)by(node)/(max(kube_node_status_capacity%7Bcluster='{cluster_name}',resource='pods'%7D)by(node)unless%20on(node)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0))*100"
	"node_pod_running_count":     "sum(kubelet_running_pods)by(node)",
	"node_pod_quota":             "max(kube_node_status_capacity{resource='pods'})by(node)unless on(node)(kube_node_status_condition{condition='Ready',status=~'unknown|false'}>0)",
	"node_disk_inode_util":       "100-(node_filesystem_files_free{mountpoint='/'}/node_filesystem_files{mountpoint='/'}*100)",
	"node_disk_inode_total":      "node_filesystem_files{mountpoint='/'}",
	"node_disk_inode_usage":      "node_filesystem_files{mountpoint='/'}-node_filesystem_files_free{mountpoint='/'}",
	"node_disk_read_iops":        "rate(node_disk_reads_completed_total[5m])",
	"node_disk_write_iops":       "rate(node_disk_writes_completed_total[5m])",
	"node_disk_read_throughput":  "irate(node_disk_read_bytes_total[5m])",
	"node_disk_write_throughput": "irate(node_disk_written_bytes_total[5m])",
	"node_net_bytes_transmitted": "irate(node_network_transmit_bytes_total{device='ens3'}[5m])",
	"node_net_bytes_received":    "irate(node_network_receive_bytes_total{device='ens3'}[5m])",
}

var appMetric = map[string]string{}

var gpuMetric = map[string]string{
	"gpu_temperature":  "nvidia_smi_temperature_gpu{$1}",
	"gpu_power":        "nvidia_smi_power_draw_watts{$1}",
	"gpu_power_limit":  "nvidia_smi_power_limit_watts{$1}",
	"gpu_memory_total": "nvidia_smi_memory_total_bytes{$1}",
	"gpu_memory_used":  "nvidia_smi_memory_used_bytes{$1}",
	"gpu_memory_free":  "nvidia_smi_memory_free_bytes{$1}",
	"gpu_ratio":        "nvidia_smi_utilization_gpu_ratio{$1}",
	"gpu_memory_ratio": "nvidia_smi_utilization_memory_ratio{$1}",
	"gpu_fan_speed":    "nvidia_smi_fan_speed_ratio{$1}",
}

func Metrics(c echo.Context) (err error) {

	kind := c.Param("kind")

	//0. parameter 입력값이 올바른지 검증한다.
	if !validateParam(c) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error": "Bad Parameter",
		})
	}

	// 1. metric_filte를 parsing 한다
	metric_filter := c.QueryParam("metric_filter")
	metrics := metricParsing(metric_filter)
	//2. metric_filter가 clusterMetric, podMetric에 다 속해 있는지 검증한다.
	if !validateMetric(kind, metrics, c) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error": "Not found metric",
		})
	}

	//2/ 필터 입력값에 대해 검증한다.
	if !validateFilter(kind, c) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error": "Bad Filter",
		})
	}

	// 3. 처리 함수를 만든다 파라미터는 c,kind,metrics 를 입력.
	mericResult(c, kind, metrics)

	return nil
}

func mericResult(c echo.Context, kind string, a []string) error {
	// fmt.Println("metricResult")
	db := db.DbManager()
	addr := "http://192.168.150.115:31298/"

	cluster := c.QueryParam("cluster_filter")
	//cluster DB 유무 체크

	switch cluster {
	case "all":
	default:
		models := FindClusterDB(db, "Name", cluster)

		if models == nil {
			common.ErrorMsg(c, http.StatusNotFound, common.ErrNotFound)
			return nil
		} else {
			log.Println("models find it !")
		}
	}

	//결과를 담기 위한
	result := map[string]model.Value{}

	for k, metric := range a {
		if metric == "" {
			continue
		}

		var data model.Value

		switch kind {
		case "cluster":
			temp_filter := map[string]string{
				"cluster": cluster,
			}
			data = QueryRange(addr, metricExpr(clusterMetric[a[k]], temp_filter), c)
		case "node":
			temp_filter := map[string]string{
				"cluster": cluster,
			}
			data = QueryRange(addr, metricExpr(nodeMetric[a[k]], temp_filter), c)
		case "pod":
			namespace := c.QueryParam("namespace_filter")
			pod := c.QueryParam("pod_filter")
			temp_filter := map[string]string{
				"cluster":   cluster,
				"namespace": namespace,
				"pod":       pod,
			}

			data = QueryRange(addr, metricExpr(podMetric[a[k]], temp_filter), c)

		case "app":
			temp_filter := map[string]string{
				"cluster": cluster,
			}
			data = QueryRange(addr, metricExpr(appMetric[a[k]], temp_filter), c)
		case "namespace":
			namespace := c.QueryParam("namespace_filter")
			temp_filter := map[string]string{
				"cluster":   cluster,
				"namespace": namespace,
			}

			data = QueryRange(addr, metricExpr(namespaceMetric[a[k]], temp_filter), c)
		case "gpu":
			temp_filter := map[string]string{
				"cluster": cluster,
			}
			data = QueryRange(addr, metricExpr(gpuMetric[a[k]], temp_filter), c)
		default:
			return c.JSON(http.StatusNotFound, echo.Map{
				"errors": echo.Map{
					"status_code": http.StatusNotFound,
					"message":     "Not Found",
				},
			})
		}

		result[a[k]] = data
	}

	return c.JSON(http.StatusOK, echo.Map{
		"items": result,
	})

	// return nil
	//t
}

func metricParsing(m string) []string {
	slice := strings.Split(m, "|")
	var arrayMetric []string
	for _, v := range slice {
		if v != "" {
			arrayMetric = append(arrayMetric, v)
		}
	}

	return arrayMetric
}

func validateMetric(k string, m []string, c echo.Context) bool {

	switch k {
	case "cluster":

		for _, v := range m {
			if clusterMetric[v] == "" {
				return false
			}
		}
	case "node":
		for _, v := range m {
			if nodeMetric[v] == "" {
				return false
			}
		}
	case "pod":
		for _, v := range m {
			if podMetric[v] == "" {
				return false
			}
		}
	case "app":
		for _, v := range m {
			if appMetric[v] == "" {
				return false
			}
		}
	case "namespace":
		for _, v := range m {
			if namespaceMetric[v] == "" {
				return false
			}
		}
	case "gpu":
		for _, v := range m {
			if gpuMetric[v] == "" {
				return false
			}
		}
	default:
	}

	return true
}

func validateFilter(k string, c echo.Context) bool {

	switch k {
	case "cluster":
		cluster := c.QueryParam("cluster_filter")
		if check := strings.Compare(cluster, "") == 0; check {
			return false
		}
	case "node":
		cluster := c.QueryParam("cluster_filter")
		if check := strings.Compare(cluster, "") == 0; check {
			return false
		}
	case "pod":
		cluster := c.QueryParam("cluster_filter")
		pod := c.QueryParam("pod_filter")
		namespace := c.QueryParam("namespace_filter")
		if check := strings.Compare(cluster, "")*strings.Compare(pod, "")*strings.Compare(namespace, "") == 0; check {
			return false
		}
	case "app":
		cluster := c.QueryParam("cluster_filter")
		if check := strings.Compare(cluster, "") == 0; check {
			return false
		}
	case "namespace":
		cluster := c.QueryParam("cluster_filter")
		namespace := c.QueryParam("namespace_filter")
		if check := strings.Compare(cluster, "")*strings.Compare(namespace, "") == 0; check {
			return false
		}
	case "gpu":
		cluster := c.QueryParam("cluster_filter")
		if check := strings.Compare(cluster, "") == 0; check {
			return false
		}
	default:
	}

	return true
}

func validateParam(c echo.Context) bool {

	if c.QueryParam("start") == "" {
		return false
	}
	if c.QueryParam("end") == "" {
		return false
	}
	if c.QueryParam("step") == "" {
		return false
	}
	return true
}

// func checkQueryParams(c echo.Context) error {

// 	if c.QueryParam("start") == "" {
// 		return c.JSON(http.StatusBadRequest, echo.Map{
// 			"Error":  "start is empty",
// 			"how-to": "start={start_time}",
// 		})
// 	}
// 	if c.QueryParam("end") == "" {
// 		return c.JSON(http.StatusBadRequest, echo.Map{
// 			"Error":  "end is empty",
// 			"how-to": "end={end_time}",
// 		})
// 	}
// 	if c.QueryParam("step") == "" {
// 		return c.JSON(http.StatusBadRequest, echo.Map{
// 			"Error":  "step is empty",
// 			"how-to": "step={step_time}",
// 		})
// 	}

// 	return nil
// }

func printError(c echo.Context) error {
	return c.JSON(http.StatusNotFound, echo.Map{
		"errors": echo.Map{
			"status_code": http.StatusNotFound,
			"message":     "Not Found",
			"command":     "cpu, memory, disk..",
		},
	})
}

func QueryRange(endpointAddr string, query string, c echo.Context) model.Value {
	log.Println("queryrange in")
	log.Println(query)
	log.Println(endpointAddr)
	var start_time time.Time
	var end_time time.Time
	var step time.Duration

	tm, _ := strconv.ParseInt(c.QueryParam("start"), 10, 64)
	start_time = time.Unix(tm, 0)
	// log.Println(start_time)

	tm2, _ := strconv.ParseInt(c.QueryParam("end"), 10, 64)
	end_time = time.Unix(tm2, 0)
	// log.Println(end_time)

	tm3, _ := time.ParseDuration(c.QueryParam("step"))
	step = tm3
	// log.Println(step)

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
		//실행 폭파 f
	}

	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	// log.Printf("Result:\n%v\n", result)
	return result
}

func GetDuration(c echo.Context) int64 {
	t, _ := time.ParseDuration(c.QueryParam("step"))
	// log.Printf("#4d - %s", t)
	returnVal := int64(t / time.Second)
	// log.Printf("#5d - %t", returnVal)
	return returnVal
}

func metricExpr(val string, filter map[string]string) string {
	var returnVal string

	for k, v := range filter {

		switch v {
		case "all":
			returnVal += fmt.Sprintf(`%s!="",`, k)
		default:
			returnVal += fmt.Sprintf(`%s="%s",`, k, v)
		}

	}

	return strings.Replace(val, "$1", returnVal, -1)
}
