package main

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntAdTable(ctx *context.Context) table.Table {

	antAd := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := antAd.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).FieldFilterable()
	info.AddField("Slot_id", "slot_id", db.Bigint)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Banner_image", "banner_image", db.Varchar)
	info.AddField("Banner_link", "banner_link", db.Varchar)
	info.AddField("Target_type", "target_type", db.Int)
	info.AddField("Status", "status", db.Int)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("ant_ad").SetTitle("AntAd").SetDescription("AntAd")

	formList := antAd.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Slot_id", "slot_id", db.Bigint, form.Number)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Banner_image", "banner_image", db.Varchar, form.Text)
	formList.AddField("Banner_link", "banner_link", db.Varchar, form.Text)
	formList.AddField("Target_type", "target_type", db.Int, form.Number)
	formList.AddField("Status", "status", db.Int, form.Number)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)

	formList.SetTable("ant_ad").SetTitle("AntAd").SetDescription("AntAd")

	return antAd
}
