package curd

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntH5Table(ctx *context.Context) table.Table {

	antH5 := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := antH5.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Bigint).FieldFilterable()
	info.AddField("名称", "title", db.Varchar)
	info.AddField("flag", "flag", db.Varchar)
	info.AddField("创建时间", "created_at", db.Datetime)
	info.AddField("更新时间", "updated_at", db.Datetime)

	info.SetTable("ant_h5")

	formList := antH5.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("flag", "flag", db.Varchar, form.Text)
	formList.AddField("详情", "rich_text", db.Mediumtext, form.RichText)
	formList.AddField("创建时间", "created_at", db.Datetime, form.Datetime)
	formList.AddField("更新时间", "updated_at", db.Datetime, form.Datetime)

	formList.SetTable("ant_h5")

	return antH5
}
