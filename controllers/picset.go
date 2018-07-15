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
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"path"
	"strings"

	"github.com/astaxie/beego"
	"github.com/uxff/taniago/utils/paginator"
	"github.com/astaxie/beego/logs"
	"github.com/buger/jsonparser"
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

var localDirRoot = "./"
var fsRoute = "/fs"
var nothumbUrl = "/static/images/nothumb.png"
var picsetRoute = "/picset"
var pageSize = 8
var dirCache map[string][]*Picset


type PicsetController struct {
	BaseController
}

func init() {
	dirCache = make(map[string][]*Picset, 0)
}


// picset list
func (this *PicsetController) Picset() {

	this.Data["appname"] = beego.AppConfig.String("appname")

	//
	fullDirName := this.Ctx.Input.Param(":splat")
	c := this.GetString("c")
	switch c {
	case "clearcache":
		dirCache = make(map[string][]*Picset, 0)
		logs.Debug("dir cache cleared")
	}

	//logs.Info("fullDirName from param=%v", fullDirName)

	curDirName := path.Base(fullDirName)
	fullParentName := path.Dir(fullDirName)

	this.Data["curDirName"] = curDirName
	this.Data["fullParentName"] = fullParentName
	this.Data["parentLink"] = picsetRoute+"/"+fullParentName

	//logs.Info("fullDirName from url param:%s curDirName=%s parentName=%s", fullDirName, curDirName, fullParentName )
	//dirpath := localDirRoot+"/"+fullDirName

	//thedirnames := make([]string, 0)
	theDirList := GetPicsetListFromDir(fullDirName)

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

func getThumbOfDir(dirpath, preRoute string) string {
	if _, err := os.Stat(localDirRoot+"/"+dirpath + "/thumb.jpg"); err == nil {
		return preRoute +"/"+ dirpath+"/thumb.jpg"
	}

	if _, err := os.Stat(localDirRoot+"/"+dirpath+"/thumb.png"); err == nil {
		return preRoute+"/"+dirpath+"/thumb.png"
	}

	return nothumbUrl
}

func getTitleOfDir(dirpath, defaultName string) string {
	if fhandle, err := os.Open(localDirRoot+"/"+dirpath + "/config.json"); err == nil {
		//logs.Info("open %v", dirpath)
		defer fhandle.Close()
		content, rerr := ioutil.ReadAll(fhandle)
		if rerr == nil {
			//logs.Info("read %s", content)
			title, _ := jsonparser.GetString(content, "title")
			return title
		}
	}

	//logs.Warn("cannot open:%s", localDirRoot+"/"+dirpath + "/config.json")

	return defaultName
}

func SetLocalDirRoot(dir string) {
	localDirRoot = dir
}

/**
	dirpath must under localRoot
*/
func GetPicsetListFromDir(dirpath string) []*Picset {
	// dirCache = localDirPath=>[]*Picset
	//logs.Info("in GetPicsetListFromDir, dirpath=%s", dirpath)

	if dirpath == "" {
		//return nil
	}

	curDirName := path.Base(dirpath)
	//parentDirName := path.Dir(dirpath)

	if len(dirpath) > 0 && dirpath[len(dirpath)-1] != '/' {
		dirpath = dirpath + "/"
	}


	// return if exist
	if existPicsetList, ok := dirCache[dirpath]; ok {
		if existPicsetList != nil {
			return existPicsetList
		}
	}

	//logs.Info("not exist:%s", dirpath)

	dirHandle, err := ioutil.ReadDir(localDirRoot+"/"+dirpath)
	if err != nil {
		logs.Warn("open dir %s error:%v", dirpath, err)
		return nil
	}


	theDirList := make([]*Picset, 0)
	picIdx := 0
	//allNum := len(dirHandle)

	for _, fi := range dirHandle {

		lName := strings.ToLower(fi.Name())
		if fi.IsDir() {
			if lName == "thumbs" {
				continue
			}

			// 目录 该目录下如果有封面，选出封面
			thumbPath := getThumbOfDir(dirpath+fi.Name(), fsRoute)
			//logs.Info("fi.name=%v thumb path=%v", fi.Name(), thumbPath)

			dirTitle := fi.Name()//getTitleOfDir(dirpath+fi.Name(), fi.Name()),//+fi.Name()+fmt.Sprintf("(%d/%d)", i+1, allNum)

			theDirList = append(theDirList, &Picset{
				Dirpath:dirpath+"/"+fi.Name(),
				Name:"[DIR]"+dirTitle,
				Thumb:thumbPath,
				Url:picsetRoute+"/"+dirpath+fi.Name(),
			})

		} else {
			if lName == "thumb.jpg" || lName == "thumb.png" || lName=="thumb.gif" {
				continue
			}

			picIdx++

			// 只有图片才展示
			fExt := path.Ext(lName)
			if fExt == ".jpg" || fExt == ".png" || fExt == ".gif" {
				thumbPath := dirpath+fi.Name()
				theDirList = append(theDirList, &Picset{
					Dirpath:dirpath+"/"+fi.Name(),
					Name:fmt.Sprintf("%s-%d", curDirName, picIdx),//fmt.Sprintf("%s-%d", getTitleOfDir(dirpath, curDirName), picIdx),//
					Thumb:fsRoute+"/"+thumbPath,
					Url:fsRoute+"/"+thumbPath,
				})

			}

		}

	}

	dirCache[dirpath] = theDirList
	logs.Info("path %s is loaded into cache", dirpath)

	return theDirList
}
