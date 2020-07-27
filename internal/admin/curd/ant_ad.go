package curd

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntAdTable(ctx *context.Context) table.Table {

	antAd := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	listPage := antAd.GetInfo().HideFilterArea().HideExportButton()

	listPage.AddField("ID", "id", db.Bigint)
	listPage.AddField("广告位", "slot_id", db.Bigint)
	listPage.AddField("广告名称", "title", db.Varchar).FieldFilterable()
	listPage.AddField("Banner", "banner_image", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	}).FieldImage("100", "50")
	listPage.AddField("跳转链接", "banner_link", db.Varchar)
	listPage.AddField("跳转方式", "target_type", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		switch value.Value {
		case "1":
			return "本窗口打开"
		case "2":
			return "新窗口打开"
		default:
			return "未知"
		}
	})
	listPage.AddField("Status", "status", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		switch value.Value {
		case "1":
			return "正常"
		case "2":
			return "隐藏"
		default:
			return "未知"
		}
	})
	listPage.AddField("创建时间", "created_at", db.Datetime)
	listPage.SetTable("ant_ad")

	formList := antAd.GetForm()
	formList.AddField("ID", "id", db.Bigint, form.Default)
	formList.AddField("广告位", "slot_id", db.Bigint, form.Number)
	formList.AddField("广告名称", "title", db.Varchar, form.Text)
	formList.AddField("Banner", "banner_image", db.Varchar, form.File).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	})
	formList.AddField("跳转链接", "banner_link", db.Varchar, form.Text)
	formList.AddField("跳转方式", "target_type", db.Int, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "本窗口打开", Value: "1"},
		{Text: "新窗口打开", Value: "2"},
	})
	formList.AddField("Status", "status", db.Int, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "正常", Value: "1"},
		{Text: "隐藏", Value: "2"},
	})

	formList.AddField("创建时间", "created_at", db.Datetime, form.Datetime)
	formList.AddField("更新时间", "updated_at", db.Datetime, form.Datetime)
	formList.SetTable("ant_ad")

	return antAd
}
