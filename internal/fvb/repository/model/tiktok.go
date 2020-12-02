package model

type AwemeFavoriteResponse struct {
	StatusMessage string `json:"status_msg"`
	AwemeList     []struct {
		ID       string `json:"aweme_id"`
		ShareURL string `json:"share_url"`
		Video    struct {
			PlayAddr struct {
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
		} `json:"video"`
	} `json:"aweme_list"`
}
