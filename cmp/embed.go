package cmp

import "fmt"

const EmbedFmt = `<object data=%s width="%d" height="%d"> <embed src=%s width="%d" height="%d"> </embed> Error: Embedded data could not be displayed. </object>`

func Embed(url string, width, height int) (code string) {
	return fmt.Sprintf(EmbedFmt, url, width, height, url, width, height)
}