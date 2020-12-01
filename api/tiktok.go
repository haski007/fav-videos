package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Haski007/fav-videos/internal/fvb/config"
	"github.com/Haski007/fav-videos/internal/fvb/repository"
	"github.com/valyala/fasthttp"
	"net/http"
)

func GetLikedVideos(secUserID string, count int) ([]repository.Video, error) {
	req := &fasthttp.Request{}
	res := &fasthttp.Response{}

	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI("https://api16-normal-c-alisg.tiktokv.com/aweme/v1/aweme/favorite/?" +
		fmt.Sprintf("aid=%d&device_id=%d&sec_user_id=%s&count=%d",
			config.Aid, 1000000000+config.SeededRand.Intn(1000000000), secUserID, count))
	req.Header.SetUserAgent(config.UserAgent)

	err := fasthttp.Do(req, res)
	if err != nil {
		return nil, err
	}

	var tiktokResp repository.AwemeFavoriteResponse
	err = json.Unmarshal(res.Body(), &tiktokResp)
	if err != nil {
		return nil, err
	}

	if tiktokResp.StatusMessage != "" {
		return nil, errors.New(tiktokResp.StatusMessage)
	}

	var videos []repository.Video
	for _, v := range tiktokResp.AwemeList {
		if len(v.Video.PlayAddr.URLList) == 0 {
			continue
		}
		videos = append(videos, repository.Video{
			ID:          v.ID,
			ShareURL:    v.ShareURL,
			DownloadURL: v.Video.PlayAddr.URLList[0],
		})
	}

	for i, j := 0, len(videos)-1; i < j; i, j = i+1, j-1 {
		videos[i], videos[j] = videos[j], videos[i]
	}

	return videos, nil
}
