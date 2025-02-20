package mysql

import (
	"byte_go/backend/app/product/biz/model"
	"byte_go/backend/app/product/conf"
	"fmt"
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

	needDemoData := !DB.Migrator().HasTable(&model.Product{})
	err = DB.AutoMigrate(&model.Product{}, &model.Category{})

	if os.Getenv("GO_ENV") != "online" && needDemoData {
		DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06',null, 'T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06',null, 'Sticker')")
		DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10',null,  'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ),                                       ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10',null,  'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ),( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20',null,  'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ),( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20',null,  'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ),( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35',null, 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ),( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20',null, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ),( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20',null, 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
		DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")

	}

	if err != nil {
		panic(err)
	}
}
