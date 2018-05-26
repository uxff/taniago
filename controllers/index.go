package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/utils/paginator"
)

type Picset struct {
	Dirpath string
	Name string
	Thumb string
	Url string
}

var staticRoot = "/55156.com-replicate"
var localDirRoot = "R:/themedia"
var fsRoute = "/fs"
var nothumbUrl = "/static/images/nothumb.png"
var picsetRoute = "/picset"
var pageSize = 10

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {

	this.TplName = "index/index.html"
}

// picset list
// todo: page
func (this *IndexController) Picset() {

	//
	fullDirName := this.Ctx.Input.Param(":splat")

	curDirName := path.Base(fullDirName)

	logs.Info("fullDirName from url param:%s curDirName=%s ", fullDirName, curDirName)
	dirpath := localDirRoot+"/"+fullDirName

	dirHandle, err := ioutil.ReadDir(dirpath)
	if err != nil {
		this.Ctx.ResponseWriter.Write([]byte("open dir error:"+err.Error()))
		return
	}

	//thedirnames := make([]string, 0)
	theDirList := make([]*Picset, 0)


	//
	if fullDirName != "" && fullDirName[len(fullDirName)-1]!='/' {
		fullDirName = fullDirName+"/"
	}

	allNum := len(dirHandle)

	p := paginator.NewPaginator(this.Ctx.Request, pageSize, int64(allNum))
	this.Data["paginator"] = p

	picIdx := 0

	for i, fi := range dirHandle {

		if i < (p.Page()-1)*pageSize || i>=p.Page()*pageSize {
			continue
		}

		if fi.IsDir() {

			// 目录 该目录下如果有封面，选出封面
			thumbPath := this.getThumbOfDir(fullDirName+fi.Name(), fsRoute)
			//logs.Info("fi.name=%v thumb path=%v", fi.Name(), thumbPath)

			theDirList = append(theDirList, &Picset{
				Dirpath:dirpath+"/"+fi.Name(),
				Name:"[DIR]"+fi.Name()+fmt.Sprintf("(%d/%d)", i+1, allNum),
				Thumb:thumbPath,
				Url:picsetRoute+"/"+fullDirName+fi.Name(),
			})

		} else {
			if fi.Name() == "thumb.jpg" || fi.Name() == "thumb.png" || fi.Name()=="thumb.gif" {
				continue
			}

			picIdx++

			// 只有图片才展示
			fExt := path.Ext(fi.Name())
			if fExt == ".jpg" || fExt == ".png" || fExt == ".gif" {
				thumbPath := fullDirName+fi.Name()
				theDirList = append(theDirList, &Picset{
					Dirpath:dirpath+"/"+fi.Name(),
					Name:fmt.Sprintf("%s-%d", curDirName, picIdx),
					Thumb:fsRoute+"/"+thumbPath,
					Url:fsRoute+"/"+thumbPath,
				})

			}

		}

	}


	this.Data["thedirnames"] = theDirList

	this.TplName = "picset/view.html"
}

func (this *IndexController) getThumbOfDir(path, preRoute string) string {
	if _, err := os.Stat(localDirRoot+"/"+path + "/thumb.jpg"); err == nil {
		return fsRoute +"/"+ path+"/thumb.jpg"
	}

	if _, err := os.Stat(localDirRoot+"/"+path+"/thumb.png"); err == nil {
		return fsRoute+"/"+path+"/thumb.png"
	}

	return nothumbUrl
}
