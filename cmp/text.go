package cmp

import (
	"fmt"
	"github.com/hoisie/web.go"
)

func TextEdit(ctx *web.Context, query string) (code string) {
	text := fmt.Sprintf("%v", ctx.Request.Params)
	return fmt.Sprintf(`<textarea style="width:375px; height:375px;">%s</textarea>`, text)
}