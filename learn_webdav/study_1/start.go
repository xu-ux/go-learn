package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"
)

func main() {
	var addr *string
	var path *string
	addr = flag.String("addr", ":8080", "") // listen端口，默认8080
	path = flag.String("path", ".", "")     // 文件路径，默认当前目录
	flag.Parse()
	fmt.Println("addr=", *addr, ", path=", *path) // 在控制台输出配置
	http.ListenAndServe(*addr, &webdav.Handler{
		FileSystem: webdav.Dir(*path),
		LockSystem: webdav.NewMemLS(),
	})
}
