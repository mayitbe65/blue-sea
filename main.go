package main

import (
	_ "blue-sea/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run("127.0.0.1:8080")
}
