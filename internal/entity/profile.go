package entity

type (
	KakaoProfile struct {
		ID           int64 `json:"id"`
		KakaoAccount struct {
			Email   string `json:"email"`
			Profile struct {
				Nickname        string `json:"nickname"`
				ProfileImageURL string `json:"profile_image_url"`
			} `json:"profile"`
		} `json:"kakao_account"`
	}

	GoogleProfile struct {
		Sub     string `json:"sub"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
)
