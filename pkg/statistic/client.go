package statistic

import (
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Client struct {
	v1.API
	Range v1.Range
}

func New(addr string, start, end time.Time, step time.Duration) *Client {
	c, err := api.NewClient(api.Config{
		Address:      addr,
		RoundTripper: api.DefaultRoundTripper,
	})
	if err != nil {
		panic(err)
	}
	return &Client{
		API: v1.NewAPI(c),
		Range: v1.Range{
			Start: start,
			End:   end,
			Step:  step,
		},
	}
}
