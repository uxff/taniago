/*
	图片化浏览某个目录
	如果浏览目录dir1,则需要目录${dir1}支持或提供以下特性：
		- ${dir1}/thumb.jpg		# 用于封面图 可以是.jpg,.png,.gif
		- ${dir1}/thubs/		# 用于存放原图对应的缩略图 计划中
		- 目录名称当成图集名称
*/
package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"path"

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

var staticRoot = "/55156.com-replicate"
var localDirRoot = "R:/themedia"
var fsRoute = "/fs"
var nothumbUrl = "/static/images/nothumb.png"
var picsetRoute = "/picset"
var pageSize = 8
var dirCache map[string][]*Picset


type IndexController struct {
	beego.Controller
}

func init() {
	dirCache = make(map[string][]*Picset, 0)
}

func (this *IndexController) Index() {

	this.TplName = "index/index.html"
}

// picset list
// todo: cache, multi domain, user login, pay and access, advertise
func (this *IndexController) Picset() {

	//
	fullDirName := this.Ctx.Input.Param(":splat")

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

	this.TplName = "picset/view.html"
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
		//this.Ctx.ResponseWriter.WriteHeader(404)
		//this.Ctx.ResponseWriter.Write([]byte("open dir error:"+err.Error()))
		return nil
	}


	theDirList := make([]*Picset, 0)
	picIdx := 0
	//allNum := len(dirHandle)

	for _, fi := range dirHandle {

		if fi.IsDir() {
			if fi.Name() == "thumbs" {
				continue
			}

			// 目录 该目录下如果有封面，选出封面
			thumbPath := getThumbOfDir(dirpath+fi.Name(), fsRoute)
			//logs.Info("fi.name=%v thumb path=%v", fi.Name(), thumbPath)

			theDirList = append(theDirList, &Picset{
				Dirpath:dirpath+"/"+fi.Name(),
				Name:"[DIR]"+fi.Name(),//getTitleOfDir(dirpath+fi.Name(), fi.Name()),//+fi.Name()+fmt.Sprintf("(%d/%d)", i+1, allNum),
				Thumb:thumbPath,
				Url:picsetRoute+"/"+dirpath+fi.Name(),
			})

		} else {
			if fi.Name() == "thumb.jpg" || fi.Name() == "thumb.png" || fi.Name()=="thumb.gif" {
				continue
			}

			picIdx++

			// 只有图片才展示
			fExt := path.Ext(fi.Name())
			if fExt == ".jpg" || fExt == ".png" || fExt == ".gif" {
				thumbPath := dirpath+fi.Name()
				theDirList = append(theDirList, &Picset{
					Dirpath:dirpath+"/"+fi.Name(),
					Name:fmt.Sprintf("%s-%d", curDirName, picIdx),
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
