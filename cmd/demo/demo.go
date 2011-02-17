package main

import (
	"fmt"
	"github.com/hoisie/web.go"
	"github.com/skelterjohn/sao"
	"github.com/skelterjohn/sao/cmp"
)

func RunDemo(tec, bec cmp.EmbeddedComponent) {
	msg := <-bec.FromCmp
	fmt.Printf("%v\n", msg)
}

func Demo() (code string) {
	tec := cmp.Embed("localhost:9999", "TextEdit", map[string]string{}, 400, 400)
	bec := cmp.Embed("localhost:9999", "Button", map[string]string{"label":"Save"}, 100, 100)
	code = tec.Code + bec.Code
	//go RunDemo(tec, bec)
	return
}

func main() {
	sao.LaunchBrowser("http://localhost:9999/Demo")
    cmp.RegisterAll()
    web.Get("/Demo", Demo)
	web.Run("0.0.0.0:9999")
}
