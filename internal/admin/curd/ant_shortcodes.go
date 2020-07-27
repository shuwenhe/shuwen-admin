package curd

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntShortcodesTable(ctx *context.Context) table.Table {

	antShortcodes := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	ListPage := antShortcodes.GetInfo().
		HideFilterArea().
		HideDeleteButton().
		HideDetailButton().
		HideEditButton().
		HideExportButton().
		HideNewButton().
		SetTitle("短信验证码")

	ListPage.AddField("ID", "id", db.Bigint)
	ListPage.AddField("手机号", "phone", db.Varchar).FieldFilterable()
	ListPage.AddField("验证码", "code", db.Varchar)
	ListPage.AddField("验证码类型", "use_type", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		switch value.Value {
		case "1":
			return "注册"
		case "2":
			return "登录"
		case "3":
			return "密码"
		default:
			return "未知"
		}
	})

	ListPage.AddField("用户IP", "ip", db.Varchar)
	ListPage.AddField("发送时间", "created_at", db.Datetime)

	ListPage.SetTable("ant_shortcodes")

	formList := antShortcodes.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("手机号", "phone", db.Varchar, form.Text)
	formList.AddField("验证码", "code", db.Varchar, form.Text)
	formList.AddField("验证码类型", "use_type", db.Int, form.Number)
	formList.AddField("用户ip", "ip", db.Varchar, form.Ip)

	formList.SetTable("ant_shortcodes")

	return antShortcodes
}
