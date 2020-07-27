package zhihu

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

type List struct {
	Paging struct {
		Total int `json:"total"`
	} `json:"paging"`
	Data []Item `json:"data"`
}

type Item struct {
	ViewCount      int    `json:"view_count"`
	FollowersCount int    `json:"followers_count"`
	IsFollowing    bool   `json:"is_following"`
	Title          string `json:"title"`
	Introduction   string `json:"introduction"`
	SectionList    []struct {
		SectionID    string `json:"section_id"`
		SectionTitle string `json:"section_title"`
	} `json:"section_list"`
	Banner  string `json:"banner"`
	Updated int    `json:"updated"`
	ID      string `json:"id"`
}
type CreateTopicArgs struct {
	Title     string `gorm:"type:varchar(100);column:title" form:"title"`
	ThemeID   uint64 `json:"theme_id" form:"theme_id" binding:"required"`
	Content   string `json:"content" form:"content" binding:"required"`
	TopicType int    `json:"topic_type" form:"topic_type"`
	Covers    []struct {
		Url   string `json:"url"`
		Title string `json:"title"`
		Type  int    `json:"type"`
	} `json:"covers" form:"covers" binding:"required"`
}

var (
	client = resty.New()
)

func CreateTopicWithImageFromZhihu() {
	r := client.R()
	resp, err := r.Get("https://www.zhihu.com/api/v4/news_specials/list?limit=100&offset=620")
	if err != nil {
		fmt.Printf("request zhihu fail:%s", err.Error())
		return
	}

	var list List
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		fmt.Printf("resp:%s", string(resp.Body()))
		fmt.Printf("json unmarshal zhihu fail:%s", err.Error())
		return
	}

	fmt.Printf("count:%d list.length:%d", list.Paging.Total, len(list.Data))

	for _, item := range list.Data {

		err = CreateTopic(CreateTopicArgs{
			Title:     item.Title,
			ThemeID:   1,
			Content:   item.Introduction,
			TopicType: 1,
			Covers: []struct {
				Url   string `json:"url"`
				Title string `json:"title"`
				Type  int    `json:"type"`
			}{
				{Url: item.Banner, Title: item.Title, Type: 2},
			},
		})

		if err != nil {
			continue
		}

	}

}

func CreateTopic(args CreateTopicArgs) error {
	req := client.R()
	req.SetBody(args)
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Connection", "keep-alive")
	req.SetHeader("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjAsImF2YXRhciI6IiIsImVtYWlsIjoiIiwiZXhwIjoxNTkyMDM0MTI3LCJpZCI6IjMiLCJsZXZlbCI6MCwibmlja25hbWUiOiIxNjYwMTE2MzU1MyIsIm9yaWdfaWF0IjoxNTg5NDQyMTI3LCJwYXNzd29yZCI6IiIsInBob25lIjoiMTY2MDExNjM1NTMiLCJzZXgiOjB9.-dpftEiZDurOpdo8Ou7uwx4KQ3UgY3ghuA7MbbZiZH4")
	resp, err := req.Post("http://api.antdate.cn/topics")
	if err != nil {
		fmt.Printf("Create fail:%s \n", err.Error())
		return err
	}
	fmt.Printf("create success:%s", resp.Status())
	return nil
}

