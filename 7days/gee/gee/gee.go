package gee

import (
	"fmt"
	"net/http"
)

// ...
type HandleFunc func(http.ResponseWriter, *http.Request)

// ...
type Engine struct {
	router map[string]HandleFunc
}

// 初始化
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// 添加路由
func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// get请求
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

// post请求
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

// run
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		fmt.Println("Listen And Serve at port 9999")
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
