package curd

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "ant_users" => http://localhost:9033/admin/info/ant_users
//
// example end
//
var Generators = map[string]table.Generator{
	"users":      GetAntUsersTable,
	"shortcodes": GetAntShortcodesTable,
	"theme":      GetAntThemeTable,
	"h5":         GetAntH5Table,
	"version":    GetAntVersionTable,
	"ad":         GetAntAdTable,
	"topic":      GetAntTopicsTable,
	// generators end
}
