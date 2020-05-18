package db

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var modelList []Model

var db *gorm.DB

func Setup() {
	var err error
	var conn = viper.GetString("mysql.dsn")
	//log.Debug(conn)
	db, err = gorm.Open("mysql", conn)
	if err != nil {
		fmt.Printf("Failed to connect database %s\n", err.Error())
		return
		//log.Fatal(fmt.Sprintf("Failed to connect database %s", err.Error()))
	} else {
		db.DB().SetMaxIdleConns(viper.GetInt("mysql.pool.min"))
		db.DB().SetMaxOpenConns(viper.GetInt("mysql.pool.max"))
		if gin.Mode() != gin.ReleaseMode {
			db.LogMode(true)
		}
	}
	//log.Info("Successfully connect to database")
	fmt.Println("Successfully connect to database")
}

func SetGormDB(gdb *gorm.DB) {
	db = gdb
}

func GetGormDB() *gorm.DB {
	return db
}

func SetDB() *sql.DB {
	return db.DB()
}

func GetDB() *sql.DB {
	return db.DB()
}

type Model interface {
	TableName() string
}

func Register(model Model) {
	rv := reflect.ValueOf(model)
	if rv.IsNil() {
		panic("register model failed, model is nil")
	}
	for _, m := range modelList {
		if m.TableName() == model.TableName() {
			panic("register model failed, already have the table name:" + model.TableName())
		}
	}
	modelList = append(modelList, model)
}

func AutoMigrate() error {
	for _, model := range modelList {
		if err := db.Debug().Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(model).Error; err != nil {
			return errors.Wrap(err, "db auto migrate failed")
		}
	}
	return nil
}
