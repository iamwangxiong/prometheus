package main

import (
	"time"

	log "github.com/gogap/logrus"
	"github.com/ucloud-lee/prometheus/pkg/statistic"
)

const prometheus_addr = "http://113.31.106.132:9090"

func main() {
	t := time.Now()

	client := statistic.New(prometheus_addr, t, t, time.Minute*1)
	if client != nil {
		log.Info("连接prometheus成功...")
	} else {
		log.Error("连接prometheus失败...")
	}

	statistic.QueryPrometheus(client)
}
