package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"
	"os"
)

var (
	flagRootDir   = flag.String("dir", ".", "webdav root dir")
	flagHttpAddr  = flag.String("http", ":8080", "http or https address")
	flagHttpsMode = flag.Bool("https-mode", false, "use https mode")
	flagCertFile  = flag.String("https-cert-file", "cert.pem", "https cert file")
	flagKeyFile   = flag.String("https-key-file", "key.pem", "https key file")
	flagUserName  = flag.String("user", "admin", "user name")
	flagPassword  = flag.String("password", "password", "user password")
	flagReadonly  = flag.Bool("read-only", false, "read only")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of WebDAV Server\n")
		flag.PrintDefaults()
	}
}
func main() {
	flag.Parse()
	fs := &webdav.Handler{
		FileSystem: webdav.Dir(*flagRootDir),
		LockSystem: webdav.NewMemLS(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if *flagUserName != "" && *flagPassword != "" {
			username, password, ok := req.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if username != *flagUserName || password != *flagPassword {
				http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
				return
			}
		}
		if *flagReadonly {
			switch req.Method {
			case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
				http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
				return
			}
		}
		fs.ServeHTTP(w, req)
	})
	if *flagHttpsMode {
		fmt.Println("HTTPS", "addr=", *flagHttpAddr, ", path=", *flagRootDir)
		http.ListenAndServeTLS(*flagHttpAddr, *flagCertFile, *flagKeyFile, nil)
	} else {
		fmt.Println("HTTP", "addr=", *flagHttpAddr, ", path=", *flagRootDir)
		http.ListenAndServe(*flagHttpAddr, nil)
	}
}
