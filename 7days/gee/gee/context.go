package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 别名
type H map[string]interface{}

// 结构体
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

// 初始化
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 表单
func (c *Context) PostFrom(key string) string {
	return c.Req.FormValue(key)
}

// 查询
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置头部
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// 字符类型
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-type", "applicatino/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML
func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.Writer.Write([]byte(html))
}
