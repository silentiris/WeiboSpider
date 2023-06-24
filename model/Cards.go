package model

type CardList struct {
	Ok   int `json:"ok"`
	Data struct {
		CardlistInfo struct {
			Containerid       string `json:"containerid"`
			VP                int    `json:"v_p"`
			ShowStyle         int    `json:"show_style"`
			Total             int    `json:"total"`
			AutoLoadMoreIndex int    `json:"autoLoadMoreIndex"`
			SinceID           int64  `json:"since_id"`
		} `json:"cardlistInfo"`
		Cards []struct {
			CardType      int    `json:"card_type"`
			ProfileTypeID string `json:"profile_type_id"`
			Itemid        string `json:"itemid"`
			Scheme        string `json:"scheme"`
			Mblog         struct {
				Visible struct {
					Type   int `json:"type"`
					ListID int `json:"list_id"`
				} `json:"visible"`
				CreatedAt                string   `json:"created_at"`
				ID                       string   `json:"id"`
				Mid                      string   `json:"mid"`
				CanEdit                  bool     `json:"can_edit"`
				ShowAdditionalIndication int      `json:"show_additional_indication"`
				Text                     string   `json:"text"`
				TextLength               int      `json:"textLength"`
				Source                   string   `json:"source"`
				Favorited                bool     `json:"favorited"`
				PicIds                   []string `json:"pic_ids"`
				IsPaid                   bool     `json:"is_paid"`
				MblogVipType             int      `json:"mblog_vip_type"`
				User                     struct {
					ID                int    `json:"id"`
					ScreenName        string `json:"screen_name"`
					ProfileImageURL   string `json:"profile_image_url"`
					ProfileURL        string `json:"profile_url"`
					StatusesCount     int    `json:"statuses_count"`
					Verified          bool   `json:"verified"`
					VerifiedType      int    `json:"verified_type"`
					VerifiedTypeExt   int    `json:"verified_type_ext"`
					VerifiedReason    string `json:"verified_reason"`
					CloseBlueV        bool   `json:"close_blue_v"`
					Description       string `json:"description"`
					Gender            string `json:"gender"`
					Mbtype            int    `json:"mbtype"`
					Svip              int    `json:"svip"`
					Urank             int    `json:"urank"`
					Mbrank            int    `json:"mbrank"`
					FollowMe          bool   `json:"follow_me"`
					Following         bool   `json:"following"`
					FollowCount       int    `json:"follow_count"`
					FollowersCount    string `json:"followers_count"`
					FollowersCountStr string `json:"followers_count_str"`
					CoverImagePhone   string `json:"cover_image_phone"`
					AvatarHd          string `json:"avatar_hd"`
					Like              bool   `json:"like"`
					LikeMe            bool   `json:"like_me"`
					Badge             struct {
						VipActivity2         int `json:"vip_activity2"`
						FollowWhitelistVideo int `json:"follow_whitelist_video"`
						UserNameCertificate  int `json:"user_name_certificate"`
						Wenchuan10Th         int `json:"wenchuan_10th"`
						NationalDay2018      int `json:"national_day_2018"`
						StatusVisible        int `json:"status_visible"`
						China2019            int `json:"china_2019"`
						Hongbao2020          int `json:"hongbao_2020"`
						Feiyan2020           int `json:"feiyan_2020"`
						PcNew                int `json:"pc_new"`
						Hongrenjie2020       int `json:"hongrenjie_2020"`
						Hongbaofeifuniu2021  int `json:"hongbaofeifuniu_2021"`
						Hongbaofeijika2021   int `json:"hongbaofeijika_2021"`
						Yuanlongping2021     int `json:"yuanlongping_2021"`
						Gaokao2021           int `json:"gaokao_2021"`
						PartyCardidState     int `json:"party_cardid_state"`
						Hongbaofei20222021   int `json:"hongbaofei2022_2021"`
						VideoVisible         int `json:"video_visible"`
						Iplocationchange2022 int `json:"iplocationchange_2022"`
						Guoqi2022            int `json:"guoqi_2022"`
						Gangqi2022           int `json:"gangqi_2022"`
						Gaokao2023           int `json:"gaokao_2023"`
					} `json:"badge"`
				} `json:"user"`
				RepostsCount          int           `json:"reposts_count"`
				CommentsCount         int           `json:"comments_count"`
				ReprintCmtCount       int           `json:"reprint_cmt_count"`
				AttitudesCount        int           `json:"attitudes_count"`
				PendingApprovalCount  int           `json:"pending_approval_count"`
				IsLongText            bool          `json:"isLongText"`
				Mlevel                int           `json:"mlevel"`
				ShowMlevel            int           `json:"show_mlevel"`
				DarwinTags            []interface{} `json:"darwin_tags"`
				Mblogtype             int           `json:"mblogtype"`
				Rid                   string        `json:"rid"`
				Cardid                string        `json:"cardid"`
				ExternSafe            int           `json:"extern_safe"`
				NumberDisplayStrategy struct {
					ApplyScenarioFlag    int    `json:"apply_scenario_flag"`
					DisplayTextMinNumber int    `json:"display_text_min_number"`
					DisplayText          string `json:"display_text"`
				} `json:"number_display_strategy"`
				ContentAuth       int `json:"content_auth"`
				SafeTags          int `json:"safe_tags"`
				CommentManageInfo struct {
					CommentPermissionType int `json:"comment_permission_type"`
					ApprovalCommentType   int `json:"approval_comment_type"`
					CommentSortType       int `json:"comment_sort_type"`
				} `json:"comment_manage_info"`
				PicNum  int `json:"pic_num"`
				HotPage struct {
					Fid            string `json:"fid"`
					FeedDetailType int    `json:"feed_detail_type"`
				} `json:"hot_page"`
				NewCommentStyle   int         `json:"new_comment_style"`
				AbSwitcher        int         `json:"ab_switcher"`
				RegionName        string      `json:"region_name"`
				RegionOpt         int         `json:"region_opt"`
				PicBg             string      `json:"pic_bg"`
				PicBgType         int         `json:"pic_bg_type"`
				PicBgBiz          interface{} `json:"pic_bg_biz"`
				MblogMenuNewStyle int         `json:"mblog_menu_new_style"`
				EditConfig        struct {
					Edited bool `json:"edited"`
				} `json:"edit_config"`
				Bid string `json:"bid"`
			} `json:"mblog,omitempty"`
		} `json:"cards"`
		Scheme      string `json:"scheme"`
		ShowAppTips int    `json:"showAppTips"`
	} `json:"data"`
}
