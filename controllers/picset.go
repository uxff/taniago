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
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/models/picset"
	"github.com/uxff/taniago/utils"
	"github.com/uxff/taniago/utils/paginator"
)

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

	//logs.Info("fullDirName from param=%v", fullDirName)

	curDirName := path.Base(fullDirName)
	fullParentName := path.Dir(fullDirName)

	this.Data["curDirName"] = curDirName
	this.Data["fullParentName"] = fullParentName
	this.Data["parentLink"] = picsetRoute + "/" + fullParentName
	this.Data["nothumbUrl"] = picset.NothumbUrl //无法在templete的循环内读取到，衰

	//logs.Info("fullDirName from url param:%s curDirName=%s parentName=%s", fullDirName, curDirName, fullParentName )
	//dirpath := localDirRoot+"/"+fullDirName

	//thedirnames := make([]string, 0)
	theDirList := picset.GetPicsetListFromDir(fullDirName, picsetRoute, fsRoute)

	if theDirList == nil {
		//logs.Warn("open dir %s error:%v", fullDirName)
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Ctx.ResponseWriter.Write([]byte("dir not exist:" + fullDirName))
		return
	}

	allNum := len(theDirList)

	p := paginator.NewPaginator(this.Ctx.Request, pageSize, int64(allNum))
	this.Data["paginator"] = p

	last := p.Page() * pageSize
	if last >= len(theDirList) {
		last = len(theDirList)
	}
	thePagedDirList := theDirList[(p.Page()-1)*pageSize : last]

	for i := range thePagedDirList {
		if thePagedDirList[i].Thumb == "" {
			thePagedDirList[i].Thumb = picset.NothumbUrl
		}
	}

	this.Data["thedirnames"] = thePagedDirList //theDirList

	this.TplName = "picset/viewlarge.tpl"
}

type PicsetResponse struct {
	Code            int          `json:"code"`
	Msg             string       `json:"msg"`
	CurrentPage     int          `json:"currentPage"`
	TotalPage       int          `json:"totalPage"`
	Total           int          `json:"total"`
	PageSize        int          `json:"pageSize"`
	FullDirpath     string       `json:"fullDirpath"`
	CurDirname      string       `json:"curDirname"`
	ParentOfDirpath string       `json:"parentOfDirpath"`
	NothumbUrl      string       `json:"nothumbUrl"`
	Data            []PicsetItem `json:"data"`
}
type PicsetItem struct {
	FileName string `json:"fileName"`
	FileSize string `json:"fileSize"`
	EditTime string `json:"editTime"`
	Url      string `json:"url"`
	Thumb    string `json:"thumb"`
	IsDir    bool   `json:"isDir"`
}

// picset list
func (this *PicsetController) ListByJson() {

	this.Data["appname"] = beego.AppConfig.String("appname")

	//
	fullDirName := this.Ctx.Request.Form.Get("path") //this.Ctx.Input.Param(":splat")
	strings.ReplaceAll(fullDirName, "//", "/")

	//logs.Info("fullDirName from param=%v", fullDirName)

	curDirName := path.Base(fullDirName)
	fullParentName := path.Dir(fullDirName)

	// this.Data["curDirName"] = curDirName
	// this.Data["fullParentName"] = fullParentName
	// this.Data["parentLink"] = picsetRoute + "/" + fullParentName

	//logs.Info("fullDirName from url param:%s curDirName=%s parentName=%s", fullDirName, curDirName, fullParentName )
	//dirpath := localDirRoot+"/"+fullDirName

	//thedirnames := make([]string, 0)
	theDirList := picset.GetPicsetListFromDir(fullDirName, picsetRoute, fsRoute)

	if theDirList == nil {
		//logs.Warn("open dir %s error:%v", fullDirName)
		this.Data["json"] = PicsetResponse{
			Code:       400,
			Msg:        "dir not exist:" + fullDirName,
			NothumbUrl: picset.NothumbUrl,
		}
		this.ServeJSON()
		return
	}

	allNum := len(theDirList)

	p := paginator.NewPaginator(this.Ctx.Request, pageSize, int64(allNum))
	// this.Data["paginator"] = p

	last := p.Page() * pageSize
	if last >= len(theDirList) {
		last = len(theDirList)
	}
	thePagedDirList := theDirList[(p.Page()-1)*pageSize : last]

	// this.Data["thedirnames"] = thePagedDirList //theDirList

	this.EnableRender = false
	resp := PicsetResponse{
		Code:            0,
		Msg:             "ok",
		Data:            convertToPicsetItemList(thePagedDirList),
		CurrentPage:     p.Page(),
		TotalPage:       p.PageNums(),
		Total:           len(theDirList),
		PageSize:        p.PerPageNums,
		CurDirname:      curDirName,
		FullDirpath:     fullDirName,
		NothumbUrl:      picset.NothumbUrl,
		ParentOfDirpath: fullParentName,
	}

	this.Data["json"] = resp

	this.ServeJSON()

	// this.TplName = "picset/view.tpl"
}

func convertToPicsetItemList(picsetList []*picset.Picset) (list []PicsetItem) {
	rand.Seed(time.Now().UnixNano())
	list = make([]PicsetItem, len(picsetList))
	for i := range picsetList {
		list[i].FileName = picsetList[i].Name
		list[i].Thumb = picsetList[i].Thumb
		list[i].Url = picsetList[i].Url
		list[i].IsDir = picsetList[i].IsDir
		list[i].FileSize = utils.Size4Human(picsetList[i].Size)
		list[i].EditTime = picsetList[i].Mtime
		//FOR DEBUG in npm run dev
		if list[i].Thumb != "" {
			list[i].Thumb = "http://localhost:6699" + picsetList[i].Thumb
		}
		//FOR DEBUG END
	}
	return
}

func (this *PicsetController) ClearCache() {
	picset.ClearCache()
	logs.Debug("picset cache cleared")
	this.EnableRender = false
}
