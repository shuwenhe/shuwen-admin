package curd

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntVersionTable(ctx *context.Context) table.Table {

	antVersion := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := antVersion.GetInfo().HideFilterArea()
	info.IsHideExportButton = true
	info.IsHideFilterArea = true

	info.AddField("ID", "id", db.Bigint).FieldFilterable()
	info.AddField("版本", "version", db.Varchar)
	info.AddField("版本介绍", "summary", db.Varchar)
	info.AddField("创建时间", "created_at", db.Datetime)
	info.AddField("修改时间", "updated_at", db.Datetime)

	info.SetTable("ant_version").SetTitle("版本管理")

	formList := antVersion.GetForm()
	formList.AddField("ID", "id", db.Bigint, form.Default)
	formList.AddField("版本", "version", db.Varchar, form.Text)
	formList.AddField("版本介绍", "summary", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Datetime, form.Datetime).FieldHide()
	formList.AddField("修改时间", "updated_at", db.Datetime, form.Datetime).FieldHide()

	formList.SetTable("ant_version")

	return antVersion
}
