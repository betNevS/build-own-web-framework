package main

import (
	"fmt"

	"github.com/betNevS/build-own-web-framework/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("foo", BarControllerHandler)
	fmt.Println("register end")
}
