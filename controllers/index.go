package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type Picset struct {
	Dirpath string
	Name string
	Thumb string
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {

	this.TplName = "index/index.html"
}

func (this *IndexController) Picset() {

	dirpath := "R:/themedia/55156.com-replicate"

	dirHandle, err := ioutil.ReadDir(dirpath)
	if err != nil {
		this.Ctx.ResponseWriter.Write([]byte("open dir error:"+err.Error()))
		return
	}

	thedirnames := make([]string, 0)

	for _, fi := range dirHandle {
		if fi.IsDir() {
			thedirnames = append(thedirnames, fi.Name())

		}
	}



	this.Data["thedirnames"] = thedirnames

	this.TplName = "picset/view.html"
}
