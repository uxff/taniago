package main

import (
	"flag"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/conf/inits"
	"github.com/uxff/taniago/models"
	"github.com/uxff/taniago/models/picset"
	_ "github.com/uxff/taniago/routers"
	"github.com/uxff/taniago/utils"
)

func main() {
	logdeep := 3
	serveDir := "e://" //"."
	addr := ":" + beego.AppConfig.String("httpport")
	serveStatic := false

	//flag.IntVar(&logdeep, "logdeep", logdeep, "log deep")
	flag.StringVar(&serveDir, "dir", serveDir, "serve dir, witch will browse")
	flag.StringVar(&addr, "addr", addr, "beego run param addr, format as ip:port")
	flag.BoolVar(&serveStatic, "s", serveStatic, "serve static")
	//flag.StringVar(&appenv, "env", appenv, "app env, in app.conf")//use env BEEGO_MODE=dev
	flag.Parse()

	logs.SetLevel(logs.LevelDebug)
	logs.SetLogFuncCallDepth(logdeep)

	// todo: use nginx instead
	if serveStatic {
		beego.SetStaticPath("fs", serveDir)
	}
	//beego.AppConfig.Set("", "")

	curDir := os.Getenv("PWD")
	localViewPath := curDir + "/views/"
	// if local views path not exist, use gopath/uxff/taniago/views
	if utils.IsDirExist(localViewPath) {
		beego.SetViewsPath(localViewPath)
	} else {
		gopath := os.Getenv("GOPATH")
		if gopath != "" {
			beego.SetViewsPath(gopath + "/src/uxff/taniago/views/")
		}
	}

	//controllers.SetLocalDirRoot(serveDir)
	picset.SetLocalDirRoot(serveDir)
	//models.LoadIndexLinksFromFile("./conf/friends.json")

	models.SetLinksPath("./conf/index.json")
	models.LoadIndexLinks()
	models.SetFriendlyLinksPath("./conf/friends.json")
	models.LoadFriendlyLinks()

	logs.Info("beego server will run. dir=%s addr=%s serveStatic=%v", serveDir, addr, serveStatic)

	inits.PrepareDb()

	beego.Run(addr)
}