type DouYinHot struct {
	BillboardData []struct {
		Author string `json:"author"`
		ImgURL string `json:"img_url"`
		Link   string `json:"link"`
		Rank   int    `json:"rank"`
		Title  string `json:"title"`
		Value  string `json:"value"`
	} `json:"billboard_data"`
	Extra struct {
		Now int64 `json:"now"`
	} `json:"extra"`
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func CreateTopicWithVideoFormDouYin() {
	r := client.R()
	r.SetHeader("cookie", "_ga=GA1.2.1413643649.1591342323; _gid=GA1.2.162417155.1591342323; SLARDAR_WEB_ID=e6960d96-220e-47e6-a777-f5de0bb428d6; ttwid=6834763348758119943; passport_csrf_token=00f7e95dc924ecbe3e89dc31989e8866; s_v_web_id=kb1w5dyz_UgnH7AlM_YVkk_40HB_9vGL_E62kQHiPf34b; csrf_token=WAHetzqALWMirgIBtAiuAGsWEnBEbEIa")
	r.SetHeader("referer", "https://creator.douyin.com/data/billboard/hot")
	//r.SetHeader(":path", "/aweme/v1/creator/data/billboard/?billboard_type=4&_signature=2R63RwAAAACvMR6nvNH7F9ket1AAIfP")

	resp, err := r.Get("https://creator.douyin.com/aweme/v1/creator/data/billboard/?billboard_type=4&_signature=2R63RwAAAACvMR6nvNH7F9ket1AAIfP")
	if err != nil {
		fmt.Printf("request zhihu fail:%s", err.Error())
		return
	}

	fmt.Printf("resp:%s \n", string(resp.Body()))
	var hotList DouYinHot
	err = json.Unmarshal(resp.Body(), &hotList)
	if err != nil {
		fmt.Printf("json unmarshal zhihu fail:%s", err.Error())
		return
	}

	for _, item := range hotList.BillboardData {
		videoID := strings.Split(strings.Trim(item.Link, "https://www.iesdouyin.com/share/video/"), "/")[0]
		videoAddr := GetVideoUrl(videoID)
		fmt.Printf("title:%s video id:%s addr:%s \n", item.Title, videoID, videoAddr)

		err = CreateTopic(CreateTopicArgs{
			Title:     item.Title,
			ThemeID:   1,
			Content:   item.Title,
			TopicType: 1,
			Covers: []struct {
				Url   string `json:"url"`
				Title string `json:"title"`
				Type  int    `json:"type"`
			}{
				{Url: videoAddr, Title: item.Title, Type: 1},
			},
		})

		if err != nil {
			continue
		}
	}

	fmt.Printf("list.length:%d", len(hotList.BillboardData))

}

type ShareVideo struct {
	StatusCode int `json:"status_code"`
	ItemList   []struct {
		GroupID      int64       `json:"group_id"`
		Desc         string      `json:"desc"`
		Position     interface{} `json:"position"`
		LabelTopText interface{} `json:"label_top_text"`
		VideoText    interface{} `json:"video_text"`
		IsPreview    int         `json:"is_preview"`
		Video        struct {
			OriginCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"origin_cover"`
			Duration int `json:"duration"`
			PlayAddr struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
			DynamicCover struct {
				URLList []string `json:"url_list"`
				URI     string   `json:"uri"`
			} `json:"dynamic_cover"`
			Width        int         `json:"width"`
			Ratio        string      `json:"ratio"`
			HasWatermark bool        `json:"has_watermark"`
			BitRate      interface{} `json:"bit_rate"`
			Vid          string      `json:"vid"`
			Cover        struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover"`
			Height int `json:"height"`
		} `json:"video"`
		ShareInfo struct {
			ShareTitle     string `json:"share_title"`
			ShareWeiboDesc string `json:"share_weibo_desc"`
			ShareDesc      string `json:"share_desc"`
		} `json:"share_info"`
		AuthorUserID int64         `json:"author_user_id"`
		TextExtra    []interface{} `json:"text_extra"`
		Geofencing   interface{}   `json:"geofencing"`
		CreateTime   int           `json:"create_time"`
		Author       struct {
			ShortID      string `json:"short_id"`
			Signature    string `json:"signature"`
			AvatarMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_medium"`
			PolicyVersion    interface{} `json:"policy_version"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			Geofencing       interface{} `json:"geofencing"`
			UID              string      `json:"uid"`
			Nickname         string      `json:"nickname"`
			AvatarLarger     struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_larger"`
			AvatarThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			UniqueID        string      `json:"unique_id"`
			FollowersDetail interface{} `json:"followers_detail"`
		} `json:"author"`
		ChaList   interface{} `json:"cha_list"`
		ShareURL  string      `json:"share_url"`
		RiskInfos struct {
			Content string `json:"content"`
			Warn    bool   `json:"warn"`
			Type    int    `json:"type"`
		} `json:"risk_infos"`
		AwemeType   int         `json:"aweme_type"`
		CommentList interface{} `json:"comment_list"`
		Promotions  interface{} `json:"promotions"`
		AwemeID     string      `json:"aweme_id"`
		Music       struct {
			CoverLarge struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_large"`
			Duration   int         `json:"duration"`
			Position   interface{} `json:"position"`
			CoverThumb struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_thumb"`
			PlayURL struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_url"`
			ID      int64  `json:"id"`
			Mid     string `json:"mid"`
			Title   string `json:"title"`
			Author  string `json:"author"`
			CoverHd struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_hd"`
			CoverMedium struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover_medium"`
			Status int `json:"status"`
		} `json:"music"`
		Duration   int `json:"duration"`
		Statistics struct {
			AwemeID      string `json:"aweme_id"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
		} `json:"statistics"`
		UniqidPosition interface{} `json:"uniqid_position"`
		ImageInfos     interface{} `json:"image_infos"`
		LongVideo      interface{} `json:"long_video"`
		IsLiveReplay   bool        `json:"is_live_replay"`
		VideoLabels    interface{} `json:"video_labels"`
	} `json:"item_list"`
	Extra struct {
		Now   int64  `json:"now"`
		Logid string `json:"logid"`
	} `json:"extra"`
}

func GetVideoUrl(videoID string) string {
	r := client.R()

	resp, err := r.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=%s&dytk=d05e31028de46bed850eae38ee0e09e1e1996c62aa93ce792f6ef7d9af656adc", videoID))
	if err != nil {
		fmt.Printf("request zhihu fail:%s", err.Error())
		return ""
	}

	var hv ShareVideo
	err = json.Unmarshal(resp.Body(), &hv)
	if err != nil {
		fmt.Printf("json unmarshal video fail:%s", err.Error())
		return ""
	}

	return hv.ItemList[0].Video.PlayAddr.URLList[0]

}
