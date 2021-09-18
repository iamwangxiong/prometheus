package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

var mfs = make([]*dto.MetricFamily, 0)

// produce metrics
func init() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/get", gogogo)
	go http.ListenAndServe(":8080", nil)
}

func gogogo(w http.ResponseWriter, r *http.Request) {
	var buff bytes.Buffer
	var err error
	textEncoder := expfmt.NewEncoder(&buff, expfmt.FmtText)
	for _, mf := range mfs {
		err = textEncoder.Encode(mf)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}
	fmt.Fprintf(w, buff.String())
}

func main() {
	praseMF()
	for {
	}
}

func praseMF() []*dto.MetricFamily {
	body, err := request()
	if err != nil {
		fmt.Println("err: ", err)
	}
	decoder := expfmt.NewDecoder(strings.NewReader(string(body)), "")

	for {
		mf := new(dto.MetricFamily)
		if err := decoder.Decode(mf); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Name: ", *mf.Name)
		fmt.Println("Metric: ", mf.Metric)
		for k1, v := range mf.Metric {
			if *mf.Name == "go_info" {
				for k2, lb := range v.Label {
					fmt.Println("Label: ", *lb.Name)
					*mf.Metric[k1].Label[k2].Name = "whoami"
				}
			}
			mfs = append(mfs, mf)
		}
	}

	return mfs
}

func request() ([]byte, error) {
	var body []byte
	resp, err := http.Get("http://127.0.0.1:8080/metrics")
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}
