package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type WebHandler1 struct {
}

func (w WebHandler1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	auth := request.Header.Get("Authorization")
	if auth == "" {
		writer.Header().Set("WWW-Authenticate", `Basic reals = "Input your name and password"`)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	authList := strings.Split(auth, " ")
	if len(authList) == 2 && authList[0] == "Basic" {
		authInfo, err := base64.StdEncoding.DecodeString(authList[1])
		if err == nil && string(authInfo) == "Leiyi:123" {
			writer.Write([]byte(fmt.Sprintf("<h1>Web1, from %s</h1>", request.RemoteAddr)))
			return
		}
	}
	writer.Write([]byte("<h1>Wrong name or password</h1>"))
}

type WebHandler2 struct {
}

func (w WebHandler2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Web2</h1>"))
}

func main() {
	c := make(chan os.Signal)

	go func() {
		http.ListenAndServe(":9091", WebHandler1{})
	}()

	go func() {
		http.ListenAndServe(":9092", WebHandler2{})
	}()

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println(s)
}
