package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

var nowNamespaceMetric = map[string]string{
	"namespace_cpu":       "round(sum(sum(irate(container_cpu_usage_seconds_total{job='kubelet',pod!='',image!='', $1}[5m]))by(namespace,pod,cluster))by(namespace,cluster),0.001)",
	"namespace_memory":    "sum(sum(container_memory_rss{job='kubelet',pod!='',image!='', $1})by(namespace,pod,cluster))by(namespace,cluster)",
	"namespace_pod_count": "count(count(container_spec_memory_reservation_limit_bytes{pod!='', $1})by(pod,cluster,namespace))by(cluster,namespace)",
}

func NowMonit(k string, c string, n string, m []string) interface{} {

	fmt.Println(c, n)

	//필요 파라미터 검증
	if check := strings.Compare(c, "")*strings.Compare(n, "")*len(m) == 0; check {
		return nil //에러 반환
	}

	//메트릭 검증
	for _, metric := range m {
		if metric == "" {
			continue
		}
		if check := strings.Compare(nowNamespaceMetric[metric], "") == 0; check {
			return nil
		}
	}

	//Prometheus call
	addr := "http://192.168.150.115:31298/"
	// result := map[string]model.Value{}

	result := map[string]model.Value{}
	for i, metric := range m {
		if metric == "" {
			continue
		}
		var data model.Value
		switch k {
		case "namespace":
			temp_filter := map[string]string{
				"cluster":   c,
				"namespace": n,
			}
			data = nowQueryRange(addr, nowMetricExpr(nowNamespaceMetric[m[i]], temp_filter))
		default:
			return nil
		}

		result[m[i]] = data
	}

	return result
}

func nowQueryRange(endpointAddr string, query string) model.Value {
	var start_time time.Time
	var end_time time.Time
	var step time.Duration

	t := time.Now()
	// tm, _ := strconv.ParseInt(c.QueryParam("start"), 10, 64)
	start_time = time.Unix(t.Unix(), 0)
	// log.Println(start_time)

	// tm2, _ := strconv.ParseInt(c.QueryParam("end"), 10, 64)
	end_time = time.Unix(t.Unix(), 0)
	// log.Println(end_time)

	tm3, _ := time.ParseDuration("1s")
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
	}

	if len(warnings) > 0 {
		log.Printf("Warnings: %v\n", warnings)
	}
	// log.Printf("Result:\n%v\n", result)
	return result
}

func nowMetricExpr(val string, filter map[string]string) string {
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
