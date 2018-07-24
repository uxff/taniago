package controllers

import (
	"github.com/uxff/taniago/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {

	theLinks := models.GetFriendLinks()

	this.Data["thelinks"] = theLinks

	this.TplName = "index/index.tpl"
}

