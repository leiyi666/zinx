package main

import (
	. "go_code/proxy/util"
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyHandler struct {
}

func (ph ProxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			writer.WriteHeader(500)
			log.Println(err)
		}
	}()
	if request.URL.Path == "/a" {
		newReq, _ := http.NewRequest(request.Method, "http://127.0.0.1:9091", request.Body)
		CloneHeader(request.Header, &newReq.Header)
		newResq, _ := http.DefaultClient.Do(newReq)
		getHeader := writer.Header()
		CloneHeader(newResq.Header, &getHeader)
		writer.WriteHeader(newResq.StatusCode)
		defer newResq.Body.Close()
		resqContent, _ := ioutil.ReadAll(newResq.Body)
		writer.Write(resqContent)
		return
	}
	writer.Write([]byte("<h1>default index</h1>"))
}

func main() {
	http.ListenAndServe(":8080", ProxyHandler{})
}
