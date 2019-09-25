package controllers

import (
	"github.com/uxff/taniago/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {


	this.TplName = "index/index.tpl"
}

func (this *IndexController) Links() {

	theLinks := models.LoadIndexLinks()

	this.Data["thelinks"] = theLinks

	this.TplName = "index/links.tpl"
}