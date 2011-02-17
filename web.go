package sao

import (
	"github.com/hoisie/web.go"
)

var block = make(chan bool, 1)

func SafeGet(where string, foo interface{}) {
	block <- true
	defer func() {<-block}()
	web.Get(where, foo)
}
func SafePost(where string, foo interface{}) {
	block <- true
	defer func() {<-block}()
	web.Post(where, foo)
}