package tiksaver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

// from github.com/mehanon/tt_2ch_media

// the website support downloading tiktok videos:
// https://ssstik.io/en
// https://www.tikwm.com/api/
// https://dlpanda.com/
// https://savett.cc/en
//

// TiktokInfo there are more fields, tho I omitted unnecessary ones
type TiktokInfo struct {
	Id         string `json:"id"`
	Play       string `json:"play,omitempty"`
	Hdplay     string `json:"hdplay,omitempty"`
	CreateTime int64  `json:"create_time"`
	Author     struct {
		UniqueId string `json:"unique_id"`
	} `json:"author"`
}

type TikwmResponse struct {
	Code          int        `json:"code"`
	Msg           string     `json:"msg"`
	ProcessedTime float64    `json:"processed_time"`
	Data          TiktokInfo `json:"data,omitempty"`
}

func TikwnGetInfo(link string) (*TiktokInfo, error) {
	payload := url.Values{"url": {link}, "hd": {"1"}}
	r, err := http.PostForm("https://www.tikwm.com/api/", payload)
	if err != nil {
		return nil, err
	}
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var resp TikwmResponse
	err = json.Unmarshal(buffer, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return &resp.Data, nil
}

// DownloadTiktokTikwm downloads video from tikwm.com
func DownloadTiktokTikwm(link string) (outputfile, desc string, err error) {
	info, err := TikwnGetInfo(link)
	if err != nil {
		return "", "", err
	}

	desc = fmt.Sprintf("%+v", *info)

	var downloadUrl string
	if info.Hdplay != "" {
		downloadUrl = info.Hdplay
	} else if info.Play != "" {
		println("warning: tikwm couldn't find HD version, downloading how it is...")
		downloadUrl = info.Play
	} else {
		return "", desc, errors.New("no download links found :c")
	}

	localFilename := fmt.Sprintf(
		"%s_%s_%s.mp4",
		info.Author.UniqueId,
		time.Unix(info.CreateTime, 0).Format("2006-01-02"),
		info.Id,
	)

	err = Wget(downloadUrl, localFilename)
	if err != nil {
		return "", desc, err
	}

	return localFilename, desc, nil
}

func Wget(url string, filename string) error {
	_, err := grab.Get(filename, url)
	return err
}
