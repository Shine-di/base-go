package mysql

import (
	"cortex3/conf"
	"cortex3/model/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

var sql *gorm.DB

func DB() *gorm.DB {

	if sql == nil {
		InitMysql()
	}
	return sql
}

func InitMysql() {
	host := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Yaml.Conf.Mysql.User, conf.Yaml.Conf.Mysql.Pwd, conf.Yaml.Conf.Mysql.Host, conf.Yaml.Conf.Mysql.Db)
	db, err := gorm.Open("mysql", host)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	sql = db
	// 全局禁用表名复数
	sql.SingularTable(true)
	
	sql.DB().SetMaxOpenConns(2000)
	sql.DB().SetMaxIdleConns(1000)

	if !sql.HasTable(&entity.SetObject{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.SetObject{})
	}
	if !sql.HasTable(&entity.SetData{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.SetData{})
	}
	if !sql.HasTable(&entity.SetStatus{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.SetStatus{})
	}
	if !sql.HasTable(&entity.Target{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Target{})
	}
	if !sql.HasTable(&entity.Data{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.Data{})
	}

	if !sql.HasTable(&entity.RecordData{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.RecordData{})
	}

	if !sql.HasTable(&entity.RecordStatus{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&entity.RecordStatus{})
	}

	fmt.Println("Mysql 初始化成功")
}
