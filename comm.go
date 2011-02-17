package sao

import (
	"fmt"
	"http"
	"bufio"
	"os"
	"bytes"
	"github.com/hoisie/web.go"
)

func RecvChanClient(host, id string) (recv <-chan string) {
	recvc := make(chan string)
	recv = recvc
	url := fmt.Sprintf("http://%s/%s/Recv", host, id)
	go func() {
		var err os.Error
		var response *http.Response
		var message string
		for err == nil {
			response, _, err = http.Get(url)
			reader := bufio.NewReader(response.Body)
			message, err = reader.ReadString('\n')
			if err != nil {
				recvc <- message
			}
		}
	}()
	return
}

func SendChanClient(host, id string) (send chan<- string) {
	sendc := make(chan string)
	send = sendc
	url := fmt.Sprintf("http://%s/%s/Send", host, id)
	go func() {
		for message := range sendc {
			http.Post(url, "text/plain", bytes.NewBufferString("msg="+message))
		}
	}()
	return
}

func SendChanServer(id string) (send chan<- string) {
	sendc := make(chan string)
	send = sendc
	
	listen := func() (code string) {
		code = <- sendc
		return
	}
	where := fmt.Sprintf("/%s/Recv", id)
	SafeGet(where, listen)
	
	return
}

func RecvChanServer(id string) (recv <-chan string) {
	recvc := make(chan string)
	recv = recvc
	
	listen := func(ctx *web.Context) {
		recvc <- ctx.Request.Params["msg"]
	} 
	where := fmt.Sprintf("/%s/Send", id)
	SafePost(where, listen)
	
	return
}