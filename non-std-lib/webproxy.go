package main

import (
	"path/filepath"
	gwv "simonwaldherr.de/go/gwv"
)

func main() {
	HTTPD := gwv.NewWebServer(8080, 60)

	HTTPD.URLhandler(
		gwv.Favicon(filepath.Join(".", "static", "favicon.ico")),
		gwv.Redirect("^/go/$", "/golang/", 301),
		gwv.Proxy("^/selfcss/", "http://selfcss.org/"),
		gwv.Proxy("^/golang/", "https://golang.org/"),
	)

	HTTPD.Start()
	HTTPD.WG.Wait()
}
