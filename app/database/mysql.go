package database

import (
	"KosKita/app/config"
	"fmt"

	ud "KosKita/features/user/data"
	kd "KosKita/features/kos/data"
	bd "KosKita/features/booking/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&ud.User{}, &kd.BoardingHouse{}, &kd.Rating{}, &bd.Booking{})

	return DB
}
