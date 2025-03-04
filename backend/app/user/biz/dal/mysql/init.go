package mysql

import (
	"byte_go/backend/app/user/biz/model"
	"byte_go/backend/app/user/conf"
	"fmt"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
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
	needDemoData := !DB.Migrator().HasTable(&model.User{})
	err = DB.AutoMigrate(&model.User{})
	if conf.GetEnv() != "online" && needDemoData {
		DB.Exec("INSERT INTO `user`.`user` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10',null, 'user1@qq.com', '$2a$10$wywzamxDGWdAgrTPgbrwKOkD4hhT8gcfH0fAGctVYK3.RCMJ71ZOu', 'user1', 'https://p3-passport.byteacctimg.com/img/mosaic-legacy/3796/2975850990~50x50.awebp')")
		DB.Exec("INSERT INTO `user`.`user` VALUES ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10',null, 'user2@qq.com', '$2a$10$wywzamxDGWdAgrTPgbrwKOkD4hhT8gcfH0fAGctVYK3.RCMJ71ZOu', 'user2', 'https://p3-passport.byteacctimg.com/img/mosaic-legacy/3796/2975850990~50x50.awebp')")
		DB.Exec("INSERT INTO `user`.`user` VALUES ( 3, '2023-12-06 15:26:19', '2023-12-09 22:29:10',null, 'user3@qq.com', '$2a$10$wywzamxDGWdAgrTPgbrwKOkD4hhT8gcfH0fAGctVYK3.RCMJ71ZOu', 'user3', 'https://p3-passport.byteacctimg.com/img/mosaic-legacy/3796/2975850990~50x50.awebp')")
	}
	if err != nil {
		panic(err)
	}
}
