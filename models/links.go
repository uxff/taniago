package models

import (
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego/logs"
)

type FriendLinks map[string][]struct{
		Name string `json:"name"`
		Url  string `json:"url"`
	}


var theLinks FriendLinks

func GetFriendLinks() *FriendLinks {
	return &theLinks
}

func LoadFriendLinksFromFile(f string) *FriendLinks {
	fhandle, err := os.Open(f)
	if err != nil {
		logs.Error("load friend links from %s error:%v", f, err)
		return nil
	}

	defer fhandle.Close()

	content, err := ioutil.ReadAll(fhandle)
	if err != nil {
		logs.Error("load friend links from %s error:%v", f, err)
		return nil
	}

	err = json.Unmarshal(content, &theLinks)
	if err != nil {
		logs.Error("load friend links from %s error:%v", f, err)
		return nil
	}

	logs.Info("load friend links from file %s ok", f)

	return &theLinks
}


