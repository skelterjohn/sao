package cmp

import (
	"github.com/hoisie/web.go"
)

func RegisterAll() {
    web.Get("/Button", Button)
	web.Get("/TextEdit", TextEdit)
}