package server

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

func Run(subCommand SubCommand) {
	webService := new(restful.WebService)

	switch subCommand {
	case FILE:
		webService.Route(webService.GET("/object").To(hello))
	case MONGODB:
		webService.Route(webService.GET("/object").To(mongo))
	}

	restful.Add(webService)
	log.Print("Start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	_, _ = io.WriteString(resp, "file return")
}

func mongo(req *restful.Request, resp *restful.Response) {
	_, _ = io.WriteString(resp, "mongo return")
}
