package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/conf/inits"
	"github.com/uxff/taniago/models"
	"github.com/uxff/taniago/models/picset"
	_ "github.com/uxff/taniago/routers"
)

func main() {
	logdeep := 3
	serveDir := "r:/themedia" //"."
	addr := ":" + beego.AppConfig.String("httpport")

	//flag.IntVar(&logdeep, "logdeep", logdeep, "log deep")
	flag.StringVar(&serveDir, "dir", serveDir, "serve dir, witch will browse")
	flag.StringVar(&addr, "addr", addr, "beego run param addr, format as ip:port")
	//flag.StringVar(&appenv, "env", appenv, "app env, in app.conf")//use env BEEGO_MODE=dev
	flag.Parse()

	logs.SetLevel(logs.LevelDebug)
	logs.SetLogFuncCallDepth(logdeep)

	// todo: use nginx instead
	//beego.SetStaticPath("fs", serveDir)
	//beego.AppConfig.Set("", "")

	//controllers.SetLocalDirRoot(serveDir)
	picset.SetLocalDirRoot(serveDir)
	//models.LoadIndexLinksFromFile("./conf/friends.json")

	models.SetLinksPath("./conf/index.json")
	models.LoadIndexLinks()
	models.SetFriendlyLinksPath("./conf/friends.json")
	models.LoadFriendlyLinks()

	logs.Info("beego server will run. dir=%s addr=%s", serveDir, addr)

	inits.PrepareDb()

	beego.Run(addr)
}
