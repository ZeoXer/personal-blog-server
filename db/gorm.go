package db

import (
	"go-server/global"
	model "go-server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db_config := global.CONFIG.Database
	dsn := db_config.Username + ":" + db_config.Password + "@tcp(" + db_config.Host + ":" + db_config.Port + ")/" + db_config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// AutoMigrate will create the tables in the database
func AutoMigrate() {
	if err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(model.User{}, model.Avatar{}); err != nil {
		panic("Error automigrating: " + err.Error())
	}
}
