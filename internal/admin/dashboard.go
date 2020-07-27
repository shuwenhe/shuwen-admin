package admin

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"
	"github.com/antdate/antdate-admin/internal/admin/curd"
	"github.com/antdate/antdate-admin/internal/admin/upload"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Dashboard() {
	r := gin.New()
	e := engine.Default()
	uploader, err := upload.NewOSSUploader(viper.GetString("oss.endpoint"),
		viper.GetString("oss.accessKeyID"), viper.GetString("oss.accessKeySecret"), viper.GetString("oss.bucketName"))
	if err != nil {
		return
	}

	file.AddUploader(upload.UploaderOSS, func() file.Uploader {
		return uploader
	})

	cfg := config.Config{
		Title:      "蚂蚁捡友",
		Theme:      "sword",
		LoginTitle: "蚂蚁捡友",
		Logo:       template.HTML("<b>ant</b>蚂蚁捡友"),
		MiniLogo:   template.HTML("<b>ant</b>"),
		Databases: config.DatabaseList{
			"default": {
				Host:       viper.GetString("db.host"),
				Port:       viper.GetString("db.port"),
				User:       viper.GetString("db.username"),
				Pwd:        viper.GetString("db.password"),
				Name:       viper.GetString("db.dbname"),
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "/",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:           language.CN,
		IndexUrl:           "/",
		Debug:              true,
		AccessAssetsLogOff: true,
		Animation: config.PageAnimation{
			Type: "fadeInUp",
		},
		ColorScheme: adminlte.ColorschemeSkinRed,
		FileUploadEngine: config.FileUploadEngine{
			Name:   upload.UploaderOSS,
			Config: map[string]interface{}{},
		},
	}

	template.AddComp(chartjs.NewChart())

	if err := e.AddConfig(cfg).
		AddGenerators(datamodel.Generators).
		AddGenerators(curd.Generators).
		AddDisplayFilterXssJsFilter().
		AddNavButtons("服务管理 portainer", "", action.JumpInNewTab(viper.GetString("portainer.addr"), "portainer")).
		AddNavButtons("接口文档", "", action.JumpInNewTab(viper.GetString("docs.addr"), "postman")).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	e.HTML("GET", "/", datamodel.GetContent)

	//go func() {
	_ = r.Run(viper.GetString("addr"))
	//}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	e.MysqlConnection().Close()
}
