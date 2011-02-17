//target:github.com/skelterjohn/soa
package soa

import (
	"os"
)

func LaunchBrowser(url string) (err os.Error) {
	_, err = os.StartProcess("/usr/bin/open", []string{"open", url}, nil, ".", nil) 
	return
}