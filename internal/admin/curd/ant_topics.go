package curd

import (
	"database/sql"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/antdate/antdate-service/pkg/json"
	"github.com/sirupsen/logrus"
)

type TopicType string

const (
	ShortTopicType TopicType = "1"
	LongTopicType  TopicType = "2"
)

const (
	Video = 1
	Image = 2
)

type Cover struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Type  int    `json:"type"`
}

func GetAntTopicsTable(ctx *context.Context) table.Table {

	antTopics := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	listPage := antTopics.GetInfo().HideFilterArea().HideExportButton()
	listPage.AddField("ID", "id", db.Bigint).FieldFilterable()
	listPage.AddField("用户", "nickname", db.Bigint)
	listPage.AddField("Chat_id", "chat_id", db.Bigint).FieldHide()
	listPage.AddField("Title", "title", db.Varchar)
	listPage.AddField("话题板块", "theme_id", db.Varchar)
	listPage.AddField("照片", "covers", db.Varchar).FieldCarousel(func(value string) []string {
		covers := UnmarshalCovers([]byte(value))
		var slides []string
		for _, cover := range covers {
			slides = append(slides, cover.Url)
		}
		return slides
	}, 150, 100)
	listPage.AddField("浏览|在线", "", db.Bigint).FieldDisplay(func(value types.FieldModel) interface{} {
		return fmt.Sprintf("%v | %v", value.Row["views"], value.Row["online"])
	})
	listPage.AddField("话题类型", "topic_type", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		if TopicType(value.Value) == LongTopicType {
			return "长话题"
		}
		return "短话题"
	})
	//listPage.AddField("Covers", "covers", db.JSON)
	listPage.AddField("发布时间", "created_at", db.Datetime)
	listPage.SetTable("ant_topics").SetTitle("AntTopics").SetDescription("AntTopics")

	formList := antTopics.GetForm()
	formList.AddField("ID", "id", db.Bigint, form.Default)
	formList.AddField("话题类型", "topic_type", db.Int, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "短话题", Value: "1"},
		{Text: "长话题", Value: "2"},
	}). // 设置默认值
		FieldDefault("1")
	formList.AddField("发帖用户", "user_id", db.Bigint, form.SelectSingle).
		FieldOptionsFromTable("ant_users", "nickname", "id")
	//formList.AddField("Chat_id", "chat_id", db.Bigint, form.Number).FieldHide()
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("话题板块", "theme_id", db.Varchar, form.SelectSingle).
		FieldOptionsFromTable("ant_theme", "theme_name", "id")
	formList.AddField("Summary", "summary", db.Text, form.RichText)
	formList.AddField("Content", "content", db.Text, form.RichText)
	formList.AddField("Covers", "covers", db.JSON, form.Multifile).FieldDisplay(func(value types.FieldModel) interface{} {
		var covers []Cover
		err := json.Unmarshal([]byte(value.Value), &covers)
		if err != nil {
			return ""
		}

		var s []string
		for _, c := range covers {
			s = append(s, c.Url)
		}

		b, err := json.Marshal(s)
		if err != nil {
			return ""
		}
		value.Value = string(b)
		return string(b)
	}).FieldPostFilterFn(func(value types.PostFieldModel) interface{} {

		var covers []Cover
		for _, v := range value.Value {
			covers = append(covers, Cover{
				Url:   v,
				Title: "图片",
				Type:  Image,
			})
		}
		b, err := json.Marshal(covers)
		if err != nil {
			return sql.NullString{}
		}

		return json.JSON(b)
	})
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)

	formList.SetTable("ant_topics").SetTitle("AntTopics").SetDescription("AntTopics")

	return antTopics
}

func UnmarshalCovers(b []byte) []Cover {
	var covers []Cover
	err := json.Unmarshal(b, &covers)
	if err != nil {
		logrus.Errorf("解析 covers 失败:%s", err.Error())
		return nil
	}
	return covers
}
