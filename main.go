package main

import (
	"flag"

	"github.com/astaxie/beego"
	_ "github.com/uxff/taniago/conf/inits"
	_ "github.com/uxff/taniago/routers"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/controllers"
	"github.com/uxff/taniago/models"
)

func main() {
	logdeep := 3
	serveDir := "r:/themedia" //"."

	flag.IntVar(&logdeep, "logdeep", logdeep, "log deep")
	flag.StringVar(&serveDir, "dir", serveDir, "serve dir, witch will browse")
	flag.Parse()

	logs.SetLevel(logs.LevelInfo)
	logs.SetLogFuncCallDepth(logdeep)

	beego.SetStaticPath("fs", serveDir)

	controllers.SetLocalDirRoot(serveDir)
	models.LoadFriendLinksFromFile("./conf/friends.json")

	logs.Info("the serve dir=%s", serveDir)

	beego.Run()
}
