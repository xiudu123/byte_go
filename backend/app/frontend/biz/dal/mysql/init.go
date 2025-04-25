package mysql

import (
	"byte_go/backend/app/frontend/conf"
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"
)

var (
	DB       *gorm.DB
	err      error
	NeedDate bool
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	NeedDate = !DB.Migrator().HasTable(&gormadapter.CasbinRule{})
	err = DB.AutoMigrate(&gormadapter.CasbinRule{})
	if err != nil {
		panic(err)
	}
}
