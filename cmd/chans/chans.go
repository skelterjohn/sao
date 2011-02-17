package main

import (
	"fmt"
	"github.com/hoisie/web.go"
	"github.com/skelterjohn/sao"
)

func main() {
    go func() {
		sch := sao.SendChanServer("1")
		sch <- "hi"
		
		rch := sao.RecvChanServer("1")
		message := <-rch
		fmt.Println(message)
	}()
	go func() {
		rch := sao.RecvChanClient("localhost:9999", "1")
		message := <- rch
		fmt.Println(message)
		
		sch := sao.SendChanClient("localhost:9999", "1")
		sch <- "omg!"
	}()
	web.Run("0.0.0.0:9999")
}
