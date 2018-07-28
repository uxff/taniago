package models

import (
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego/logs"
)

type FriendLinks []struct{
	Name string `json:"name"`
	Links[]struct{
		Name string `json:"name"`
		Url  string `json:"url"`
	}	`json:"links"`
}


var theLinks FriendLinks
var theLinksPath string

func GetFriendLinks() FriendLinks {
	return theLinks
}

func SetLinksPath(p string) {
	theLinksPath = p
	//LoadFriendLinksFromFile(theLinksPath)
}

func LoadFriendLinks() FriendLinks {
	return LoadFriendLinksFromFile(theLinksPath)
}

func LoadFriendLinksFromFile(f string) FriendLinks {
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

	return theLinks
}


