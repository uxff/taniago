package picset

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/buger/jsonparser"
)

type Picset struct {
	Dirpath     string `json:"-"`
	Name        string `json:"name"`
	Thumb       string `json:"thumb"`
	Url         string `json:"url"` //for view tpl mode, beego format this, vue us another format
	Size        int64  `json:"size"`
	ThumbLoaded bool   `json:"-"`
	IsDir       bool   `json:"-"`
	Mtime       string `json:"mtime"`
}

var localDirRoot = "./"

//var fsRoute = "/fs"
const NothumbUrl = "/static/images/nothumb.png"

//var picsetRoute = "/picset"
//var pageSize = 8
var dirCache map[string][]*Picset

func init() {
	dirCache = make(map[string][]*Picset, 0)

	// 定时清理缓存
	go func() {
		tick := time.Tick(time.Second * 300)
		for {
			select {
			case <-tick:
				ClearCache()
			}
		}
	}()
}

func GetThumbOfDir(dirpath, preRoute string) string {
	if _, err := os.Stat(localDirRoot + "/" + dirpath + "/thumb.jpg"); err == nil {
		return preRoute + "/" + dirpath + "/thumb.jpg"
	}

	if _, err := os.Stat(localDirRoot + "/" + dirpath + "/thumb.png"); err == nil {
		return preRoute + "/" + dirpath + "/thumb.png"
	}

	return ""
}

func GetThumbFromSubdirs(dirpath, preRoute string, deep int) string {

	if deep <= 0 {
		return ""
	}

	dirHandle, err := ioutil.ReadDir(localDirRoot + "/" + dirpath)
	if err != nil {
		logs.Warn("cannot open this dir:%s :%v", dirpath, err)
		return ""
	}

	for _, fi := range dirHandle {
		if fi.IsDir() {
			if subThumb := GetThumbOfDir(dirpath+"/"+fi.Name(), preRoute); subThumb != "" {
				logs.Debug("dir %s sub %s HAS thumb:%s", dirpath, fi.Name(), subThumb)
				return subThumb
			}

			// logs.Debug("dir %s sub %s has NO thumb, try subdirs, deep=%d", dirpath, fi.Name(), deep)
			if subThumb := GetThumbFromSubdirs(dirpath+"/"+fi.Name(), preRoute, deep-1); subThumb != "" {
				return subThumb
			}
		}
	}

	logs.Warn("this dir has NO thumb:%s", dirpath)
	return ""
}

func GetTitleOfDir(dirpath, defaultName string) string {
	if fhandle, err := os.Open(localDirRoot + "/" + dirpath + "/config.json"); err == nil {
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
func GetPicsetListFromDir(dirpath, dirPreRoute, filePreRoute string) []*Picset {
	// dirCache = localDirPath=>[]*Picset
	//logs.Info("in GetPicsetListFromDir, dirpath=%s", dirpath)

	if dirpath == "" {
		//return nil
	}

	//curDirName := path.Base(dirpath)
	//parentDirName := path.Dir(dirpath)

	// set last char '/'
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

	dirHandle, err := ioutil.ReadDir(localDirRoot + "/" + dirpath)
	if err != nil {
		logs.Warn("open dir %s error:%v", dirpath, err)
		return nil
	}

	theDirList := make(PicsetSlice, 0)
	picIdx := 0
	//allNum := len(dirHandle)

	for _, fi := range dirHandle {

		lName := strings.ToLower(fi.Name())
		if len(fi.Name()) > 0 && fi.Name()[0] == '.' {
			continue
		}

		if fi.IsDir() {
			// if lName == "thumbs" || lName == "thumb" {
			// 	continue
			// }

			// 目录 该目录下如果有封面，选出封面
			//thumbPath := GetThumbOfDir(dirpath+fi.Name(), filePreRoute)
			//logs.Info("fi.name=%v thumb path=%v", fi.Name(), thumbPath)

			dirTitle := fi.Name() //getTitleOfDir(dirpath+fi.Name(), fi.Name()),//+fi.Name()+fmt.Sprintf("(%d/%d)", i+1, allNum)

			picItem := &Picset{
				Dirpath: dirpath + fi.Name(),
				Name:    dirTitle,
				Url:     dirPreRoute + "/" + dirpath + fi.Name(),
				Thumb:   "",
				IsDir:   true,
				Size:    fi.Size(),
				Mtime:   fi.ModTime().Format("2006-01-02 15:04"),
			}

			if thumbPath := GetThumbOfDir(dirpath+fi.Name(), filePreRoute); thumbPath != "" {
				picItem.Thumb = thumbPath
				picItem.ThumbLoaded = true
			}
			go picItem.LoadThumb(filePreRoute)

			theDirList = append(theDirList, picItem)

		} else {
			// if lName == "thumb.jpg" || lName == "thumb.png" || lName == "thumb.gif" {
			// 	continue
			// }

			picIdx++

			// 只有图片才展示
			fExt := path.Ext(lName)
			thumbPath := dirpath + fi.Name()
			picItem := &Picset{
				Dirpath: dirpath + fi.Name(),
				//Name:    fmt.Sprintf("%s-%d", curDirName, picIdx), //fmt.Sprintf("%s-%d", getTitleOfDir(dirpath, curDirName), picIdx),//
				Name:  fi.Name(),
				Thumb: "",
				Url:   filePreRoute + "/" + dirpath + fi.Name(),
				Size:  fi.Size(),
				Mtime: fi.ModTime().Format("2006-01-02 15:04"),
			}
			if fExt == ".jpg" || fExt == ".png" || fExt == ".gif" {
				picItem.Thumb = filePreRoute + "/" + thumbPath
				// picItem.Url = filePreRoute + "/" + thumbPath
			}
			theDirList = append(theDirList, picItem)
		}
	}

	sort.Sort(theDirList)

	dirCache[dirpath] = theDirList
	logs.Debug("path %s is loaded into cache", dirpath)

	return theDirList
}

func ClearCache() {
	dirCache = make(map[string][]*Picset, 0)
}

type PicsetSlice []*Picset

func (ps PicsetSlice) Len() int {
	return len(ps)
}

func (ps PicsetSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// 对子目录排序
func (ps PicsetSlice) Less(i, j int) bool {

	iName, jName := path.Base(ps[i].Dirpath), path.Base(ps[j].Dirpath)
	if len(iName) > 0 && ('0' <= iName[0] && iName[0] <= '9') &&
		len(jName) > 0 && ('0' <= jName[0] && jName[0] <= '9') {

		in, jn := 0, 0
		fmt.Sscanf(iName, "%d", &in)
		fmt.Sscanf(jName, "%d", &jn)

		return in < jn
	}

	return iName < jName
}

// 有瑕疵，会扫描券目录
func (p *Picset) LoadThumb(preRoute string) {

	if p.ThumbLoaded {
		return
	}

	if thumbPath := GetThumbFromSubdirs(p.Dirpath, preRoute, 1); thumbPath != "" {
		p.Thumb = thumbPath
		p.ThumbLoaded = true
	}
}
