package main

import (
	"Airttp/modules"
	"net/rpc"
	"net"
	"log"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Http int

func main() {
	http := new(Http)

	server := rpc.NewServer()
	server.RegisterName("Http", http)

	l, e := net.Listen("tcp", ":5003")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("server start")
	server.Accept(l)
}

func (t *Http) Module(params modules.ModuleParams, result *modules.ModuleParams) error {
	fmt.Print("New Request : ")
	result.Copy(params)
	if (result.Req.Method != "GET") {
		return nil
	}
	dat, err := ioutil.ReadFile(result.Req.Uri)
	if (err != nil) {
		result.Res.Body = []byte(err.Error())
		result.Res.Headers["Content-Length"] = strconv.Itoa(len(result.Res.Body))
		result.Res.Code = 404
		result.Res.Message = "Not Found"
	} else {
		result.Res.Body = dat
		result.Res.Headers["Content-Length"] = strconv.Itoa(len(dat))
		result.Res.Code = 200
		result.Res.Message = "OK"
	}
	fmt.Println("OK")
	return nil
}