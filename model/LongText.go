package model

type LongText struct {
	Ok   int `json:"ok"`
	Data struct {
		Ok              int    `json:"ok"`
		LongTextContent string `json:"longTextContent"`
		RepostsCount    int    `json:"reposts_count"`
		CommentsCount   int    `json:"comments_count"`
		AttitudesCount  int    `json:"attitudes_count"`
	} `json:"data"`
}
