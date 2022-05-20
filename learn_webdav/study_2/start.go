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
	var user = *flag.String("user", "admin", "")
	var pass = *flag.String("password", "password", "")
	addr = flag.String("addr", ":8080", "")
	path = flag.String("path", ".", "")
	flag.Parse()

	fs := &webdav.Handler{
		FileSystem: webdav.Dir(*path),
		LockSystem: webdav.NewMemLS(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// 获取用户名/密码
		username, password, ok := req.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// 验证用户名/密码
		if username != user || password != pass {
			http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
			return
		}
		fs.ServeHTTP(w, req)
	})
	fmt.Println("addr=", *addr, ", path=", *path)
	http.ListenAndServe(*addr, nil)
}
