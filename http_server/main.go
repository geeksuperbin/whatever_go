package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// 一个 http server 监听 63199 端口
// 请求地址返回 json {"code": 1}
// 可用于回调服务器

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		glog.SetPath("/tmp/logs")
		glog.File("access-{Ymd}.log").Println(r.GetBodyString())

		r.Response.WriteExit("{\"code\":1}")
	})
    s.SetPort(63199)
	s.Run()
}