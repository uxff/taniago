package controllers

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	subdirname := this.Ctx.Input.Param(":splat")

	logs.Info("subdirname from url param:%s", subdirname)
	dirpath := localDirRoot+"/"+subdirname

	dirHandle, err := ioutil.ReadDir(dirpath)
	if err != nil {
		this.Ctx.ResponseWriter.Write([]byte("open dir error:"+err.Error()))
		return
	}

	//thedirnames := make([]string, 0)
	theDirList := make([]*Picset, 0)

	//
	if subdirname != "" && subdirname[len(subdirname)-1]!='/' {
		subdirname = subdirname+"/"
	}

	for _, fi := range dirHandle {
		if fi.IsDir() {

			thumbPath := this.getThumbOfDir(subdirname+fi.Name(), fsRoute)
			//logs.Info("fi.name=%v thumb path=%v", fi.Name(), thumbPath)


			theDirList = append(theDirList, &Picset{
				Dirpath:dirpath+"/"+fi.Name(),
				Name:"[DIR]"+fi.Name(),
				Thumb:thumbPath,
				Url:picsetRoute+"/"+subdirname+fi.Name(),
			})


		} else {
			if fi.Name() == "thumb.jpg" || fi.Name() == "thumb.png" {
				continue
			}

			fExt := path.Ext(fi.Name())
			if fExt == ".jpg" || fExt == ".png" || fExt == ".gif" {
				thumbPath := subdirname+fi.Name()
				theDirList = append(theDirList, &Picset{
					Dirpath:dirpath+"/"+fi.Name(),
					Name:fi.Name(),
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
