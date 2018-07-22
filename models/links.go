package models

import (
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego/logs"
)

type FriendLinks struct {
	List []struct{
		Name string `json:"name"`
		Url  string `json:"url"`
	}				`json:"list"`
}

var theLinks FriendLinks

func GetFriendLinks() *FriendLinks {
	return &theLinks
}

func LoadFriendLinksFromFile(f string) {
	fhandle, err := os.Open(f)
	if err != nil {
		logs.Error("load friend links from %s error:%v", f, err)
		return
	}

	defer fhandle.Close()

	content, err := ioutil.ReadAll(fhandle)
	if err != nil {
		logs.Error("load friend links from %s error:%v", f, err)
		return
	}

	err = json.Unmarshal(content, &theLinks)
	if err != nil {
		logs.Error("load friend links from %s error:%v", f, err)
		return
	}

	logs.Info("load friend links from file %s ok", f)

}


