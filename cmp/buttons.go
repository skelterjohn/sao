package cmp

import (
	"fmt"
	"github.com/skelterjohn/sao"
	"github.com/hoisie/web.go"
)

func RunButton(id string) {
	send := sao.SendChanServer(id)
	recv := sao.RecvChanServer(id)
	val := <-recv
	send <- val
}

func Button(ctx *web.Context, val string) string {
	label := ctx.Request.Params["label"]
	//go RunButton(ctx.Request.Params["sao_id"])
	return fmt.Sprintf(`<input type="button" value="%s" />`, label)
}