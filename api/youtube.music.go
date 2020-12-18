package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Haski007/fav-music-bot/internal/fmb/config"
	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/model"
	"github.com/valyala/fasthttp"
)

func GetLikedVideos(publisher *model.Publisher, count int) ([]model.Video, error) {
	req := &fasthttp.Request{}
	res := &fasthttp.Response{}

	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI("https://api16-normal-c-alisg.tiktokv.com/aweme/v1/aweme/favorite/?" +
		fmt.Sprintf("aid=%d&device_id=%d&sec_user_id=%s&count=%d",
			config.Aid, 1000000000+config.SeededRand.Intn(1000000000), publisher.SecureID, count))
	req.Header.SetUserAgent(config.UserAgent)

	err := fasthttp.Do(req, res)
	if err != nil {
		return nil, err
	}

	var tiktokResp model.AwemeFavoriteResponse
	err = json.Unmarshal(res.Body(), &tiktokResp)
	if err != nil {
		return nil, err
	}

	if tiktokResp.StatusMessage != "" {
		return nil, errors.New(tiktokResp.StatusMessage)
	}

	var videos []model.Video
	for _, v := range tiktokResp.AwemeList {
		if len(v.Video.PlayAddr.URLList) == 0 {
			continue
		}
		videos = append(videos, model.Video{
			ID:          v.ID,
			ShareURL:    v.ShareURL,
			DownloadURL: v.Video.PlayAddr.URLList[0],
			Publisher:   publisher.Username,
		})
	}

	for i, j := 0, len(videos)-1; i < j; i, j = i+1, j-1 {
		videos[i], videos[j] = videos[j], videos[i]
	}

	return videos, nil
}

func GetSecureUserID(username string) (string, error) {
	req := &fasthttp.Request{}
	res := &fasthttp.Response{}

	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI("https://www.tiktok.com/@" + username)
	req.Header.SetUserAgent(config.UserAgent)

	err := fasthttp.Do(req, res)
	if err != nil {
		return "", err
	}

	matches := config.SecUIDReg.FindStringSubmatch(res.String())
	if len(matches) != 2 {
		return "", errors.New("no matches")
	}

	return matches[1], nil
}
