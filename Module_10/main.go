package main

import (
	"github.com/golang/glog"
	"github.com/mycncamp/Module_10/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func handler(w http.ResponseWriter, r *http.Request) {

	glog.V(4).Info("entering root handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

	//接收客户端 request，并将 request 中带的 header 写入 response header
	for name, headers := range r.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
		}
	}

	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	path := os.Getenv("PATH")
	w.Header().Set("version", path)

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	var statusCode = 200
	w.WriteHeader(statusCode)
	io.WriteString(os.Stdout, "ClientIP: "+r.RemoteAddr+"\n"+"StatusCode: "+strconv.Itoa(statusCode)+"\n")

	glog.V(4).Infof("Respond in %d ms", delay)
}

//当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
	//fmt.Fprintf(w, "200\n")
}

func main() {
	metrics.Register()
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":80", nil)
	http.Handle("/metrics", promhttp.Handler())
}
