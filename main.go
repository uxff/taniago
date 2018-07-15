package main

import (
	"github.com/astaxie/beego"
	_ "github.com/uxff/taniago/conf/inits"
	_ "github.com/uxff/taniago/routers"
	"flag"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/controllers"
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

	logs.Info("the serve dir=%s", serveDir)

	beego.Run()
}
