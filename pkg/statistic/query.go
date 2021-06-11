package statistic

import (
	"context"
	"time"

	log "github.com/gogap/logrus"
	"github.com/prometheus/common/model"
)

const NODE_STATUS = `up{container="node-exporter"}`

type Query struct {
	Time time.Time
}

func QueryPrometheus(client *Client) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	q := Query{}
	q.Time = time.Now()

	if v, err := q.query(ctx, client); err != nil {
		log.Errorf("prometheus query failed: %s", err)
		return
	} else {
		log.Info("查询数据为: ", v)
	}
}

func (q *Query) query(ctx context.Context, client *Client) (model.SampleValue, error) {
	value, _, err := client.Query(ctx, NODE_STATUS, q.Time)
	if err != nil {
		return -1, err
	}
	if value.(model.Vector).Len() == 0 {
		log.Warnf("prometheus query got nil value: %s", NODE_STATUS)
		return 0, nil
	}
	return value.(model.Vector)[0].Value, nil
}
