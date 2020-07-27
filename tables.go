package main

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/antdate/antdate-admin/internal/admin/curd"
)

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "ant_topics" => http://localhost:9033/admin/info/ant_topics
//
// example end
//
var Generators = map[string]table.Generator{
	"ant_topics": curd.GetAntTopicsTable,

	// generators end
}
