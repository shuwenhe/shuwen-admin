package curd

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAntUsersTable(ctx *context.Context) table.Table {

	antUsers := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	// 列表页
	listPage := antUsers.GetInfo()
	listPage.Title = "用户搜索"
	listPage.Description = "用户搜索"
	listPage.IsHideEditButton = true
	listPage.IsHideExportButton = true
	listPage.IsHideNewButton = true
	listPage.IsHideRowSelector = true
	listPage.IsHideFilterArea = true

	listPage.AddField("ID", "id", db.Bigint)
	listPage.AddField("头像", "avatar", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	}).FieldImage("50", "50")
	listPage.AddField("昵称", "nickname", db.Varchar)
	listPage.AddField("手机号", "phone", db.Varchar).FieldFilterable()
	listPage.AddField("性别", "sex", db.Smallint)
	listPage.AddField("最近登录", "recent_login_at", db.Datetime)

	listPage.DeleteHook = func(ids []string) error {
		return nil
	}

	listPage.SetTable("ant_users")
	// 详情页
	detailPage := antUsers.GetDetail()

	detailPage.AddField("ID", "id", db.Int)
	detailPage.AddField("昵称", "nickname", db.Varchar)
	detailPage.AddField("头像", "avatar", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	}).FieldImage("50", "50")
	detailPage.AddField("手机号", "phone", db.Varchar).FieldFilterable()
	detailPage.AddField("Email", "email", db.Varchar)
	detailPage.AddField("性别", "sex", db.Smallint)
	detailPage.AddField("年龄", "age", db.Int)
	detailPage.AddField("等级", "level", db.Int)
	detailPage.AddField("最近登录", "recent_login_at", db.Datetime)
	detailPage.AddField("注册时间", "created_at", db.Datetime)

	detailPage.SetTable("users").
		SetTitle("用户详情").
		SetDescription("用户详情")

	// 表单页
	formList := antUsers.GetForm()
	formList.AddField("ID", "id", db.Bigint, form.Default)
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Avatar", "avatar", db.Varchar, form.File).FieldDisplay(func(value types.FieldModel) interface{} {
		return GetOSSDomain(value.Value)
	})
	formList.AddField("Nickname", "nickname", db.Varchar, form.Text)
	formList.AddField("Sex", "sex", db.Smallint, form.Number)
	formList.AddField("Age", "age", db.Int, form.Number)
	formList.AddField("Level", "level", db.Int, form.Number)
	formList.AddField("Recent_login_at", "recent_login_at", db.Datetime, form.Datetime)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)

	formList.SetTable("ant_users").SetTitle("AntUsers").SetDescription("AntUsers")
	return antUsers
}
