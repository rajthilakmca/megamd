package api

import (
	"net/http"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"github.com/googollee/go-socket.io"
)

type MegdHandler struct {
	method string
	path   string
	h      http.Handler
}

var megdHandlerList []MegdHandler

//RegisterHandler inserts a handler on a list of handlers
func RegisterHandler(path string, method string, h http.Handler) {
	var th MegdHandler
	th.path = path
	th.method = method
	th.h = h
	megdHandlerList = append(megdHandlerList, th)
}

// RunServer starts vertice httpd server.
func NewNegHandler() *negroni.Negroni {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})

	m := &delayedRouter{}
	for _, handler := range megdHandlerList {
		m.Add(handler.method, handler.path, handler.h)
	}

	server, err := socketio.NewServer(nil)
    if err != nil {
        fmt.Println(err)
    }

	m.Add("Get", "/", Handler(index))
	m.Add("Post", "/logs/", server)
	m.Add("Get", "/logs/", server)
	//m.Add("Get", "/logs", Handler(logs))
	m.Add("Get", "/ping", Handler(ping))
	m.Add("Get", "/vnc/", server)
	m.Add("Post", "/vnc/", server)
	//m.Add("Get", "/vnc/", Handler(vnc))

  socketHandler(server)
	//socketProcess(server)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(c)
	n.Use(newLoggerMiddleware())
	n.UseHandler(m)
	n.Use(negroni.HandlerFunc(contextClearerMiddleware))
	n.Use(negroni.HandlerFunc(flushingWriterMiddleware))
	n.Use(negroni.HandlerFunc(errorHandlingMiddleware))
	n.Use(negroni.HandlerFunc(authTokenMiddleware))
	n.UseHandler(http.HandlerFunc(runDelayedHandler))
	return n
}
