/*
	图片化浏览某个目录
	如果浏览目录dir1,则需要目录${dir1}支持或提供以下特性：
		- ${dir1}/thumb.jpg		# 用于封面图 可以是.jpg,.png,.gif
		- ${dir1}/thumbs/		# 用于存放原图对应的缩略图
		- 目录名称当前图集显示名称
	// done: cache, paysapi, static fs, page
	// todo: multi domain, login/register, pay and access, advertise, shopping mall, bitpay, ethereum pay，
, fake order, friendly link(hard)
*/
package controllers

import (
	"time"
	"path"
	"github.com/astaxie/beego"
	"github.com/uxff/taniago/utils/paginator"
	"github.com/uxff/taniago/models/picset"
	"github.com/astaxie/beego/logs"
)

type Picset struct {
	Dirpath string
	Name string
	Thumb string
	Url string
}

type CacheItem struct {
	Created time.Time
	PicsetList []*Picset
}

var fsRoute = "/fs"
var picsetRoute = "/picset"
var pageSize = 20


type PicsetController struct {
	BaseController
}


// picset list
func (this *PicsetController) Picset() {

	this.Data["appname"] = beego.AppConfig.String("appname")

	//
	fullDirName := this.Ctx.Input.Param(":splat")
	//c := this.GetString("c")
	//switch c {
	//case "clearcache":
	//	//dirCache = make(map[string][]*Picset, 0)
	//	logs.Debug("dir cache cleared")
	//}

	//logs.Info("fullDirName from param=%v", fullDirName)

	curDirName := path.Base(fullDirName)
	fullParentName := path.Dir(fullDirName)

	this.Data["curDirName"] = curDirName
	this.Data["fullParentName"] = fullParentName
	this.Data["parentLink"] = picsetRoute+"/"+fullParentName

	//logs.Info("fullDirName from url param:%s curDirName=%s parentName=%s", fullDirName, curDirName, fullParentName )
	//dirpath := localDirRoot+"/"+fullDirName

	//thedirnames := make([]string, 0)
	theDirList := picset.GetPicsetListFromDir(fullDirName, picsetRoute, fsRoute)

	if theDirList == nil {
		//logs.Warn("open dir %s error:%v", fullDirName)
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Ctx.ResponseWriter.Write([]byte("dir not exist:"+fullDirName))
		return
	}

	allNum := len(theDirList)

	p := paginator.NewPaginator(this.Ctx.Request, pageSize, int64(allNum))
	this.Data["paginator"] = p

	last:= p.Page()*pageSize
	if last >= len(theDirList) {
		last = len(theDirList)
	}
	thePagedDirList := theDirList[(p.Page()-1)*pageSize:last]

	this.Data["thedirnames"] = thePagedDirList //theDirList

	this.TplName = "picset/view.tpl"
}

func (this *PicsetController) ClearCache() {
	picset.ClearCache()
	logs.Debug("picset cache cleared")
	this.EnableRender = false
}
