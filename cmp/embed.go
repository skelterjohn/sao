package cmp

import "fmt"

const EmbedFmt = 
`
<object data=%s width="%d" height="%d">
 <embed src=%s width="%d" height="%d">
 </embed>
Error: Embedded data could not be displayed.
</object>
`

var idBlock = make(chan bool, 1)
var NextID int = 1

func GetNextID() (id string) {
	idBlock<-true
	id = fmt.Sprintf("%d", NextID)
	NextID++
	<-idBlock
	return
}

type EmbeddedComponent struct {
	ID string
	Code string
	ToCmp chan<- interface{}
	FromCmp <-chan interface{}
}

func Embed(host, kind string, params map[string]string, width, height int) (ec EmbeddedComponent) {
	qstring := ""
	for k, v := range params {
		qstring += fmt.Sprintf("%s=%s&", k, v)
	}
	id := GetNextID()
	qstring += fmt.Sprintf("sao_width=%d&sao_height=%d&sao_id=%s", width, height, id)
	url := fmt.Sprintf("http://%s/%s?%s", host, kind, qstring)
	ec.Code = fmt.Sprintf(EmbedFmt, url, width, height, url, width, height)
	ec.ID = id
	
	
	return
}
