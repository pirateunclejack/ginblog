package model

import (
	"fmt"
	"ginblog/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() {

	dsn := fmt.Sprintf("host= %s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		utils.DbHost,
		utils.DbUser,
		utils.DbPassWord,
		utils.DbName,
		utils.DbPort,
		utils.DbSslMode,
		utils.DbTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database connect failed with error: ", err)
	}

	// disalbe gingulartable
	// db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqldb, err := db.DB()

	if err != nil {
		fmt.Println("Get sql.db from gorm.db error: ", err)
	}

	// SetMaxIdleConns set most of idle connections in connection pool
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns set most connections
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime set max time which a connect can be used again
	sqldb.SetConnMaxLifetime(10)

}
