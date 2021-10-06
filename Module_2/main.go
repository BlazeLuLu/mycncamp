package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
}

//当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
	//fmt.Fprintf(w, "200\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":80", nil)
}
