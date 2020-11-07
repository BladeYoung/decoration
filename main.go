package main

import (
	_ "decoration/boot"
	_ "decoration/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
