package curd

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntThemeTable(ctx *context.Context) table.Table {

	antTheme := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := antTheme.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Bigint).FieldFilterable()
	info.AddField("主题名称", "theme_name", db.Varchar)
	info.AddField("主题图片", "img_url", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	}).FieldImage("200", "100")
	info.AddField("主题介绍", "summary", db.Varchar).FieldWidth(300)
	info.AddField("创建时间", "created_at", db.Datetime)

	info.SetTable("ant_theme")

	formList := antTheme.GetForm()
	formList.AddField("ID", "id", db.Bigint, form.Default)
	formList.AddField("主题名称", "theme_name", db.Varchar, form.Text)
	formList.AddField("主题图片", "img_url", db.Varchar, form.File).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	})
	formList.AddField("主题介绍", "summary", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Datetime, form.Datetime)
	formList.AddField("更新时间", "updated_at", db.Datetime, form.Datetime)

	formList.SetTable("ant_theme")

	return antTheme
}
