package main

import (
	"github.com/hoisie/web.go"
	"github.com/skelterjohn/soa"
	"github.com/skelterjohn/soa/cmp"
)

func Container(val string) (code string) {
	code = cmp.Embed("http://localhost:9999/Button", 100, 100)
	return
}

func main() {
	soa.LaunchBrowser("http://localhost:9999/Container")
    web.Get("/(Container)", Container)
    web.Get("/Button?(.*)", cmp.Button)
    web.Run("0.0.0.0:9999")
}